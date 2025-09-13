package Application

import (
	"SSO_BE_API/Config"
	DTO "SSO_BE_API/Model/DTO/Application"
	"SSO_BE_API/Model/Entity"
	"errors"
	"strconv"
)

func StoreApplicationCallbackService(request DTO.StoreApplicationCallback) (Entity.CallbackApplication, error) {
	var Application Entity.Application
	conn := Config.DB
	if err := conn.Where("id = ?", request.ApplicationId).First(&Application).Error; err != nil {
		return Entity.CallbackApplication{}, err
	}

	OwnerId, _ := strconv.Atoi(request.OwnerId)
	if Application.OwnerId != OwnerId {
		return Entity.CallbackApplication{}, errors.New("You are not the owner of the application")
	}
	NewCallbackApplication := Entity.CallbackApplication{
		Application: Application,
		Callback:    request.CallbackUrl,
	}

	if err := conn.Create(&NewCallbackApplication).Error; err != nil {
		return Entity.CallbackApplication{}, err
	}

	return NewCallbackApplication, nil
}
