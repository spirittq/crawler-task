package utils

import (
	"os"
	"strconv"
)

func GetEnvOrDefault(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetEnvAsIntOrDefault(key string, fallback int) int {
	value := os.Getenv(key)
	intValue, err := strconv.Atoi(value)

	if err != nil {
		return fallback
	}
	return intValue
}
