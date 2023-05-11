package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func (api *API) handleGetUser(w http.ResponseWriter, r *http.Request) {
	//TODO: "Validate uuid"
	userID := chi.URLParam(r, "userID")
	user, err := api.Repo.GetUser(userID)
	if err != nil {
		if err == ErrUserNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
