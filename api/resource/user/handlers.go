package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/go-chi/chi"
)

var ErrInternalServer = errors.New("Issue getting user.")

func (api *API) handleGetUser(w http.ResponseWriter, r *http.Request) {
	l := logger.FromContext(r.Context())
	//TODO: "Validate uuid"
	userID := chi.URLParam(r, "userID")
	user, err := api.Repo.GetUser(userID)
	if err != nil {
		if err == ErrUserNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, ErrInternalServer.Error(), http.StatusInternalServerError)
		l.Errorf("Issue getting user: %s", err)
		return
	}
	json.NewEncoder(w).Encode(user)
}
