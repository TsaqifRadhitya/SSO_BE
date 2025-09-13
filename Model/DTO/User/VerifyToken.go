package DTO

type VerifyToken struct {
	Token string `json:"verify_token" validate:"required,string"`
}
