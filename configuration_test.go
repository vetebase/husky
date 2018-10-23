package husky

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigurationLoad(t *testing.T) {
	h := New()

	config := h.Config.Load()

	assert.True(t, reflect.TypeOf(config).String() == "map[string]string")
}

func TestConfigurationValues(t *testing.T) {
	h := New()

	config := h.Config.Load()

	assert.True(t, config["NAME"] == "husky")
	assert.True(t, config["PORT"] == "8080")
	assert.True(t, config["JWT_SECRET"] == "12345")
}
