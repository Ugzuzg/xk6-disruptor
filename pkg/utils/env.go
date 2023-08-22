package utils

import (
	"os"
	"strconv"
)

// GetBooleanEnvVar returns a boolean environment variable.
// If variable is not set or invalid value, returns the default value
func GetBooleanEnvVar(envVar string, defaultValue bool) bool {
	value, err := strconv.ParseBool(os.Getenv(envVar))
	if err != nil {
		return defaultValue
	}
	return value
}

// GetStringEnvVar returns a string environment variable.
// If variable is not set returns the default value
func GetStringEnvVar(envVar string, defaultValue string) string {
	name := os.Getenv(envVar)
	if name == "" {
		return defaultValue
	}
	return name
}
