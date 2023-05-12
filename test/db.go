package test

import (
	"testing"

	"github.com/Wave-95/pgserver/db"
	"github.com/Wave-95/pgserver/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DB(t *testing.T) *pgxpool.Pool {
	cfg := config.New()
	db, err := db.Setup(cfg.DatabaseConfig)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	return db
}
