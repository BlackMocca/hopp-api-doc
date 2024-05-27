package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/Blackmocca/hopp-api-doc/domain"
	"github.com/Blackmocca/hopp-api-doc/domain/constants"
	"github.com/Blackmocca/hopp-api-doc/domain/models"
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type microsoftProviderRepository struct {
	client  *models.SSOClient
	debug   bool
	timeout int
}

func NewMicrosoftProviderRepository(client *models.SSOClient, debug bool, timeout int) domain.AuthProvider {
	return &microsoftProviderRepository{client: client, debug: debug, timeout: timeout}
}

func (k microsoftProviderRepository) R(providerURI string) *resty.Request {
	return resty.New().
		SetDebug(k.debug).
		SetBaseURL(providerURI).
		SetTimeout(time.Duration(k.timeout) * time.Minute).
		R()
}

func (k microsoftProviderRepository) getBodyJSON(resp *resty.Response) (map[string]interface{}, error) {
	var m = make(map[string]interface{})
	if err := json.Unmarshal(resp.Body(), &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (a microsoftProviderRepository) LinkAuth(ctx context.Context) (string, error) {
	uri, err := url.Parse(a.client.ProviderURI)
	if err != nil {
		return "", err
	}
	uri.Path = constants.GET_PATH_AUTH_MICROSOFT(a.client.Tenant)
	uri.ForceQuery = true

	queryParams := uri.Query()
	queryParams.Add("client_id", a.client.ClientId)
	queryParams.Add("response_type", constants.RESPONSE_TYPE)
	queryParams.Add("response_mode", "query")
	queryParams.Add("scope", strings.Join(strings.Split(a.client.Scope, "+"), "+"))
	queryParams.Add("redirect_uri", a.client.RedirectURI)
	queryParams.Add("state", constants.GET_STATE_ENCODE(map[string]interface{}{"provider": string(a.client.AuthProvider)}))

	uri.RawQuery, err = url.QueryUnescape(queryParams.Encode())
	if err != nil {
		return "", err
	}

	return uri.String(), nil
}

func (a microsoftProviderRepository) SignToken(ctx context.Context, authorizeCode string, state string) (*models.JWTToken, error) {
	var path = constants.GET_PATH_AUTH_MICROSOFT_TOKEN(a.client.Tenant)
	var body = map[string]string{
		"client_id":     a.client.ClientId,
		"client_secret": a.client.ClientSecret,
		"code":          authorizeCode,
		"state":         state,
		"grant_type":    constants.GRANT_TYPE_AUTHORIZATION_CODE,
		"redirect_uri":  a.client.RedirectURI,
	}

	resp, err := a.R(a.client.ProviderURI).
		SetContentLength(true).
		SetHeader("Content-Type", echo.MIMEMultipartForm).
		SetFormData(body).
		Post(path)
	if err != nil {
		return nil, fmt.Errorf("can't connect to host %s on message %s", a.client.ProviderURI, err.Error())
	}

	respBody, err := a.getBodyJSON(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("error status '%s' on message '%s'", cast.ToString(respBody["error"]), cast.ToString(respBody["error_description"]))
	}

	jwtToken := &models.JWTToken{
		AccessToken:  cast.ToString(respBody["access_token"]),
		RefreshToken: cast.ToString(respBody["refresh_token"]),
		AccessTokenRegisterClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(cast.ToInt(respBody["expires_in"])))),
		},
	}
	if jwtToken.RefreshToken == "" {
		jwtToken.RefreshToken = jwtToken.AccessToken
	}

	return jwtToken, nil
}

func (a microsoftProviderRepository) Me(ctx context.Context, jwtToken string) (*models.ProfileProvider, error) {
	var path = constants.GET_PATH_AUTH_MICROSOFT_ME()
	resp, err := a.R("https://graph.microsoft.com").
		SetContentLength(true).
		SetAuthToken(jwtToken).
		Post(path)
	if err != nil {
		return nil, fmt.Errorf("can't connect to host %s on message %s", a.client.ProviderURI, err.Error())
	}
	if resp.StatusCode() == 401 {
		regexp := regexp.MustCompile(`error_description=+(.+)`)
		msg := strings.ReplaceAll(regexp.FindString(resp.Header().Get("Www-Authenticate")), "error_description=", "")
		return nil, errors.New(strings.Trim(msg, "\""))
	}

	respBody, err := a.getBodyJSON(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("error status '%s' on message '%s'", cast.ToString(respBody["error"]), cast.ToString(respBody["error_description"]))
	}

	return &models.ProfileProvider{
		Username: cast.ToString(respBody["email"]),
	}, nil
}
