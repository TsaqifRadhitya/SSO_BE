package Auth

import (
	"SSO_BE_API/Config"
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Utils"
	"errors"
	"strconv"
)

func SSOService(data DTOAuth.SSO) (string, error) {
	conn := Config.DB
	var ApplicationData Entity.Application
	if err := conn.Preload("CallbackApplication").Where("application_key = ?", data.ApplicationKey).First(&ApplicationData).Error; err != nil {
		return "", err
	}

	isAuthorizedCallback := false

	for _, Callback := range ApplicationData.CallbackApplication {
		if Callback.Callback == data.CallbackUrl {
			isAuthorizedCallback = true
			break
		}
	}

	if !isAuthorizedCallback {
		return "", errors.New("Callback is Not Authorized")
	}

	userId, _ := strconv.Atoi(data.UserId)
	verifyToken := Utils.GenerateVerifyToken(Entity.User{
		ID: userId,
	})
	data.GetCallbackUrlWithToken(verifyToken)
	return data.CallbackUrl, nil
}
