package Auth

type RefreshTokenJson struct {
	RefresToken string `json:"refres_token" validate:"required"`
}
