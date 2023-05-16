package extraction

import (
	"io"
	"strings"
	"time"

	"github.com/Wave-95/pgserver/internal/models"
	"github.com/Wave-95/pgserver/pkg/llm"
	"github.com/Wave-95/pgserver/pkg/ocr"
	"github.com/google/uuid"
)

type ExtractedData string

type Service interface {
	CreateExtraction(file io.ReadCloser, targets string, sessionId uuid.UUID) (*models.Extraction, error)
}

type service struct {
	Dependencies
}

type Dependencies struct {
	repository Repository
	ocr        ocr.OCR
	llm        llm.LLM
}

func NewService(deps Dependencies) Service {
	return &service{
		Dependencies: deps,
	}
}

// CreateExtraction takes in an OCR instance, a image file, and targets to extract. It uses the OCR to
// extract the text from the image file and then interacts with an LLM for further processing.
func (s *service) CreateExtraction(file io.ReadCloser, targets string, sessionId uuid.UUID) (*models.Extraction, error) {
	// extract text from image using ocr
	textocr, err := s.ocr.ImageToText(file)

	// process ocr text
	extractMessage := buildExtractMessage(PromptExtract, textocr, targets)
	llmResponse, err := s.llm.Chat(extractMessage)
	if err != nil {
		return nil, err
	}

	// store result as new extraction object
	extraction := &models.Extraction{
		Id:        uuid.New(),
		SessionId: sessionId,
		Result:    llmResponse,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = s.repository.CreateExtraction(extraction)
	if err != nil {
		return nil, err
	}

	return extraction, err
}

func buildExtractMessage(prompt string, textocr string, targets string) string {
	str := strings.Replace(prompt, "{{textocr}}", textocr, 1)
	str = strings.Replace(str, "{{targets}}", targets, 1)
	return str
}
