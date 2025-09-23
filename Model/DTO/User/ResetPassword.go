package DTO

type ResetPassword struct {
	Email string `json:"email" form:"email" validate:"required,email"`
}
