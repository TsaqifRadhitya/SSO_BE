package DTO

type VerifyAccess struct {
	ApplicationKey string `json:"application_key" validate:"required,string"`
	CallbackURL    string `json:"callback_url" validate:"required,string,url"`
}
