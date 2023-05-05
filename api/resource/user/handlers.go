package user

import (
	"encoding/json"
	"net/http"
)

func (api *API) handleGetUser(w http.ResponseWriter, r *http.Request) {
	user, err := api.Repo.GetUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(user)
}
