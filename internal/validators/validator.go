package validators

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func NewCustomValidator() *CustomValidator {
	v := validator.New()

	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		if name == "" {
			return field.Name
		}

		return name
	})

	return &CustomValidator{Validator: v}
}