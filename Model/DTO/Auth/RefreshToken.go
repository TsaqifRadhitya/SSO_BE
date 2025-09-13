package DTO

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
	JwtToken     string `json:"jwt_token" validate:"required"`
}
