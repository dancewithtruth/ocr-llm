package db

import (
	"context"
	"fmt"

	"github.com/Wave-95/pgserver/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	url := buildUrl(cfg)
	db, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func buildUrl(cfg config.DatabaseConfig) string {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	return url
}
