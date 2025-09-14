package DTO

type VerifyToken struct {
	Token string `form:"verify_token"  validate:"required"`
}
