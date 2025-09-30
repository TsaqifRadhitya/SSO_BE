package User

import DTOUser "SSO_BE_API/Model/DTO/User"

func VerifyResetPasswordService(data DTOUser.VerifyResetPassword, compared DTOUser.VerifyResetPassword) bool {
	if data.Otp == compared.Otp && data.Email == compared.Email {
		return true
	}
	return false
}
