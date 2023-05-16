package apiresponse

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInternalServer = errors.New("Internal Server Error")
)

type ErrResponse struct {
	Message string `json:"message"`
}

func RespondWithError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	errResponse := ErrResponse{err.Error()}
	json.NewEncoder(w).Encode(errResponse)
}
