package Auth

type SSOJson struct {
	CalbackUrl string `json:"calback_url" validate:"required,url"`
}
