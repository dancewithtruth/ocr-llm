package user

import (
	"context"

	"github.com/Wave-95/pgserver/db/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	GetUser(ctx context.Context, input GetUserRequest) (*models.User, error)
}

type service struct {
	repository Repository
}

func NewUserService(db *pgxpool.Pool) Service {
	userRepository := NewUserRepository(db)
	return &service{repository: userRepository}
}

func (s *service) GetUser(ctx context.Context, input GetUserRequest) (*models.User, error) {
	user, err := s.repository.GetUser(ctx, input.UserID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
