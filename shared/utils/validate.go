package utils

import (
	"github.com/go-playground/validator/v10"
)

var validateSingleton *validator.Validate

// Initiate validator if it is not initialized yet
func initValidator() *validator.Validate {
	if validateSingleton == nil {
		validateSingleton = validator.New()
	}
	return validateSingleton
}

// Validate provided type data
func Validate[T any](structToValidate T) error {
	validate := initValidator()
	err := validate.Struct(structToValidate)
	return err
}
