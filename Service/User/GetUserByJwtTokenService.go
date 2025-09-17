package User

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Model/Entity"
)

func GetUserByJwtTokenService(id string) (Entity.User, error) {
	var user Entity.User

	conn := Config.DB

	if err := conn.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
