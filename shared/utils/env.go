package utils

import (
	"os"
	"strconv"
)

// Extract env variable based on the key or fallback to provided default value if key doesn't exist
func GetEnvOrDefault(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}
	return value
}

// Extract env variable as integer based on the key or fallback to provided default value if key doesn't exist
func GetEnvAsIntOrDefault(key string, fallback int) int {
	value := os.Getenv(key)
	intValue, err := strconv.Atoi(value)

	if err != nil {
		return fallback
	}
	return intValue
}
