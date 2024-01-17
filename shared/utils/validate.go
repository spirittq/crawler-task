package utils

import (
	"github.com/go-playground/validator/v10"
)

var validateSingleton *validator.Validate

func initValidator() *validator.Validate {
	if validateSingleton == nil {
		validateSingleton = validator.New()
	}
	return validateSingleton
}

func Validate[T any](structToValidate T) error {
	validate := initValidator()
	err := validate.Struct(structToValidate)
	return err
}
