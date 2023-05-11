package user

import (
	"context"
	"testing"

	"github.com/Wave-95/pgserver/db"
	"github.com/Wave-95/pgserver/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	cfg := config.New()
	db, err := db.Setup(cfg.DatabaseConfig)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	userRepository := NewUserRepository(db)

	ctx := context.Background()
	_, err = userRepository.GetUser(ctx, "abc-123")

	assert.Equal(t, ErrUserNotFound, err)
}
