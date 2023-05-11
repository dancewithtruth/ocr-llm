package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	ServerPort     string
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

const (
	defaultServerPort = ":8080"
	defaultDBName     = "pg_db"
	defaultDBUser     = "postgres"
	defaultDBPW       = "postgres"
	defaultDBHost     = "localhost"
	defaultDBPort     = "5432"
)

// New loads and returns a config object
func New() Config {
	dbConfig := DatabaseConfig{
		Host:     getEnv("DB_HOST", defaultDBHost),
		Port:     getEnv("DB_PORT", defaultDBPort),
		Name:     getEnv("DB_NAME", defaultDBName),
		User:     getEnv("DB_USER", defaultDBUser),
		Password: getEnv("DB_PASSWORD", defaultDBPW),
	}

	cfg := Config{ServerPort: defaultServerPort, DatabaseConfig: dbConfig}
	return cfg
}

// getEnv looks up env vars and returns a fallback if not found
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
