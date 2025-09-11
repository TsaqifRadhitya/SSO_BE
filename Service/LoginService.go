package Service

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Entity"
	"SSO_BE_API/Model"
	"time"
)

func LoginService(data Model.LoginRequest) (string, error) {
	var accessToken Entity.AccessToken
	if err := Config.DB.Where("token = ? AND expires_at > ?", data.AccessToken, time.Now()).Preload("Client").First(&accessToken).Error; err != nil {
		return "", err
	}
	callbackUrl, err := data.GetCallbackUrl(accessToken.Token, accessToken.Client.Callback)
	if err != nil {
		return "", err
	}
	return callbackUrl, nil
}
