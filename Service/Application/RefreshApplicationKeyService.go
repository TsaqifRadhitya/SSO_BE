package Application

import (
	"SSO_BE_API/Config"
	DTOApplication "SSO_BE_API/Model/DTO/Application"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Utils"
	"errors"
	"strconv"
)

func RefreshApplicationKeyService(request DTOApplication.RefreshApplicationKey) (string, error) {
	var Application Entity.Application
	conn := Config.DB
	if err := conn.Where("id = ?", request.ApplicationId).Find(&Application).Error; err != nil {
		return "", err
	}

	OwnerId, _ := strconv.Atoi(request.OwnerId)
	if Application.OwnerId != OwnerId {
		return "", errors.New("You are not the owner of the application")
	}

	generateKey, err := Utils.GenerateRandomString(32)

	if err != nil {
		return "", err
	}
	newApplicationKey, errHash := Utils.CreateHash(generateKey)

	if errHash != nil {
		return "", err
	}

	Application.ApplicationKey = newApplicationKey

	if err = conn.Save(&Application).Error; err != nil {
		return "", err
	}

	return newApplicationKey, nil
}
