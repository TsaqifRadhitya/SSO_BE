package DTO

type RefreshToken struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" validate:"required"`
	JwtToken     string `form:"jwt_token" json:"jwt_token" validate:"required"`
}
