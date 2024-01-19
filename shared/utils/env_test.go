package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvOrDefault(t *testing.T) {

	t.Run("succesfully retrieves env variable", func(t *testing.T) {

		expectedValue := "test"
		os.Setenv("TEST_ENV", expectedValue)

		actualValue := GetEnvOrDefault("TEST_ENV", "")
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("fallback to default value if env variable is not set", func(t *testing.T) {

		expectedValue := "default"
		actualValue := GetEnvOrDefault("NONE", expectedValue)
		assert.Equal(t, expectedValue, actualValue)
	})
}

func TestGetEnvAsIntOrDefault(t *testing.T) {
	t.Run("succesfully retrieves env variable as int", func(t *testing.T) {

		expectedValue := 10
		os.Setenv("TEST_ENV", "10")

		actualValue := GetEnvAsIntOrDefault("TEST_ENV", 0)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("fallback to default value if env variable is not set", func(t *testing.T) {

		expectedValue := 12
		actualValue := GetEnvAsIntOrDefault("NONE", expectedValue)
		assert.Equal(t, expectedValue, actualValue)
	})

	t.Run("fallback to default value if env variable can't be int", func(t *testing.T) {
		expectedValue := 12
		os.Setenv("TEST_ENV", "Not a number")
		actualValue := GetEnvAsIntOrDefault("TEST_ENV", expectedValue)
		assert.Equal(t, expectedValue, actualValue)
	})

}
