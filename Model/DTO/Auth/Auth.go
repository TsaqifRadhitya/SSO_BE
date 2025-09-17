package DTO

type Auth struct {
	Token        string `json:"token"`
	RefreshToken string `json:"-"`
	CallbackUrl  string `json:"callback_url,omitempty"`
}
