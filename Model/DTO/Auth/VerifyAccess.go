package DTO

type VerifyAccess struct {
	AccessToken string `json:"access_token" validate:"required,string"`
	CallbackURL string `json:"callback_url" validate:"required,string,url"`
}
