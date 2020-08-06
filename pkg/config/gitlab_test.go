package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadGitLabConfig(t *testing.T) {
	// Setting env vars
	tokenKey := "EAGLEYE_TOKEN"
	baseURLKey := "EAGLEYE_BASE_URL"
	tokenValue := "1234"
	baseURLValue := "https://git.mydomain.com/api/v4"
	os.Setenv(tokenKey, tokenValue)
	os.Setenv(baseURLKey, baseURLValue)
	defer os.Setenv(tokenKey, "")
	defer os.Setenv(baseURLKey, "")

	gitlabConfig := LoadGitLabConfig()
	assert.Equal(t, gitlabConfig, GitLabConfig{tokenValue, baseURLValue})
}

func TestSafelyGetenv(t *testing.T) {
	// Setting env vars
	key := "EAGLEYE_TESTING_TESTING"
	os.Setenv(key, "1234")
	defer os.Setenv(key, "")

	value := safelyGetenv(key)
	assert.Equal(t, value, "1234", "They should be equal")
}
