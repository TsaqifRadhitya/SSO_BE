package Application

import (
	"SSO_BE_API/Config"
	DTOApplication "SSO_BE_API/Model/DTO/Application"
	"SSO_BE_API/Model/Entity"
	"errors"
	"strconv"
)

func UpdateApplicationCallbackService(request DTOApplication.UpdateApplicationCallaback) (Entity.CallbackApplication, error) {
	var CallbackApplication Entity.CallbackApplication

	conn := Config.DB

	if err := conn.Preload("Application").Where("id = ?", request.CallbackId).First(&CallbackApplication).Error; err != nil {
		return Entity.CallbackApplication{}, err
	}
	OwnerId, _ := strconv.Atoi(request.OwnerId)
	if CallbackApplication.Application.OwnerId != OwnerId {
		return Entity.CallbackApplication{}, errors.New("You don't have the permission to update the callback application")
	}

	CallbackApplication.Callback = request.CallbackUrl

	if err := conn.Save(&CallbackApplication).Error; err != nil {
		return Entity.CallbackApplication{}, err
	}

	return CallbackApplication, nil
}
