package user

import (
	"context"
	"testing"

	"github.com/Wave-95/pgserver/test"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	db := test.DB(t)
	service := NewUserService(db)
	assert.NotNil(t, service)

	t.Run("GetUser", func(t *testing.T) {
		t.Run("not found", func(t *testing.T) {
			input := GetUserRequest{UserID: "abc123"}
			_, err := service.GetUser(context.Background(), input)
			assert.ErrorIs(t, err, ErrUserNotFound)
		})

	})
}
