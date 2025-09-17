package Application

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Model/Entity"
	"errors"
	"gorm.io/gorm"
)

func GetAllApplicationService(id string) ([]Entity.Application, error) {
	conn := Config.DB
	var Applications []Entity.Application

	if err := conn.Preload("CallbackApplication").Where("owner_id = ?", id).Find(&Applications).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return make([]Entity.Application, 0), nil
		}
		return nil, err
	}

	return Applications, nil
}
