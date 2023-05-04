package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	fmt.Println("Connecting to db")
	conn, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}
	return conn, nil
}
