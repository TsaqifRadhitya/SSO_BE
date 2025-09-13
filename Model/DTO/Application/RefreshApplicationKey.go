package DTO

type RefreshApplicationKey struct {
	ApplicationId string `json:"application_id" validate:"required,numeric"`
	OwnerId       string `json:"owner_id" validate:"required,numeric"`
}
