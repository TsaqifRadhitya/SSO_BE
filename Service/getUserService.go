package Service

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Entity"
	"SSO_BE_API/Model"
	"time"
)

func GetUserService(data Model.VerifyToken) (Entity.User, error) {
	var verifyToken Entity.VerifyToken
	if err := Config.DB.Where("token = ? AND expires_at > ?", data.Token, time.Now()).Preload("User").First(&verifyToken).Error; err != nil {
		return Entity.User{}, err
	}

	return verifyToken.User, nil
}
