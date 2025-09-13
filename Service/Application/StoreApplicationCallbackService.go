package Application

import (
	DTO "SSO_BE_API/Model/DTO/Application"
	"SSO_BE_API/Model/Entity"
)

func StoreApplicationCallbackService(request DTO.StoreApplicationCallback) (Entity.CallbackApplication, error) {
	return Entity.CallbackApplication{}, nil
}
