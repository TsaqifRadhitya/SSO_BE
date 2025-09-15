package Application

import (
	"SSO_BE_API/Config"
	DTOApplication "SSO_BE_API/Model/DTO/Application"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Utils"
	"strconv"
)

func StoreApplicationService(request DTOApplication.StoreApplication) (Entity.Application, error) {
	var Application Entity.Application
	Application.ApplicationName = request.ApplicationName
	OwnerId, _ := strconv.Atoi(request.OwnerId)
	Application.OwnerId = OwnerId

	ApplicationCallbacks := []Entity.CallbackApplication{}

	if Application.CallbackApplication != nil {
		for _, callback := range request.CallbackUrls {
			ApplicationCallbacks = append(ApplicationCallbacks, Entity.CallbackApplication{
				Callback: callback,
			})
		}
	}

	conn := Config.DB

	trx := conn.Begin()
	if len(ApplicationCallbacks) > 0 {
		if err := conn.Create(&ApplicationCallbacks).Error; err != nil {
			return Entity.Application{}, err
		}
		Application.CallbackApplication = ApplicationCallbacks
	}

	generateApplicationKey, err := Utils.GenerateRandomString(32)
	
	if err != nil {
		trx.Rollback()
		return Entity.Application{}, err
	}

	hashedApplicationKey, errHash := Utils.CreateHash(generateApplicationKey)

	if errHash != nil {
		trx.Rollback()
		return Entity.Application{}, err
	}

	Application.ApplicationKey = hashedApplicationKey

	if err := conn.Create(&Application).Error; err != nil {
		trx.Rollback()
		return Entity.Application{}, err
	}

	trx.Commit()
	return Application, nil
}
