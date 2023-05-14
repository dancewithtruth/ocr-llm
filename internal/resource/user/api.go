package user

import (
	"github.com/Wave-95/pgserver/pkg/validator"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

type API struct {
	service  Service
	validate validator.Validate
}

func NewAPI(db *pgxpool.Pool, v validator.Validate) *API {
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	return &API{service: userService, validate: v}
}

func (api *API) RegisterHandlers(r chi.Router) {
	r.Get("/users/{userID}", api.handleGetUser)
}
