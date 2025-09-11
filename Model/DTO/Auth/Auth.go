package Auth

type Auth struct {
	Token        string  `json:"token"`
	RefreshToken string  `json:"refresh_token"`
	CallbackUrl  *string `json:"callback_url,omitempty"`
}
