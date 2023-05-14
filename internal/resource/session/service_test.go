package session

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	mockRepository := NewMockRepository()
	service := NewService(&mockRepository)
	assert.NotNil(t, service)

	t.Run("CreateSession", func(t *testing.T) {
		ipAddress := "127.0.0.1"
		session, err := service.CreateSession(ipAddress)
		assert.Nil(t, err)
		assert.NotEmpty(t, session.Id)
	})
}
