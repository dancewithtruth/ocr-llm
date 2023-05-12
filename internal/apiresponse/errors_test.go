package apiresponse

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondError(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RespondError(w, 500, ErrInternalServer)
	})

	req := httptest.NewRequest("Get", "/", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	var response ErrResponse
	json.NewDecoder(rec.Body).Decode(&response)

	gotMsg := response.Message
	if gotMsg != ErrInternalServer.Error() {
		t.Errorf("got %q, want %q", gotMsg, ErrInternalServer.Error())
	}

	gotStatusCode := rec.Result().StatusCode
	if gotStatusCode != 500 {
		t.Errorf("got %d, want %d", gotStatusCode, 500)
	}
}
