package Auth

import (
	"SSO_BE_API/Config"
	DTOAuth "SSO_BE_API/Model/DTO/Auth"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Utils"
	"errors"
)

func LoginService(data DTOAuth.Login) (DTOAuth.Auth, error) {
	conn := Config.DB
	var userData Entity.User
	if err := conn.Where("email = ?", data.Email).First(&userData).Error; err != nil {
		return DTOAuth.Auth{}, err
	}

	if isValid := Utils.CompareHash(userData.Password, data.Password); !isValid {

		return DTOAuth.Auth{}, errors.New("Invalid password")
	}

	jwt := Utils.GenerateJwtToken(userData)
	refreshToken := Utils.GenerateRefreshToken(userData)

	SessionData := Entity.Session{
		UserId:       userData.ID,
		RefreshToken: refreshToken,
		JwtToken:     jwt,
	}

	if err := conn.Create(&SessionData).Error; err != nil {
		return DTOAuth.Auth{}, err
	}

	if data.CallbackURL != "" && data.ApplicationKey != "" {
		var ApplicationData Entity.Application
		if err := conn.Preload("CallbackApplication").Where("application_key = ?", data.ApplicationKey).First(&ApplicationData).Error; err != nil {
			return DTOAuth.Auth{}, err
		}

		isWhitelisted := false
		for _, v := range ApplicationData.CallbackApplication {
			if v.Callback == data.CallbackURL {
				isWhitelisted = true
				break
			}
		}

		if !isWhitelisted {
			return DTOAuth.Auth{}, errors.New("callback application is not authorized")
		}

		VerifyToken := Utils.GenerateVerifyToken(userData)
		VerifyTokenData := Entity.VerifyToken{
			Token:          VerifyToken,
			UserId:         userData.ID,
			ApplicationKey: data.ApplicationKey,
		}
		if err := conn.Create(&VerifyTokenData).Error; err != nil {
			return DTOAuth.Auth{}, err
		}

		data.GetCallbackUrlWithToken(VerifyToken)
	}

	return DTOAuth.Auth{
		Token:        SessionData.JwtToken,
		RefreshToken: SessionData.RefreshToken,
		CallbackUrl:  data.CallbackURL,
	}, nil
}
