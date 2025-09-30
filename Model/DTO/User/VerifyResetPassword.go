package DTO

type VerifyResetPassword struct {
	Otp   string `json:"otp" form:"otp" validate:"required"`
	Email string `json:"email" form:"email" validate:"required,email"`
}
