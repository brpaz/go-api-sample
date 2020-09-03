package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	os.Setenv("APP_PORT", "1000")

	err := Load()

	assert.Nil(t, err)
	assert.Equal(t, 1000, Get().Port)

	// Test if it loads defaults correctly
	assert.Equal(t, "prod", Get().Env)
	assert.False(t, Get().Debug)
}
