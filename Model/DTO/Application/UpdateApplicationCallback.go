package DTO

type UpdateApplicationCallaback struct {
	ApplicationId string `json:"application_id" validate:"required,numeric"`
	OwnerId       string `json:"owner_id" validate:"required,numeric"`
	CallbackId    string `json:"callback_id" validate:"required,numeric"`
	CallbackUrl   string `json:"callback_url" validate:"required,url"`
}
