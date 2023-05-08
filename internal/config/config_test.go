package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := New()
	defaultDBConfig := DatabaseConfig{
		Host:     defaultDBHost,
		Port:     defaultDBPort,
		Name:     defaultDBName,
		User:     defaultDBUser,
		Password: defaultDBPW,
	}
	assert.Equal(t, cfg.ServerPort, defaultServerPort, "Config should have default server port of 8080")
	assert.Equal(t, cfg.DatabaseConfig, defaultDBConfig, "Config should have default database config")

	customDBName := "customdbname"
	customDBPW := "secretpassword"
	os.Setenv("DB_NAME", customDBName)
	os.Setenv("DB_PASSWORD", customDBPW)
	customDBConfig := DatabaseConfig{
		Host:     defaultDBHost,
		Port:     defaultDBPort,
		Name:     customDBName,
		User:     defaultDBUser,
		Password: customDBPW,
	}
	cfg = New()
	assert.Equal(t, cfg.DatabaseConfig, customDBConfig, "Config should have database config with custom fields")
}
