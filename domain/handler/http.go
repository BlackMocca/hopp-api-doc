package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"archive/zip"

	"github.com/Blackmocca/hopp-api-doc/domain"
	"github.com/Blackmocca/hopp-api-doc/domain/constants"
	"github.com/Blackmocca/hopp-api-doc/domain/models"
	docs "github.com/Blackmocca/hopp-api-doc/hopp-cli/methods"
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type HttpHandler struct {
	datasource    domain.Datasource
	microsoftAuth domain.AuthProvider
}

func NewHttpHandler(datasource domain.Datasource, microsoftAuth domain.AuthProvider) HttpHandler {
	return HttpHandler{datasource: datasource, microsoftAuth: microsoftAuth}
}

func (h HttpHandler) readMetadata(path string) (map[string]interface{}, error) {
	var meta = make(map[string]interface{})
	f, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return meta, err
	}
	if errors.Is(err, os.ErrNotExist) {
		return meta, nil
	}

	if err := json.Unmarshal(f, &meta); err != nil {
		return meta, err
	}

	return meta, nil
}

func (h HttpHandler) getSessionUser(c echo.Context) *models.User {
	if user := c.Get("session"); user != nil {
		return user.(*models.User)
	}
	return nil
}

func (h HttpHandler) fillTeamCollectionMetaData(teams []models.Team) error {
	if len(teams) == 0 {
		return nil
	}
	meta, err := h.readMetadata(constants.TEAM_META_DATA_PATH)
	if err != nil {
		return err
	}

	if len(meta) > 0 {
		for index := range teams {
			if v, ok := meta[teams[index].Id]; ok {
				ti, err := time.Parse(time.RFC3339, cast.ToString(v.(map[string]interface{})["last_updated"]))
				if err != nil {
					return err
				}
				teams[index].LastUpdated = ti.Format(constants.TIMESTAMP_LAYOUT)
			}
		}
	}

	return nil
}

func (h HttpHandler) Index(c echo.Context) error {
	if session := h.getSessionUser(c); session == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	resp := map[string]interface{}{
		"user": h.getSessionUser(c),
	}
	return c.Render(http.StatusOK, "index", resp)
}

