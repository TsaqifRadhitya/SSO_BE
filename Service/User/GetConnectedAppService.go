package User

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Model/Entity"
)

func GetConnectedAppService(id string) ([]Entity.ConnectedApplication, error) {
	conn := Config.DB
	var ConnectedApp []Entity.ConnectedApplication
	if err := conn.Preload("Application", "CallbackApplication").Where("user_id = ?").Find(&ConnectedApp).Error; err != nil {
		return []Entity.ConnectedApplication{}, nil
	}
	return ConnectedApp, nil
}
