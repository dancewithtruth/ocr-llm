package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Wave-95/pgserver/db"
	"github.com/Wave-95/pgserver/handlers"
)

type MockUserGetter struct{}

func (ug *MockUserGetter) GetUser() (*db.User, error) {
	return &db.User{Id: 1, FirstName: "Victor", LastName: "Wu"}, nil
}

func TestGetUserHandler(t *testing.T) {
	mockUserGetter := MockUserGetter{}
	req, err := http.NewRequest(http.MethodGet, "/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	handler := handlers.UserGet(&mockUserGetter)
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var user db.User
	err = json.NewDecoder(rec.Body).Decode(&user)
	if err != nil {
		t.Fatal(err)
	}

	if user.Id != 1 {
		t.Errorf("Expected ID to be 1, but got %d", user.Id)
	}
}
