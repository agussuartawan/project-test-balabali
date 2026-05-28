package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error, s any) map[string][]string {
	errors := make(map[string][]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errors["error"] = []string{err.Error()}
		return errors
	}

	for _, fieldError := range validationErrors {
		field := fieldError.Field()

		var message string
		switch fieldError.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", field)
		case "email":
			message = fmt.Sprintf("%s is not a valid email address", field)
		case "min":
			message = fmt.Sprintf("%s must be at least %s characters long", field, fieldError.Param())
		case "max":
			message = fmt.Sprintf("%s must be at most %s characters long", field, fieldError.Param())
		default:
			message = fmt.Sprintf("%s is invalid", field)
		}

		errors[field] = append(errors[field], message)
	}

	return errors
}