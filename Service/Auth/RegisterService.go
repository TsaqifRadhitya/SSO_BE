package Auth

import (
	"SSO_BE_API/Config"
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Utils"
)

func RegisterService(json DTOAuth.Register) (Entity.User, error) {
	hashedPassword, err := Utils.CreateHash(json.Password)
	if err != nil {
		return Entity.User{}, err
	}
	UserData := Entity.User{
		Name:     json.Name,
		Email:    json.Email,
		Password: hashedPassword,
	}

	conn := Config.DB

	if err = conn.Create(&UserData).Error; err != nil {
		return UserData, err
	}
	return UserData, nil
}
