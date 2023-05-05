package user

import (
	"github.com/Wave-95/pgserver/db/models"
)

type mockUserRepository struct {
	GetUserResponse func() (*models.User, error)
}

func (m *mockUserRepository) GetUser() (*models.User, error) {
	return m.GetUserResponse()
}

func NewMockUserRepository() mockUserRepository {
	return mockUserRepository{}
}
