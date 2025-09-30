package Auth

import (
	"SSO_BE_API/Config"
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Utils"
	"fmt"
)

func RegisterService(json DTOAuth.Register) (DTOAuth.Auth, error) {
	hashedPassword, err := Utils.CreateHash(json.Password)
	if err != nil {
		return DTOAuth.Auth{}, err
	}
	UserData := Entity.User{
		Name:     json.Name,
		Email:    json.Email,
		Password: hashedPassword,
		Phone:    json.Phone,
	}

	conn := Config.DB

	if err = conn.Create(&UserData).Error; err != nil {
		fmt.Println("error hash:", err)
		return DTOAuth.Auth{}, err
	}

	authCredential, _ := LoginService(DTOAuth.Login{
		Email:    json.Email,
		Password: json.Password,
	})

	fmt.Println(authCredential)

	return authCredential, nil
}
