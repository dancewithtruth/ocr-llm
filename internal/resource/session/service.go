package session

import (
	"time"

	"github.com/Wave-95/pgserver/internal/models"
	"github.com/google/uuid"
)

type Service interface {
	CreateSession(ipAddress string) (*models.Session, error)
	GetSession(sessionID uuid.UUID) (*models.Session, error)
}

type service struct {
	repository Repository
}

func (s *service) CreateSession(ipAddress string) (*models.Session, error) {
	session := generateSession(ipAddress)
	err := s.repository.CreateSession(session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *service) GetSession(sessionID uuid.UUID) (*models.Session, error) {
	session, err := s.repository.GetSession(sessionID)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

// generateSession is a helper function that takes an IPAddress and returns a session model
func generateSession(ipAddress string) *models.Session {
	return &models.Session{
		Id:        uuid.New(),
		IPAddress: ipAddress,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