func (h HttpHandler) Login(c echo.Context) error {
	var ctx = c.Request().Context()
	var session = h.getSessionUser(c)
	var isAuth = (session != nil)
	if isAuth {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	providers, err := h.datasource.FetchAuthProviders(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := map[string]interface{}{
		"providers": providers,
	}
	return c.Render(http.StatusOK, "login", resp)
}

func (h HttpHandler) Download(c echo.Context) error {
	var collectionId = c.Param("collection_id")
	var name = c.QueryParam("name")
	var path = fmt.Sprintf("%s/%s", constants.DOCUMENT_BASE_PATH, collectionId)
	if name == "" {
		name = collectionId
	}
	var zipOutput = fmt.Sprintf("%s/%s.zip", constants.DOCUMENT_BASE_PATH, name)

	if err := h.zip(path, zipOutput); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	bu, err := os.ReadFile(zipOutput)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("Content-Type", "application/zip")
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.zip"`, name))
	c.Response().Write(bu)
	return nil
}

func (h HttpHandler) TeamCollection(c echo.Context) error {
	var ctx = context.Background()
	var user = c.Get("session").(*models.User)

	teams, err := h.datasource.FetchMyTeams(ctx, user.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := h.fillTeamCollectionMetaData(teams); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := map[string]interface{}{
		"teams": teams,
		"user":  h.getSessionUser(c),
	}
	return c.Render(http.StatusOK, "collection", resp)
}

func (h HttpHandler) MyCollection(c echo.Context) error {
	var user = h.getSessionUser(c)
	var teams, err = h.getMyCollection(user.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := map[string]interface{}{
		"teams": teams,
		"user":  user,
	}
	return c.Render(http.StatusOK, "mycollection", resp)
}

func (h HttpHandler) zip(source, target string) error {
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	// 2. Go through all the files of the source
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}

func (h HttpHandler) AuthProvider(c echo.Context) error {
	var ctx = c.Request().Context()
	var provider = constants.AuthProvider(strings.ToUpper(c.QueryParam("provider")))
	if !provider.Valid() {
		return echo.NewHTTPError(http.StatusBadRequest, "provider was invalid")
	}

	var uri string
	var err error
	switch provider {
	case constants.AUTH_PROVIDER_MICROSOFT:
		uri, err = h.microsoftAuth.LinkAuth(ctx)
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusTemporaryRedirect, uri)
}

func (h HttpHandler) AuthProviderCallback(c echo.Context) error {
	var ctx = c.Request().Context()
	var code = c.QueryParam("code")
	var state = c.QueryParam("state")
	var provider = constants.AuthProvider(strings.ToUpper(c.Param("provider")))

	var profile *models.ProfileProvider
	var errProfile error
	switch provider {
	case constants.AUTH_PROVIDER_MICROSOFT:
		jwt, err := h.microsoftAuth.SignToken(ctx, code, state)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		profile, errProfile = h.microsoftAuth.Me(ctx, jwt.AccessToken)
	}

	if errProfile != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errProfile.Error())
	}

	var user, err = h.datasource.FetchOneUser(ctx, profile.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errProfile.Error())
	}

	if user == nil {
		resp := map[string]interface{}{
			"error":             "User Not Found",
			"error_description": strings.Title("Please Register User at hoppscotch app before usage this app"),
		}
		return c.Render(http.StatusOK, "error", resp)
	}

	c.SetCookie(&http.Cookie{
		Name:     constants.COOKIE_SESSION_NAME,
		Value:    user.EncodeHex(),
		Expires:  time.Now().AddDate(0, 0, 30),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h HttpHandler) Signout(c echo.Context) error {
	_, err := c.Cookie(constants.COOKIE_SESSION_NAME)
	if err != nil && err != http.ErrNoCookie {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.SetCookie(&http.Cookie{
		Name:     constants.COOKIE_SESSION_NAME,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return c.JSON(http.StatusOK, "success")
}

func (h HttpHandler) ImportCollection(c echo.Context) error {
	var session = h.getSessionUser(c)
	var file, err = c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	f, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer f.Close()

	var buf = bytes.Buffer{}
	io.Copy(&buf, f)

	/* manual + filename  */
	var now = time.Now().Format(time.RFC3339)
	var getDirName = func(prefix string, userId string, filename string) string {
		return fmt.Sprintf("%s/%s_manual_%s_%s", prefix, userId, now, slug.Make(strings.TrimSuffix(filename, ".json")))
	}
	var dirTeam = filepath.Join(getDirName("./assets", session.Id, file.Filename))
	os.Mkdir(dirTeam, 0755)

	var filePath = filepath.Join(dirTeam, "./data.json")
	os.WriteFile(filePath, buf.Bytes(), 0755)

	var pathOut = fmt.Sprintf(getDirName("./docs", session.Id, file.Filename))
	if err := docs.GenerateDocs(pathOut, filePath, 0, false, "./hopp-cli/templates"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return h.MyCollection(c)
}

func (h HttpHandler) DeleteMycollection(c echo.Context) error {
	var collectionId = c.Param("collection_id")
	var dirPath = "./docs"
	var fileInfo, err = ioutil.ReadDir(dirPath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for _, file := range fileInfo {
		if !file.IsDir() {
			continue
		}
		if index := strings.Index(file.Name(), collectionId); index == -1 {
			continue
		}

		path := filepath.Join(dirPath, file.Name())
		if err := os.RemoveAll(path); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return h.MyCollection(c)
}

func (h HttpHandler) getMyCollection(userId string) ([]models.Team, error) {
	var collections = make([]models.Team, 0)
	var fileInfo, err = ioutil.ReadDir("./docs")
	if err != nil && err != os.ErrNotExist {
		return collections, nil
	}

	for _, file := range fileInfo {
		if !file.IsDir() {
			continue
		}

		split := strings.Split(file.Name(), "_")
		if len(split) > 1 {
			if !strings.EqualFold(userId, split[0]) {
				continue
			}
			lastUpdated, err := time.Parse(time.RFC3339, split[2])
			if err != nil {
				return nil, err
			}
			name := fmt.Sprintf("%s_%s", lastUpdated.Format("20060102_150405"), split[3])
			var ptr = models.Team{
				Id:          file.Name(),
				Name:        name,
				LastUpdated: lastUpdated.Format(constants.TIMESTAMP_LAYOUT),
			}
			collections = append(collections, ptr)
		}
	}
	return collections, nil
}
