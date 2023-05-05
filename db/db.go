package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup() (*pgxpool.Pool, error) {
	url := "postgres://postgres:postgres@localhost:5432/postgres"
	db, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return db, nil
}
