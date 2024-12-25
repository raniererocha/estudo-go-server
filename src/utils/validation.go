package utils

import (
	"github.com/go-playground/validator/v10"
)

func Validate[T any](obj T) []string {
	var validate = validator.New()

	var errors []string
	if err := validate.Struct(&obj); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, "O campo "+err.Field()+" é iválido ("+err.Tag()+")")
		}
	}

	return errors
}
