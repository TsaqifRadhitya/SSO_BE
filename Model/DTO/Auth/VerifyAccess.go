package DTO

type VerifyAccess struct {
	ApplicationKey string `json:"application_key" form:"application_key" validate:"required"`
	CallbackURL    string `json:"callback_url" form:"callback_url"  validate:"required,url"`
}
