package Application

import (
	DTOApplication "SSO_BE_API/Model/DTO/Application"
	"SSO_BE_API/Model/Entity"
)

func UpdateApplicationCallbackService(request DTOApplication.UpdateApplicationCallaback) (Entity.CallbackApplication, error) {
	return Entity.CallbackApplication{}, nil
}
