package User

import (
	"SSO_BE_API/Config"
	DTOUser "SSO_BE_API/Model/DTO/User"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Utils"
	"errors"
)

func SetNewPasswordService(data DTOUser.SetNewPassword) error {
	conn := Config.DB
	var account Entity.User
	if err := conn.Where("email = ?", data.Email).First(&account).Error; err != nil {
		return err
	}
	if !isDifferentPass(data.Password, account.Password) {
		return errors.New("password already use on preview's password")
	}

	newPass, _ := Utils.CreateHash(data.Password)

	account.Password = newPass

	if err := conn.Save(&account).Error; err != nil {
		return err
	}

	return nil
}

func isDifferentPass(newPass string, hashedPass string) bool {
	return !Utils.CompareHash(hashedPass, newPass)
}
