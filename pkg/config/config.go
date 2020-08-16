package config

import (
	"os"
	"strings"

	"github.com/Matt-Gleich/eagleye/pkg/utils"
	"github.com/Matt-Gleich/logoru"
)

type GitLabConfig struct {
	Token   string
	BaseURL string
}

func LoadConfiguration() (string, GitLabConfig) {
	logoru.Info("Loading configuration from env")
	platform := utils.SafelyGetenv("EAGLEYE_PLATFORM")
	var platformName string
	var gitlabConfig GitLabConfig
	switch strings.ToLower(platform) {
	case "gitlab":
		platformName = "gitlab"
		gitlabConfig = loadGitLabConfig()
	case "github":
		logoru.Error("GitHub is currently not supported")
		os.Exit(1)
	default:
		logoru.Error("Platform: \"", platform, "\" is not supported")
		os.Exit(1)
	}
	logoru.Success("Loaded configuration!")
	return platformName, gitlabConfig
}

func loadGitLabConfig() GitLabConfig {
	token := utils.SafelyGetenv("EAGLEYE_GITLAB_TOKEN")
	baseURL := utils.SafelyGetenv("EAGLEYE_GITLAB_BASE_URL")
	return GitLabConfig{token, baseURL}
}
