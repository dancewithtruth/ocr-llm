package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) GetUser() (*User, error) {
	user := User{}
	getUserQuery := "select * from users order by id desc limit 1"
	err := repo.db.QueryRow(context.Background(), getUserQuery).Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
