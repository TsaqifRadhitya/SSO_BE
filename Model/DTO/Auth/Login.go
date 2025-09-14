package DTO

import (
	"net/url"
)

type Login struct {
	Email          string `form:"email" json:"email" validate:"required"`
	Password       string `form:"password" json:"password" validate:"required"`
	ApplicationKey string `form:"application_key" json:"application_key" validate:"required_with=CallbackURL"`
	CallbackURL    string `form:"callback_url" json:"callback_url" validate:"required_with=ApplicationKey,omitempty,url"`
}

func (req *Login) GetCallbackUrlWithToken(token string) {
	if req.CallbackURL == "" {
		return
	}
	url, err := url.Parse(req.CallbackURL)
	if err != nil {
		return
	}
	q := url.Query()
	q.Add("token", token)
	q.Add("access_token", req.ApplicationKey)
	url.RawQuery = q.Encode()
	callbackUrl := url.String()
	req.CallbackURL = callbackUrl
	return
}
