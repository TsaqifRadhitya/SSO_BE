package Application

import (
	DTOApplication "SSO_BE_API/Model/DTO/Application"
	"SSO_BE_API/Model/Entity"
)

func StoreApplicationService(request DTOApplication.StoreApplication) (Entity.Application, error) {
	return Entity.Application{}, nil
}
