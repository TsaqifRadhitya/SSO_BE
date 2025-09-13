package Application

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Model/Entity"
)

func GetAllApplicationService(id string) ([]Entity.Application, error) {
	conn := Config.DB
	var Applications []Entity.Application

	if err := conn.Where("owner_id = ?", id).Find(&Applications).Error; err != nil {
		return nil, err
	}
	
	return Applications, nil
}
