package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := New()
	assert.Equal(t, cfg.ServerPort, defaultServerPort, "Config should have default server port of 8080")
	assert.Equal(t, cfg.DSN, buildDBUrl(defaultDBuser, defaultDBpw, defaultDBname, defaultDBPort), "Config should have default DSN")

	os.Setenv("DB_NAME", "PG_DB")
	cfg = New()
	assert.Equal(t, cfg.DSN, buildDBUrl(defaultDBuser, defaultDBpw, "PG_DB", defaultDBPort), "Config should have DSN with db name from env var")
}
