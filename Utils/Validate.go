package Utils

import "github.com/go-playground/validator/v10"

func Validate(data interface{}) map[string]string {
	v := validator.New()
	err := v.Struct(data)

	if err == nil {
		return nil
	}

	errors := map[string]string{}

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			errors[fieldErr.Field()] = fieldErr.Tag()
		}
	} else {
		errors["error"] = err.Error()
	}

	return errors
}
