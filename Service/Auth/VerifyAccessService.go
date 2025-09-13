package Auth

import (
	"SSO_BE_API/Config"
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/Entity"
)

func VerifyAccessService(request DTOAuth.VerifyAccess) bool {
	conn := Config.DB
	var ApplicationData Entity.Application
	if err := conn.Preload("CallbackApplication").Where("application_key = ?", request.ApplicationKey).First(&ApplicationData).Error; err != nil {
		return false
	}

	for _, Callback := range ApplicationData.CallbackApplication {
		if Callback.Callback == request.CallbackURL {
			return true
		}
	}
	return false
}
