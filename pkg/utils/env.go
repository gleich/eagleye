package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// SafelyGetenv ... Safely get environment variables failing when it doesn't exist
func SafelyGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		logrus.Fatal("Failed to get " + key + " key")
	}
	return value
}
