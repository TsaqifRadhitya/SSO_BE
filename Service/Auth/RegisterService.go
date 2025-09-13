package Auth

import (
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/Entity"
)

func RegisterService(json DTOAuth.Register) (Entity.User, error) {
	return Entity.User{}, nil
}
