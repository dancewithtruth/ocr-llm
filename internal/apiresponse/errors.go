package apiresponse

import (
	"encoding/json"
	"net/http"

	"github.com/Wave-95/pgserver/pkg/logger"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func RespondError(w http.ResponseWriter, statusCode int, err error, l logger.Logger) {
	w.WriteHeader(statusCode)
	errResponse := ErrorResponse{err.Error()}
	err = json.NewEncoder(w).Encode(errResponse)
	if err != nil {
		l.Errorf("Could not encode error response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
