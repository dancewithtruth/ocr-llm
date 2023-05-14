package user

import (
	"context"
	"errors"

	"github.com/Wave-95/pgserver/internal/models"
	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrUserNotFound = errors.New("user not found")

type Repository interface {
	GetUser(ctx context.Context, userID string) (*models.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func (r *userRepository) GetUser(ctx context.Context, userID string) (*models.User, error) {
	l := logger.FromContext(ctx)
	getUserQuery := "select * from users where id = $1"
	user := models.User{}
	err := r.db.QueryRow(ctx, getUserQuery, userID).Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		l.Errorf("repository: error getting user: %s", err)
		if err == pgx.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *pgxpool.Pool) Repository {
	return &userRepository{db: db}
}
