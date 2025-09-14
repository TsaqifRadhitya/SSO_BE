package User

import (
	"SSO_BE_API/Config"
	DTOUser "SSO_BE_API/Model/DTO/User"
	"SSO_BE_API/Model/Entity"
	"errors"
	"time"
)

func GetUserByVerifyTokenService(request DTOUser.VerifyToken) (Entity.User, error) {
	conn := Config.DB
	var verifyTokenData Entity.VerifyToken

	if err := conn.Preload("User").Where("token = ?", request.Token).First(&verifyTokenData).Error; err != nil {
		return Entity.User{}, err
	}

	if verifyTokenData.ExpiresAt.After(time.Now()) {
		return Entity.User{}, errors.New("Token is expired")
	}

	return verifyTokenData.User, nil
}
