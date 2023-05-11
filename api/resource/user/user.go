package user

import (
	"github.com/Wave-95/pgserver/pkg/validator"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

type API struct {
	//Dependencies stored here
	Repo      Repository
	Validator *validator.Validate
}

func NewUserApi(db *pgxpool.Pool, v *validator.Validate) *API {
	userRepository := NewUserRepository(db)
	return &API{Repo: &userRepository, Validator: v}
}

func (api *API) SetupRoutes(r chi.Router) {
	r.Get("/users/{userID}", api.handleGetUser)
}
