package Model

type ResponseSuccess struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    *interface{} `json:"data,omitempty"`
}

type ResponseError struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Error   *interface{} `json:"error,omitempty"`
}
