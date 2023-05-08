package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
