package Auth

import (
	"SSO_BE_API/Config"
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Utils"
	"errors"
	"time"
)

func RefreshTokenService(data DTOAuth.RefreshToken) (DTOAuth.Auth, error) {
	var session Entity.Session
	conn := Config.DB
	if err := conn.Preload("User").Where("jwt_token = ? and refresh_token = ?", data.JwtToken, data.RefreshToken).First(&session).Error; err != nil {
		return DTOAuth.Auth{}, err
	}

	if time.Now().After(session.RefreshExpiry) {
		return DTOAuth.Auth{}, errors.New("Refresh Token Expired")
	}

	NewRefreshToken := Utils.GenerateRefreshToken(session.User)
	NewJwtToken := Utils.GenerateJwtToken(session.User)

	session.RefreshToken = NewRefreshToken
	session.JwtToken = NewJwtToken

	if err := conn.Save(&session).Error; err != nil {
		return DTOAuth.Auth{}, err
	}

	return DTOAuth.Auth{
		RefreshToken: NewRefreshToken,
		Token:        NewJwtToken,
	}, nil
}
