package Auth

import (
	"SSO_BE_API/Config"
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Utils"
	"errors"
	"net/url"
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
		requestBaseUrl, _ := url.Parse(data.CallbackUrl)
		callbackBaseUrl, _ := url.Parse(Callback.Callback)
		if requestBaseUrl.Host == callbackBaseUrl.Host {
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

	if err := conn.Create(&Entity.VerifyToken{
		Token:          verifyToken,
		ApplicationKey: data.ApplicationKey,
		UserId:         userId,
	}).Error; err != nil {
		return "", err
	}

	go func() {
		accessLogger := Entity.AccessLog{
			UserId:        userId,
			ApplicationId: int(ApplicationData.ID),
		}

		conn.Create(&accessLogger)

		connectedApplication := Entity.ConnectedApplication{
			UserId:        userId,
			ApplicationId: int(ApplicationData.ID),
		}

		conn.Where("application_id = ? and user_id = ?", int(ApplicationData.ID), userId).FirstOrCreate(&connectedApplication)
	}()

	data.GetCallbackUrlWithToken(verifyToken)
	return data.CallbackUrl, nil
}
