package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Wave-95/pgserver/api/resource/user"
	"github.com/Wave-95/pgserver/db"
	"github.com/Wave-95/pgserver/internal/config"
	"github.com/Wave-95/pgserver/middleware"
	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/go-chi/chi"
)

func main() {
	cfg := config.New()
	l := logger.New()

	db, err := db.Setup(cfg.DatabaseConfig)
	if err != nil {
		l.Fatalf("Issue connecting to db: %s", err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.RequestLogger(l))

	userApi := user.NewUserApi(db)
	userApi.SetupRoutes(r)

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
