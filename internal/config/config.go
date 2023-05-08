package config

import (
	"fmt"
	"os"
)

type Config struct {
	ServerPort string
	DSN        string
}

const (
	defaultServerPort = "8080"
	defaultDBname     = "postgres"
	defaultDBuser     = "postgres"
	defaultDBpw       = "postgres"
	defaultDBPort     = "5432"
)

func New() Config {
	dbuser := os.Getenv("DB_USER")
	dbpw := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	dsn := buildDBUrl(dbuser, dbpw, dbname, dbport)

	cfg := Config{ServerPort: defaultServerPort, DSN: dsn}
	return cfg
}

func buildDBUrl(dbuser, dbpw, dbname, dbport string) string {
	if dbuser == "" {
		dbuser = defaultDBuser
	}
	if dbpw == "" {
		dbpw = defaultDBpw
	}
	if dbname == "" {
		dbname = defaultDBname
	}
	if dbport == "" {
		dbport = defaultDBPort
	}
	return fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", dbuser, dbpw, dbport, dbname)
}
