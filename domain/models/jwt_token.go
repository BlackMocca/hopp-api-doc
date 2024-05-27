package models

import (
	"github.com/gofrs/uuid"
	LibJwt "github.com/golang-jwt/jwt/v5"
)

type JWTToken struct {
	AccessToken                string                   `json:"access_token"`
	RefreshToken               string                   `json:"refresh_token"`
	AccessTokenRegisterClaims  *LibJwt.RegisteredClaims `json:"access_token_cliams"`
	RefreshTokenRegisterClaims *LibJwt.RegisteredClaims `json:"refresh_token_cliams"`
}

func NewJWTToken(accessToken string, refreshToken string, cliams *LibJwt.RegisteredClaims, refreshTokenCliams *LibJwt.RegisteredClaims) *JWTToken {
	return &JWTToken{
		AccessToken:                accessToken,
		RefreshToken:               refreshToken,
		AccessTokenRegisterClaims:  cliams,
		RefreshTokenRegisterClaims: refreshTokenCliams,
	}
}

type JWTTokenPayload struct {
	JTI         *uuid.UUID `json:"jti"`
	Id          *uuid.UUID `json:"id"`
	Username    string     `json:"username"`
	SSOClientId *uuid.UUID `json:"sso_client_id"`
}

func NewJwtTokenPayload(userId, SSOClientId *uuid.UUID, username string) JWTTokenPayload {
	jti, _ := uuid.NewV4()
	return JWTTokenPayload{
		JTI:         &jti,
		Id:          userId,
		Username:    username,
		SSOClientId: SSOClientId,
	}
}
