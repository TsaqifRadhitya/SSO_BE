package Utils

import "SSO_BE_API/Model/DTO/Response"

func ErrorFormater(err error) DTO.ResponseError[*interface{}] {
	return DTO.ResponseError[*interface{}]{}
}
