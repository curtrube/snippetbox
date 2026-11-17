package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad_Defaults(t *testing.T) {
	cfg, err := Load()

	require.NoError(t, err)

	assert.Equal(t, cfg.Environment, "dev")
	assert.Equal(t, cfg.MySQL.Host, "localhost")
	assert.Equal(t, cfg.MySQL.Port, 3306)
	assert.Equal(t, cfg.MySQL.Database, "snippetbox")
	assert.Equal(t, cfg.MySQL.Username, "web")
	assert.Equal(t, cfg.MySQL.Password, "pass")
}
