package DTO

type ResponseSuccess[T interface{}] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ResponseError[T interface{}] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   T      `json:"error,omitempty"`
}
