package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	ValidateField *string `validate:"required"`
}

func TestValidate(t *testing.T) {

	t.Run("successfull validation", func(t *testing.T) {
		field := "test"

		testStruct := TestStruct{
			ValidateField: &field,
		}
		err := Validate[TestStruct](testStruct)
		assert.Nil(t, err)
	})

	t.Run("fail validation return error", func(t *testing.T) {
		testStruct := TestStruct{}
		err := Validate[TestStruct](testStruct)
		assert.Error(t, err)
	})

}
