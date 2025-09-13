package DTO

import (
	"net/url"
)

type Login struct {
	Email          string  `json:"email" validate:"required,string"`
	Password       string  `json:"password" validate:"required,string"`
	ApplicationKey *string `json:"application_key" validate:"required,string"`
	CallbackURL    *string `json:"callback_url" validate:"required,string,url"`
}

func (req *Login) GetCallbackUrlWithToken(token string) {
	if *req.CallbackURL == "" {
		return
	}
	url, err := url.Parse(*req.CallbackURL)
	if err != nil {
		return
	}
	q := url.Query()
	q.Add("token", token)
	q.Add("access_token", *req.ApplicationKey)
	url.RawQuery = q.Encode()
	callbackUrl := url.String()
	req.CallbackURL = &callbackUrl
	return
}
