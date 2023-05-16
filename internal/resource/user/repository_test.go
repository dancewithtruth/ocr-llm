package user

import (
	"context"
	"testing"

	"github.com/Wave-95/pgserver/test"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	db := test.DB(t)
	userRepository := NewRepository(db)
	ctx := context.Background()
	t.Run("GetUser", func(t *testing.T) {
		t.Run("it should return an err if user not found", func(t *testing.T) {
			_, err := userRepository.GetUser(ctx, "abc-123")
			assert.Equal(t, ErrUserNotFound, err)
		})
	})

}
