package Auth

import (
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
)

func LoginService(data DTOAuth.Login) (DTOAuth.Auth, error) {
	return DTOAuth.Auth{}, nil
}
