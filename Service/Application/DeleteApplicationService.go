package Application

import (
	"SSO_BE_API/Config"
	DTOApplication "SSO_BE_API/Model/DTO/Application"
	"SSO_BE_API/Model/Entity"
	"errors"
	"strconv"
)

func DeleteApplicationService(request DTOApplication.DeleteApplication) error {
	var Application Entity.Application
	conn := Config.DB
	if err := conn.Where("id = ?", request.ApplicationId).First(&Application).Error; err != nil {
		return err
	}

	OwnerId, _ := strconv.Atoi(request.OwnerId)
	if Application.OwnerId != OwnerId {
		return errors.New("You are not the owner of the application")
	}

	if err := conn.Delete(&Application).Error; err != nil {
		return err
	}

	return nil
}
