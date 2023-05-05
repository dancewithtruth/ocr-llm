package main

import (
	"log"
	"net/http"

	"github.com/Wave-95/pgserver/api/resource/user"
	"github.com/Wave-95/pgserver/db"
	"github.com/Wave-95/pgserver/middleware/logMiddleware"
	"github.com/Wave-95/pgserver/middleware/requestid"
	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/go-chi/chi"
)

func main() {
	l := logger.New()

	// Initializes database
	db, err := db.Setup()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(requestid.Middleware())
	r.Use(logMiddleware.Middleware(l))

	userApi := user.NewUserApi(db, l)
	userApi.SetupRoutes(r)

	http.ListenAndServe(":8080", r)
}
