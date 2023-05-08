package main

import (
	"net/http"

	"github.com/Wave-95/pgserver/api/resource/user"
	"github.com/Wave-95/pgserver/db"
	"github.com/Wave-95/pgserver/internal/config"
	"github.com/Wave-95/pgserver/middleware/logMiddleware"
	"github.com/Wave-95/pgserver/middleware/request"
	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/go-chi/chi"
)

func main() {
	cfg := config.New()
	l := logger.New()

	// Initializes database
	db, err := db.Setup(cfg.DSN)
	if err != nil {
		l.Fatalf("Issue connecting to db: %s", err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(request.Middleware())
	r.Use(logMiddleware.Middleware(l))

	userApi := user.NewUserApi(db, l)
	userApi.SetupRoutes(r)

	http.ListenAndServe(":8080", r)
}
