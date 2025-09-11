package Model

import "net/url"

type LoginRequest struct {
	Email       string `json:"email" valid:"required,string"`
	Password    string `json:"password" valid:"required,string"`
	AccessToken string `json:"access_token" valid:"required,string"`
}

func (req LoginRequest) GetCallbackUrl(token string, callbackUrl string) (string, error) {
	url, err := url.Parse(callbackUrl)
	if err != nil {
		return "", err
	}
	q := url.Query()
	q.Add("token", token)
	url.RawQuery = q.Encode()

	return url.String(), nil
}
