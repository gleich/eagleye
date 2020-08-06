package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfiguration(t *testing.T) {
	platformKey := "EAGLEYE_PLATFORM"
	os.Setenv(platformKey, "gitlab")
	defer os.Setenv(platformKey, "")
	setGitLabEnvVars()
	defer tearDownGitLabEnvVars()

	platformName1, _ := LoadConfiguration()
	assert.Equal(t, platformName1, "gitlab")
}

func TestLoadGitLabConfig(t *testing.T) {
	// Setting env vars
	tokenValue, baseURLValue := setGitLabEnvVars()
	defer tearDownGitLabEnvVars()

	gitlabConfig := loadGitLabConfig()
	assert.Equal(t, gitlabConfig, GitLabConfig{tokenValue, baseURLValue})
}

func setGitLabEnvVars() (string, string) {
	tokenKey := "EAGLEYE_GITLAB_TOKEN"
	baseURLKey := "EAGLEYE_GITLAB_BASE_URL"
	tokenValue := "1234"
	baseURLValue := "https://git.mydomain.com/api/v4"
	os.Setenv(tokenKey, tokenValue)
	os.Setenv(baseURLKey, baseURLValue)
	return tokenValue, baseURLValue
}

func tearDownGitLabEnvVars() {
	os.Setenv("EAGLEYE_GITLAB_TOKEN", "")
	os.Setenv("EAGLEYE_GITLAB_BASE_URL", "")
}
