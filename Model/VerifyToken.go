package Model

type VerifyToken struct {
	Token string `json:"token" validate:"required,string"`
}
