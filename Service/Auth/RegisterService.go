package Auth

import (
	"SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/Entity"
)

func RegisterService(json Auth.RegisterJson) (Entity.User, error) {
	return Entity.User{}, nil
}
