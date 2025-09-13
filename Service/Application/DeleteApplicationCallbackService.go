package Application

import (
	"SSO_BE_API/Config"
	DTOApplication "SSO_BE_API/Model/DTO/Application"
	"SSO_BE_API/Model/Entity"
	"errors"
	"strconv"
)

func DeleteApplicationCallbackService(request DTOApplication.DeleteApplicationCallaback) error {
	var CallbackApplication Entity.CallbackApplication

	conn := Config.DB

	if err := conn.Preload("Application").Where("id = ?", request.CallbackId).First(&CallbackApplication).Error; err != nil {
		return err
	}
	OwnerId, _ := strconv.Atoi(request.OwnerId)

	if CallbackApplication.Application.OwnerId != OwnerId {
		return errors.New("You don't have the permission to delete the callback application")
	}

	if err := conn.Delete(&CallbackApplication).Error; err != nil {
		return err
	}
	return nil
}
