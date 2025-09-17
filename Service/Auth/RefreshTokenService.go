package Auth

import (
	"SSO_BE_API/Config"
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Utils"
	"errors"
	"time"
)

func RefreshTokenService(refres_token string) (DTOAuth.Auth, error) {
	var session Entity.Session
	conn := Config.DB
	if err := conn.Preload("User").Where("refresh_token = ?", refres_token).First(&session).Error; err != nil {
		return DTOAuth.Auth{}, err
	}

	if time.Now().After(session.RefreshExpiry) {
		return DTOAuth.Auth{}, errors.New("Refresh Token Expired")
	}

	NewRefreshToken := Utils.GenerateRefreshToken(session.User)
	NewJwtToken := Utils.GenerateJwtToken(session.User)

	session.RefreshToken = NewRefreshToken
	session.JwtToken = NewJwtToken
	session.RefreshExpiry = time.Now().AddDate(0, 0, 1)
	session.JwtExpiry = time.Now().Add(time.Minute * 15)

	if err := conn.Save(&session).Error; err != nil {
		return DTOAuth.Auth{}, err
	}

	return DTOAuth.Auth{
		RefreshToken: NewRefreshToken,
		Token:        NewJwtToken,
	}, nil
}
