package user

import (
	"context"
	"time"

	"github.com/Wave-95/pgserver/db/models"
	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetUser() (*models.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
	l  logger.Logger
}

func (r *userRepository) GetUser() (*models.User, error) {
	getUserQuery := "select * from users order by id desc limit 1"
	user := models.User{}
	err := r.db.QueryRow(context.Background(), getUserQuery).Scan(&user.Id, &user.FirstName, &user.LastName)
	r.l.Info("Getting User")
	time.Sleep(2 * time.Second)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *pgxpool.Pool, l logger.Logger) userRepository {
	return userRepository{db: db, l: l}
}
