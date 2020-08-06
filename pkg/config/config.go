package config

import (
	"strings"

	"github.com/Matt-Gleich/eagleye/pkg/utils"
	"github.com/sirupsen/logrus"
)

type GitLabConfig struct {
	Token   string
	BaseURL string
}

func LoadConfiguration() (string, GitLabConfig) {
	logrus.Info("Loading configuration from env")
	platform := utils.SafelyGetenv("EAGLEYE_PLATFORM")
	var platformName string
	var gitlabConfig GitLabConfig
	switch strings.ToLower(platform) {
	case "gitlab":
		platformName = "gitlab"
		gitlabConfig = loadGitLabConfig()
	case "github":
		logrus.Fatal("GitHub is currently not supported")
	default:
		logrus.Fatal(platform + " is not supported")
	}
	utils.Success("Loaded configuration!")
	return platformName, gitlabConfig
}

func loadGitLabConfig() GitLabConfig {
	token := utils.SafelyGetenv("EAGLEYE_GITLAB_TOKEN")
	baseURL := utils.SafelyGetenv("EAGLEYE_GITLAB_BASE_URL")
	return GitLabConfig{token, baseURL}
}
