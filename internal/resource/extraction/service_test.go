package extraction

import (
	"mime/multipart"
	"testing"

	"github.com/Wave-95/pgserver/internal/models"
	"github.com/Wave-95/pgserver/pkg/llm"
	"github.com/Wave-95/pgserver/pkg/ocr"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	deps := Dependencies{
		repository: &MockRepository{extractions: make(map[uuid.UUID]*models.Extraction)},
		ocr:        &ocr.MockOCR{},
		llm:        &llm.MockLLM{},
	}
	service := NewService(deps)
	assert.NotNil(t, service)

	t.Run("CreateExtraction", func(t *testing.T) {
		sessionId := uuid.New()
		var mockFile multipart.File
		targets := "xyz"
		extraction, err := service.CreateExtraction(mockFile, targets, sessionId)
		assert.Nil(t, err)
		assert.NotNil(t, extraction)
	})
}
