package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"archive/zip"

	"github.com/Blackmocca/hopp-api-doc/domain"
	"github.com/Blackmocca/hopp-api-doc/domain/constants"
	"github.com/Blackmocca/hopp-api-doc/domain/models"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type HttpHandler struct {
	datasource domain.Datasource
}

func NewHttpHandler(datasource domain.Datasource) HttpHandler {
	return HttpHandler{datasource: datasource}
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
	fmt.Println("cookie context", c.Cookies())

	return c.Render(http.StatusOK, "index", nil)
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
	teams, err := h.datasource.FetchAllTeams(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := h.fillTeamCollectionMetaData(teams); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := map[string]interface{}{
		"teams": teams,
	}
	return c.Render(http.StatusOK, "collection", resp)
}

func (h HttpHandler) MyCollection(c echo.Context) error {
	var userId = c.Param("user_id")
	fmt.Println(userId)

	var teams = []models.Team{
		{
			Id:   "1",
			Name: "Wait for signin microsoft",
		},
	}
	resp := map[string]interface{}{
		"teams": teams,
	}
	return c.Render(http.StatusOK, "collection", resp)
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
