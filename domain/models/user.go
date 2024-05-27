package models

import (
	"encoding/hex"
	"encoding/json"

	"github.com/guregu/null/zero"
)

type User struct {
	Id          string      `json:"id" db:"uid"`
	DisplayName zero.String `json:"display_name" db:"displayName"`
	Email       string      `json:"email" db:"email"`
	PhotoURL    zero.String `json:"photo_url" db:"photoURL"`
	IsAdmin     bool        `json:"is_admin" db:"isAdmin"`
}

func (u User) EncodeHex() string {
	src, _ := json.Marshal(u)
	return hex.EncodeToString(src)
}

func NewUserSession(session string) *User {
	bu, _ := hex.DecodeString(session)
	var user = User{}
	json.Unmarshal(bu, &user)
	return &user
}
