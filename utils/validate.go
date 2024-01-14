package utils

import (
	"github.com/go-playground/validator/v10"
)

var validateSingleton *validator.Validate

func intiValidator() *validator.Validate {
	if validateSingleton == nil {
		validateSingleton = validator.New()
	}
	return validateSingleton
}

func Validate[T any](structToValidate T) error {
	validate := intiValidator()
	err := validate.Struct(structToValidate)
	return err
}
