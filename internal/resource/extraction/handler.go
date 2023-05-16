package extraction

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Wave-95/pgserver/internal/apiresponse"
	"github.com/Wave-95/pgserver/internal/models"
	"github.com/Wave-95/pgserver/middleware"
	"github.com/Wave-95/pgserver/pkg/llm"
	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/Wave-95/pgserver/pkg/ocr"
	"github.com/Wave-95/pgserver/pkg/validator"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrCreateExtractionEncodeJSON = errors.New("Issue encoding response.")

type Handler struct {
	db *pgxpool.Pool
	v  validator.Validate
}

func NewHandler(db *pgxpool.Pool, validate validator.Validate) Handler {
	return Handler{db: db, v: validate}
}

type CreateExtractionRequest struct {
	File      io.ReadCloser `form:"file" validate:"required,uuid4`
	Targets   string        `form:"targets" validate:"required`
	LLMAPIKey string        `form:"llm_api_key"`
}

func (r *CreateExtractionRequest) Validate(v validator.Validate) error {
	return v.Struct(r)
}

type CreateExtractionResponse struct {
	Id        uuid.UUID `json:"id"`
	Result    string    `json:"result"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Create handles create extraction requests
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)

	// Parse multipart form (max 32mb file)
	req, err := parseAndValidateForm(r, h.v)
	if err != nil {
		logger.Errorf("Issue parsing or validating form: %q", err)
		apiresponse.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Get session ID from cookies
	sessionId, err := getSessionId(r)
	if err != nil {
		logger.Errorf("Issue getting session ID: %q", err)
		apiresponse.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	// Setup dependencies and create extraction service
	deps := Dependencies{
		repository: NewRepository(h.db),
		ocr:        ocr.New(),
		llm:        llm.New(req.LLMAPIKey),
	}
	defer deps.ocr.Close()
	extractionService := NewService(deps)
	extraction, err := extractionService.CreateExtraction(req.File, req.Targets, sessionId)
	if err != nil {
		logger.Errorf("Issue creating extraction: %q", err)
		apiresponse.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	writeCreateExtraction(w, extraction)
}

func parseAndValidateForm(r *http.Request, v validator.Validate) (*CreateExtractionRequest, error) {
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("Issue parsing file from request form: %w", err)
	}
	defer file.Close()
	targets := r.FormValue("targets")
	llmAPIKey := r.FormValue("llm_api_key")

	// Validate request
	req := CreateExtractionRequest{
		File:      file,
		Targets:   targets,
		LLMAPIKey: llmAPIKey,
	}
	if err := req.Validate(v); err != nil {
		return nil, fmt.Errorf("Issue validating request: %w", err)
	}
	return &req, nil
}

func getSessionId(r *http.Request) (uuid.UUID, error) {
	sessionCookie, err := r.Cookie(middleware.CookieNameSession)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Issue getting session cookie from request: %w", err)
	}
	sessionId, err := uuid.Parse(sessionCookie.Value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Issue parsing session cookie: %w", err)
	}
	return sessionId, nil
}

func writeCreateExtraction(w http.ResponseWriter, extraction *models.Extraction) {
	// Write extraction response
	res := CreateExtractionResponse{
		Id:        extraction.Id,
		Result:    extraction.Result,
		CreatedAt: extraction.CreatedAt,
		UpdatedAt: extraction.UpdatedAt,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		apiresponse.RespondWithError(w, http.StatusInternalServerError, ErrCreateExtractionEncodeJSON)
	}
}

func (h *Handler) RegisterHandlers(r chi.Router) {
	r.Post("/extractions", h.Create)
}
