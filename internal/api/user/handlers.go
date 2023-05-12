package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Wave-95/pgserver/internal/apiresponse"
	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/Wave-95/pgserver/pkg/validator"
	"github.com/go-chi/chi"
)

var (
	ErrInternalServer        = errors.New("Internal server error")
	ErrGetUserInvalidRequest = errors.New("Invalid get user request")
	ErrGetUserEncodeJSON     = errors.New("Error encoding user to JSON")
)

type GetUserRequest struct {
	UserID string `validate:"required,uuid4"`
}

func (r GetUserRequest) Validate(v validator.Validate) error {
	return v.Struct(r)
}

func (api *API) handleGetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := logger.FromContext(ctx)
	// Validate get user request
	userID := chi.URLParam(r, "userID")
	input := GetUserRequest{UserID: userID}
	if err := input.Validate(api.validate); err != nil {
		apiresponse.RespondError(w, http.StatusBadRequest, ErrGetUserInvalidRequest)
		return
	}

	// Get user and handle errors
	user, err := api.service.GetUser(ctx, input)
	if err != nil {
		switch err {
		case ErrUserNotFound:
			apiresponse.RespondError(w, http.StatusNotFound, ErrUserNotFound)
		default:
			apiresponse.RespondError(w, http.StatusInternalServerError, ErrInternalServer)
			l.Errorf("Issue getting user: %s", err)
		}
		return
	}

	// Write user response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		apiresponse.RespondError(w, http.StatusInternalServerError, ErrGetUserEncodeJSON)
	}
}
