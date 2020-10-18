package util

import (
	"os"
	"strconv"
)

// GetBoolEnv returns envrionment variable as boolean value
func GetBoolEnv(name string) (bool, error) {
	value := os.Getenv(name)

	if value == "" {
		return false, nil
	}

	return strconv.ParseBool(value)
}
