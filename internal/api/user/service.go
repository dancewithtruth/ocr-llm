package user

import (
	"context"

	"github.com/Wave-95/pgserver/internal/models"
)

type Service interface {
	GetUser(ctx context.Context, userID string) (*models.User, error)
}

type service struct {
	repository Repository
}

func NewUserService(userRepository Repository) Service {
	return &service{repository: userRepository}
}

func (s *service) GetUser(ctx context.Context, userID string) (*models.User, error) {
	user, err := s.repository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
