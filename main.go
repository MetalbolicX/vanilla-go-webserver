package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MetalbolicX/vanilla-go-webserver/handlers"
	"github.com/MetalbolicX/vanilla-go-webserver/middlewares"
	"github.com/MetalbolicX/vanilla-go-webserver/server"
	"github.com/MetalbolicX/vanilla-go-webserver/utils"
)

func main() {

	if err := utils.LoaderEnvFile(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("SERVER_PORT")
	DB_URL := os.Getenv("DB_URL")
	DB_MANAGEMENT_SYSTEM := os.Getenv("DB_MANAGEMENT_SYSTEM")
	STATIC_FOLDER := os.Getenv("STATIC_FOLDER")

	server := server.NewServer(PORT, STATIC_FOLDER)
	bindRoutes(server)
	if err := server.Listen(DB_MANAGEMENT_SYSTEM, DB_URL); err != nil {
		log.Fatal("Server cannot be started")
	}

}

// The bindRoutes function is responsible for binding
// the routes (URL paths) to their corresponding handler
// functions in the server.Server instance s.
// It sets up the routing configuration for
// various HTTP methods (GET, POST, PUT)
// and associates each route with its respective handler function.
func bindRoutes(s *server.Server) {
	s.Handle(http.MethodGet, "/", handlers.IndexHandler)
	s.Handle(http.MethodGet, "/exercises", handlers.GetExercisesHandler)
	s.Handle(http.MethodGet, "/home", handlers.HomeHandler)
	s.Handle(http.MethodPost, "/customer", handlers.NewCustomerHandler)
	s.Handle(http.MethodGet, "/customer/\\d+", handlers.GetCustomerByIdHandler)
	s.Handle(http.MethodPut, "/customer/\\d+", handlers.UpdateCustomerHandler)
	s.Handle(http.MethodDelete, "/customer/\\d+",
		s.AddMiddleware(handlers.DeleteCustomerHandler,
			middlewares.CheckAuth(), middlewares.Logging()))
}
