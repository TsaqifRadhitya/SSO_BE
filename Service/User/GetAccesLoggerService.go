package User

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Model/Entity"
)

func GetAccesLoggerService(id string) ([]Entity.AccessLog, error) {
	conn := Config.DB
	var AccessLog []Entity.AccessLog
	if err := conn.Preload("Application").Where("user_id = ?").Find(&AccessLog).Error; err != nil {
		return []Entity.AccessLog{}, nil
	}
	return AccessLog, nil
}
