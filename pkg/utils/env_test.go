package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafelyGetenv(t *testing.T) {
	// Setting env vars
	key := "EAGLEYE_TESTING_TESTING"
	os.Setenv(key, "1234")
	defer os.Setenv(key, "")

	value := SafelyGetenv(key)
	assert.Equal(t, value, "1234", "They should be equal")
}
