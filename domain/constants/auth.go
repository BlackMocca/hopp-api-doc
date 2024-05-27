package constants

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type AuthProvider string

const (
	AUTH_PROVIDER_KEYCLOAK  AuthProvider = "KEYCLOAK" /* KEYCLOAK */
	AUTH_PROVIDER_MICROSOFT AuthProvider = "MICROSOFT"
)

func (a AuthProvider) Valid() bool {
	providers := []AuthProvider{
		AUTH_PROVIDER_KEYCLOAK,
		AUTH_PROVIDER_MICROSOFT,
	}
	for _, item := range providers {
		if strings.EqualFold(string(a), string(item)) {
			return true
		}
	}
	return false
}

const (
	RESPONSE_TYPE                 = "code"
	GRANT_TYPE_AUTHORIZATION_CODE = "authorization_code"
	GRANT_TYPE_PASSWORD           = "password"
)

var (
	GET_PATH_AUTH_KEY_CLOAK = func(realm string, protocol string) string {
		return fmt.Sprintf("realms/%s/protocol/%s/auth", realm, protocol)
	}

	GET_PATH_AUTH_KEY_CLOAK_TOKEN = func(realm string, protocol string) string {
		return fmt.Sprintf("realms/%s/protocol/%s/token", realm, protocol)
	}

	GET_PATH_AUTH_KEY_CLOAK_SIGNOUT = func(realm string, protocol string) string {
		return fmt.Sprintf("realms/%s/protocol/%s/logout", realm, protocol)
	}

	GET_PATH_AUTH_KEY_CLOAK_ME = func(realm string, protocol string) string {
		return fmt.Sprintf("realms/%s/protocol/%s/userinfo", realm, protocol)
	}

	GET_PATH_AUTH_MICROSOFT = func(tenantId string) string {
		return fmt.Sprintf("%s/oauth2/v2.0/authorize", tenantId)
	}

	GET_PATH_AUTH_MICROSOFT_TOKEN = func(tenantId string) string {
		return fmt.Sprintf("%s/oauth2/v2.0/token", tenantId)
	}

	GET_PATH_AUTH_MICROSOFT_SIGNOUT = func(tenantId string) string {
		return fmt.Sprintf("%s/oauth2/v2.0/logout", tenantId)
	}

	GET_PATH_AUTH_MICROSOFT_ME = func() string {
		return "oidc/userinfo"
	}

	GET_STATE_ENCODE = func(state map[string]interface{}) string {
		if len(state) > 0 {
			bu, _ := json.Marshal(state)
			return base64.StdEncoding.EncodeToString(bu)
		}
		return ""
	}

	GET_STATE_DECODE = func(state string) (map[string]interface{}, error) {
		var m = map[string]interface{}{}
		stateBu, err := base64.StdEncoding.DecodeString(state)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(stateBu, &m)
		return m, nil
	}

	GET_PROVIDER_SESSION_KEY = func(authProvider AuthProvider, userPrincipal string) string {
		return fmt.Sprintf("provider_%s_%s", string(authProvider), userPrincipal)
	}
)
