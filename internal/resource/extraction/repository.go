package extraction

import (
	"context"

	"github.com/Wave-95/pgserver/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateExtraction(*models.Extraction) error
	DeleteExtraction(extractionId uuid.UUID) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) CreateExtraction(extraction *models.Extraction) error {
	ctx := context.Background()
	sql := "INSERT INTO extractions (id, session_id, result, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(ctx, sql, extraction.Id, extraction.SessionId, extraction.Result, extraction.CreatedAt, extraction.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteExtraction(extractionId uuid.UUID) error {
	ctx := context.Background()
	sql := "DELETE FROM extractions where id = $1"
	_, err := r.db.Exec(ctx, sql, extractionId)
	if err != nil {
		return err
	}
	return nil
}

type MockRepository struct {
	extractions map[uuid.UUID]*models.Extraction
}

func (r *MockRepository) CreateExtraction(extraction *models.Extraction) error {
	r.extractions[extraction.Id] = extraction
	return nil
}

func (r *MockRepository) DeleteExtraction(extractionId uuid.UUID) error {
	delete(r.extractions, extractionId)
	return nil
}
