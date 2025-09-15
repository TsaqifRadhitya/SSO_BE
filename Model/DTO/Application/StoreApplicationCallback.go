package DTO

type StoreApplicationCallback struct {
	ApplicationId string `json:"application_id" validate:"required,numeric"`
	OwnerId       string `json:"owner_id" validate:"required,numeric"`
	CallbackUrl   string `json:"callback_url" form:"callback_url" validate:"required,url"`
}
