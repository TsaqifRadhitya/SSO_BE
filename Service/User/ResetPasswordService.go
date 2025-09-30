package User

import (
	"SSO_BE_API/Config"
	DTOUser "SSO_BE_API/Model/DTO/User"
	"SSO_BE_API/Model/Entity"
	"SSO_BE_API/Provider"
	"fmt"
	"math/rand"
)

func ResetPasswordService(request DTOUser.ResetPassword) (int, error) {
	fmt.Println(request.Email)
	conn := Config.DB

	if err := conn.Where("email = ?", request.Email).First(&Entity.User{}).Error; err != nil {
		return 0, err
	}

	otp := rand.Intn(999999)
	savedOtp := fmt.Sprint(otp)
	mailerClient := Provider.InitClientMailer()

	if err := mailerClient.SendResetPasswordOTP(request.Email, savedOtp); err != nil {
		return 0, err
	}

	return otp, nil
}
