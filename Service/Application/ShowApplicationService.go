package Application

import (
	"SSO_BE_API/Config"
	DTOApplication "SSO_BE_API/Model/DTO/Application"
	"SSO_BE_API/Model/Entity"
	"errors"
	"strconv"
)

func ShowApplicationService(request DTOApplication.ShowApplication) (Entity.Application, error) {
	var Application Entity.Application
	conn := Config.DB
	if err := conn.Where("id = ?", request.ApplicationId).First(&Application).Error; err != nil {
		return Entity.Application{}, err
	}

	OwnerId, _ := strconv.Atoi(request.OwnerId)
	if Application.OwnerId != OwnerId {
		return Entity.Application{}, errors.New("You don't have permission to view this application")
	}

	return Application, nil
}
