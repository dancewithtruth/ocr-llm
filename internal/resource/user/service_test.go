package user

import (
	"context"
	"testing"

	"github.com/Wave-95/pgserver/test"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	db := test.DB(t)
	userRepository := NewUserRepository(db)
	service := NewUserService(userRepository)
	assert.NotNil(t, service)

	t.Run("GetUser", func(t *testing.T) {
		t.Run("not found", func(t *testing.T) {
			_, err := service.GetUser(context.Background(), "abc123")
			assert.ErrorIs(t, err, ErrUserNotFound)
		})

	})
}
