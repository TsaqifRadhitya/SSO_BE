package Utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func Validate(data interface{}) map[string]string {
	fmt.Println(data)
	v := validator.New()
	err := v.Struct(data)

	fmt.Println(err)

	if err == nil {
		return nil
	}

	errors := make(map[string]string)

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrs {
			var msg string
			switch e.Tag() {
			case "required":
				msg = fmt.Sprintf("%s is required", e.Field())
			case "email":
				msg = fmt.Sprintf("%s must be a valid email", e.Field())
			case "min":
				msg = fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
			case "max":
				msg = fmt.Sprintf("%s must be at most %s characters", e.Field(), e.Param())
			default:
				msg = fmt.Sprintf("%s is not valid", e.Field())
			}

			errors[e.Field()] = msg
		}
	}

	return errors
}
