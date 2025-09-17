package Auth

import (
	"SSO_BE_API/Config"
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/Entity"
	"net/url"
)

func VerifyAccessService(request DTOAuth.VerifyAccess) (bool, string) {
	conn := Config.DB
	var ApplicationData Entity.Application
	if err := conn.Preload("CallbackApplication").Where("application_key = ?", request.ApplicationKey).First(&ApplicationData).Error; err != nil {
		return false, ""
	}

	for _, Callback := range ApplicationData.CallbackApplication {
		requestBaseUrl, _ := url.Parse(request.CallbackURL)
		callbackBaseUrl, _ := url.Parse(Callback.Callback)
		if requestBaseUrl.Host == callbackBaseUrl.Host {
			return true, ApplicationData.ApplicationName
		}
	}
	return false, ApplicationData.ApplicationName
}
