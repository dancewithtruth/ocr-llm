package extraction

import (
	"testing"
	"time"

	"github.com/Wave-95/pgserver/internal/models"
	"github.com/Wave-95/pgserver/internal/resource/session"
	"github.com/Wave-95/pgserver/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	db := test.DB(t)
	repository := NewRepository(db)
	sessionRepository := session.NewRepository(db)
	assert.NotNil(t, repository)
	extractionId := uuid.New()
	sessionId := uuid.New()

	// Set up test by creating session
	err := sessionRepository.CreateSession(&models.Session{
		Id:        sessionId,
		IPAddress: "127.0.0.1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		t.Fatalf("could not set up session properly: %q", err)
	}

	t.Run("CreateExtraction", func(t *testing.T) {
		extraction := models.Extraction{
			Id:        extractionId,
			SessionId: sessionId,
			Result:    "abc123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		err := repository.CreateExtraction(&extraction)
		assert.Nil(t, err)
	})

	t.Run("DeleteExtraction", func(t *testing.T) {
		err := repository.DeleteExtraction(extractionId)
		assert.Nil(t, err)
	})

	// clean up session
	sessionRepository.DeleteSession(sessionId)
}
