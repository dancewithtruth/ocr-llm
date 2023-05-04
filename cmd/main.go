package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Wave-95/pgserver/db"
	"github.com/Wave-95/pgserver/handlers"
	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Initiating server")

	//Connect to DB
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("Issue connecting to db: %v\n", err)
	}
	defer conn.Close()

	userRepo := db.NewUserRepo(conn)
	r := chi.NewRouter()

	r.Get("/user", handlers.UserGet(userRepo))

	http.ListenAndServe(":8080", r)
}
