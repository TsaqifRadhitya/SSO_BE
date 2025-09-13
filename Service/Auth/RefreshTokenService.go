package Auth

import (
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
)

func RefreshTokenService(data DTOAuth.RefreshToken) (DTOAuth.Auth, error) {
	return DTOAuth.Auth{}, nil
}
