package session

import (
	"testing"
	"time"

	"github.com/Wave-95/pgserver/internal/models"
	"github.com/Wave-95/pgserver/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	db := test.DB(t)
	r := NewRepository(db)
	assert.NotNil(t, r)
	sessionID := uuid.New()
	ipAddress := "127.0.0.1"
	t.Run("CreateSession", func(t *testing.T) {
		session := &models.Session{
			Id:        sessionID,
			IPAddress: ipAddress,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		err := r.CreateSession(session)
		assert.Nil(t, err)
	})

	t.Run("GetSession", func(t *testing.T) {
		t.Run("session exists", func(t *testing.T) {
			session, err := r.GetSession(sessionID)
			if err != nil {
				t.Fatalf("got %q, want nil", err)
			}
			assert.Equal(t, session.Id, sessionID)
		})

		t.Run("session not found", func(t *testing.T) {
			randomId := uuid.New()
			_, err := r.GetSession(randomId)
			assert.ErrorIs(t, err, ErrSessionNotFound)
		})

	})

	t.Run("DeleteSession", func(t *testing.T) {
		err := r.DeleteSession(sessionID)
		assert.Nil(t, err)
		_, err = r.GetSession(sessionID)
		assert.ErrorIs(t, err, ErrSessionNotFound)
	})
}
