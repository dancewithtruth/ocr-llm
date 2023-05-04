package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Wave-95/pgserver/db"
)

type UserGetter interface {
	GetUser() (*db.User, error)
}

func UserGet(userGetter UserGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := userGetter.GetUser()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(user)
	}
}
