package user

import (
	"context"

	"github.com/Wave-95/pgserver/db/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetUser() (*models.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func (r *userRepository) GetUser() (*models.User, error) {
	getUserQuery := "select * from users order by id desc limit 1"
	user := models.User{}
	err := r.db.QueryRow(context.Background(), getUserQuery).Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *pgxpool.Pool) userRepository {
	return userRepository{db: db}
}
