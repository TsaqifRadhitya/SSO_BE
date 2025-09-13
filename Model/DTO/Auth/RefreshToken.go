package DTO

type RefreshToken struct {
	RefresToken string `json:"refres_token" validate:"required"`
}
