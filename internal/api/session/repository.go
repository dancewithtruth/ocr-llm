package session

import (
	"context"
	"errors"

	"github.com/Wave-95/pgserver/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrSessionNotFound = errors.New("session not found")
)

type Repository interface {
	CreateSession(session *models.Session) error

	GetSession(sessionID uuid.UUID) (*models.Session, error)

	DeleteSession(sessionID uuid.UUID) error
}

func New(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

type repository struct {
	db *pgxpool.Pool
}

func (r *repository) CreateSession(session *models.Session) error {
	ctx := context.Background()
	sql := "INSERT INTO sessions (id, ip_address, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(ctx, sql, session.Id, session.IPAddress, session.CreatedAt, session.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetSession(sessionID uuid.UUID) (*models.Session, error) {
	ctx := context.Background()
	sql := "select * from sessions where id = $1"
	session := models.Session{}
	err := r.db.QueryRow(ctx, sql, sessionID).Scan(&session.Id, &session.IPAddress, &session.CreatedAt, &session.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}
	return &session, nil
}

func (r *repository) DeleteSession(sessionID uuid.UUID) error {
	ctx := context.Background()
	sql := "DELETE FROM sessions where id = $1"
	_, err := r.db.Exec(ctx, sql, sessionID)
	if err != nil {
		return err
	}
	return nil
}
