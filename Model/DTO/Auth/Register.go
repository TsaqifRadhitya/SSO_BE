package DTO

type Register struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
	Name     string `json:"name" form:"name" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required,numeric,min=11,max=15"`
}
