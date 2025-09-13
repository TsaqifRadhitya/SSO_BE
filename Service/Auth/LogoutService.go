package Auth

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Model/Entity"
	"errors"
)

func LogoutService(jwt string) error {
	conn := Config.DB

	if conn.Where("jwt_token = ?", jwt).Delete(&Entity.Session{}).RowsAffected == 0 {
		return errors.New("Session Not Exist")
	}
	
	return nil
}
