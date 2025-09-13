package DTO

type StoreApplication struct {
	OwnerId      string   `json:"owner_id" validate:"required,numeric"`
	CallbackUrls []string `json:"callback_url" validate:"required,dive,url"`
}
