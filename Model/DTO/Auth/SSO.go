package DTO

type SSO struct {
	CalbackUrl string `json:"calback_url" validate:"required,url"`
}
