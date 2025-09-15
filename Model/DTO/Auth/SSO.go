package DTO

import "net/url"

type SSO struct {
	CallbackUrl    string `json:"callback_url" form:"callback_url" validate:"required,url"`
	ApplicationKey string `json:"application_key" form:"application_key" validate:"required"`
	UserId         string `validate:"required,numeric"`
}

func (req *SSO) GetCallbackUrlWithToken(token string) {
	if req.CallbackUrl == "" {
		return
	}
	url, err := url.Parse(req.CallbackUrl)
	if err != nil {
		return
	}
	q := url.Query()
	q.Add("token", token)
	q.Add("access_token", req.ApplicationKey)
	url.RawQuery = q.Encode()
	callbackUrl := url.String()
	req.CallbackUrl = callbackUrl
	return
}
