package models

import (
	"encoding/json"
	"strings"

	"github.com/spf13/cast"
)

type ProfileProvider struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (p ProfileProvider) String() string {
	bu, _ := json.Marshal(p)
	return string(bu)
}

func NewProfileProviderWithParams(params map[string]interface{}) *ProfileProvider {
	var profile = new(ProfileProvider)

	for key, val := range params {
		switch key {
		case "upn":
			profile.Username = strings.ToLower(cast.ToString(val))
		case "userPrincipalName":
			profile.Username = strings.ToLower(cast.ToString(val))
		case "given_name":
			profile.Firstname = cast.ToString(val)
		case "family_name":
			profile.Lastname = cast.ToString(val)
		case "firstname":
			profile.Firstname = cast.ToString(val)
		case "lastname":
			profile.Lastname = cast.ToString(val)
		}
	}

	return profile
}
