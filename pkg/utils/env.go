package utils

import (
	"os"

	"github.com/Matt-Gleich/logoru"
)

// SafelyGetenv ... Safely get environment variables failing when it doesn't exist
func SafelyGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		logoru.Error("Failed to get", key, "key")
		os.Exit(1)
	}
	return value
}
