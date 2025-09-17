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

	if time.Now().After(verifyTokenData.ExpiresAt) {
		return Entity.User{}, errors.New("token has expired")
	}

	if verifyTokenData.IsUsed {
		return Entity.User{}, errors.New("token has already been used")
	}

	verifyTokenData.IsUsed = true

	if err := conn.Save(&verifyTokenData).Error; err != nil {
		return Entity.User{}, err
	}
	return verifyTokenData.User, nil
}
