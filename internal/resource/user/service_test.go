package user

import (
	"context"
	"testing"

	"github.com/Wave-95/pgserver/test"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	db := test.DB(t)
	userRepository := NewRepository(db)
	userService := NewService(userRepository)
	assert.NotNil(t, userService)

	t.Run("GetUser", func(t *testing.T) {
		t.Run("not found", func(t *testing.T) {
			_, err := userService.GetUser(context.Background(), "abc123")
			assert.ErrorIs(t, err, ErrUserNotFound)
		})

	})
}
