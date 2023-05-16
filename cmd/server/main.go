package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Wave-95/pgserver/db"
	"github.com/Wave-95/pgserver/internal/config"
	"github.com/Wave-95/pgserver/internal/resource/extraction"
	"github.com/Wave-95/pgserver/internal/resource/session"
	"github.com/Wave-95/pgserver/internal/resource/user"
	"github.com/Wave-95/pgserver/middleware"
	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/Wave-95/pgserver/pkg/validator"
	"github.com/go-chi/chi"
	middlewarechi "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.New()
	l := logger.New()
	v := validator.New()

	db, err := db.Setup(cfg.DatabaseConfig)
	if err != nil {
		l.Fatalf("Issue connecting to db: %s", err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.RequestLogger(l))
	r.Use(middlewarechi.RealIP)
	r.Use(middleware.Session(createSessionService(db)))

	userApi := user.NewAPI(db, v)
	userApi.RegisterHandlers(r)

	extractionHandler := extraction.NewHandler(db, v)
	extractionHandler.RegisterHandlers(r)

	server := &http.Server{
		Addr:    cfg.ServerPort,
		Handler: r,
	}

	// Set up signal channel and send interrupts signals to channel
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatalf("Could not start server: %s", err)
		}
	}()

	// Block channel until interrupt signal is received
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		l.Fatalf("Server forced shutdown: %s", err)
	}
}

func createSessionService(db *pgxpool.Pool) session.Service {
	sessionRepository := session.NewRepository(db)
	return session.NewService(sessionRepository)
}
