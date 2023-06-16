package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MetalbolicX/vanilla-go-webserver/handlers"
	"github.com/MetalbolicX/vanilla-go-webserver/server"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("SERVER_PORT")
	DB_URL := os.Getenv("DB_URL")
	DB_MANAGEMENT_SYSTEM := os.Getenv("DB_MANAGEMENT_SYSTEM")
	STATIC_FOLDER := os.Getenv("STATIC_FOLDER")

	server := server.NewServer(PORT, STATIC_FOLDER)
	//
	bindRoutes(server)
	if err := server.Listen(DB_MANAGEMENT_SYSTEM, DB_URL); err != nil {
		log.Fatal("Server cannot be started")
	}

}

func bindRoutes(s *server.Server) {
	s.Handle(http.MethodGet, "/", handlers.IndexHandler)
	s.Handle(http.MethodGet, "/exercises", handlers.GetExercisesHandler)
	s.Handle(http.MethodGet, "/home", handlers.HomeHandler)
	s.Handle(http.MethodPost, "/user", handlers.NewCustomerHandler)
}
