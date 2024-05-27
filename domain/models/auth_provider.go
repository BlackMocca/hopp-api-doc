package models

import "github.com/Blackmocca/hopp-api-doc/domain/constants"

type SSOClient struct {
	AuthProvider constants.AuthProvider
	ProviderURI  string
	ClientId     string
	ClientSecret string
	RedirectURI  string
	Scope        string
	Tenant       string
}

func NewAuthMicrosoftClient(tenant string, clientId string, clientSecret string, redirectURI string, scope string) *SSOClient {
	return &SSOClient{
		AuthProvider: constants.AUTH_PROVIDER_MICROSOFT,
		ProviderURI:  "https://login.microsoftonline.com",
		ClientId:     clientId,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
		Scope:        scope,
		Tenant:       tenant,
	}
}
