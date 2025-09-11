package Auth

import (
	"errors"
	"net/url"
)

type LoginRequest struct {
	Email       string  `json:"email" validate:"required,string"`
	Password    string  `json:"password" validate:"required,string"`
	AccessToken string  `json:"access_token" validate:"required,string"`
	CallbackURL *string `json:"callback_url" validate:"required,string,url"`
}

func (req LoginRequest) GetCallbackUrl(token string) (string, error) {
	if *req.CallbackURL == "" {
		return "", errors.New("callback_url is required")
	}
	url, err := url.Parse(*req.CallbackURL)
	if err != nil {
		return "", err
	}
	q := url.Query()
	q.Add("token", token)
	url.RawQuery = q.Encode()

	return url.String(), nil
}
