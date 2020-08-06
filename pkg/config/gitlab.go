package config

import (
	"os"

	"github.com/Matt-Gleich/eagleye/pkg/utils"
	"github.com/sirupsen/logrus"
)

type GitLabConfig struct {
	Token   string
	BaseURL string
}

func LoadGitLabConfig() GitLabConfig {
	logrus.Info("Loading configuration from env")
	token := safelyGetenv("EAGLEYE_TOKEN")
	baseURL := safelyGetenv("EAGLEYE_BASE_URL")
	utils.Success("Loaded configuration!")
	return GitLabConfig{token, baseURL}
}

func safelyGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		logrus.Fatal("Failed to get " + key + " key")
	}
	return value
}
