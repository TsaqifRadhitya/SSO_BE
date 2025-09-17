package Auth

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Model/Entity"
	"errors"
)

func LogoutService(refresh_token string) error {
	conn := Config.DB

	if conn.Where("refresh_token = ?", refresh_token).Delete(&Entity.Session{}).RowsAffected == 0 {
		return errors.New("Session Not Exist")
	}

	return nil
}
