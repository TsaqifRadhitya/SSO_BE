package DTO

type DeleteApplicationCallaback struct {
	ApplicationId string `json:"application_id" validate:"required,numeric"`
	CallbackId    string `json:"callback_id" validate:"required,numeric"`
	OwnerId       string `json:"owner_id" validate:"required,numeric"`
}
