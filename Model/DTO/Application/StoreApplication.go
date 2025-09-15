package DTO

type StoreApplication struct {
	OwnerId         string   `json:"owner_id" validate:"required,numeric"`
	ApplicationName string   `json:"application_name" form:"application_name" validate:"required,max=255"`
	CallbackUrls    []string `json:"callback_url" form:"callback_url" validate:"dive,url"`
}
