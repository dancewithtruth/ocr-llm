package user

import (
	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

type API struct {
	//Dependencies stored here
	Repo Repository
}

func NewUserApi(db *pgxpool.Pool, l logger.Logger) *API {
	userRepository := NewUserRepository(db, l)
	return &API{Repo: &userRepository}
}

func (api *API) SetupRoutes(r chi.Router) {
	r.Get("/users", api.handleGetUser)
}
