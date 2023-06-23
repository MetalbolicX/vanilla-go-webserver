package routes

import (
	"net/http"

	"github.com/MetalbolicX/vanilla-go-webserver/internal/handlers"
	"github.com/MetalbolicX/vanilla-go-webserver/internal/middlewares"
	"github.com/MetalbolicX/vanilla-go-webserver/internal/pages"
	"github.com/MetalbolicX/vanilla-go-webserver/pkg/server"
)

// The bindRoutes function is responsible for binding
// the routes (URL paths) to their corresponding handler
// functions in the server.Server instance s.
// It sets up the routing configuration for
// various HTTP methods (GET, POST, PUT)
// and associates each route with its respective handler function.
func BindRoutes(s *server.Server) {
	s.Handle(http.MethodGet, "/", pages.HomeHandler)
	s.Handle(http.MethodPost, "/customer", handlers.NewCustomerHandler)
	s.Handle(http.MethodGet, "/customer/\\d+", handlers.GetCustomerByIdHandler)
	s.Handle(http.MethodPut, "/customer/\\d+", handlers.UpdateCustomerHandler)
	s.Handle(http.MethodDelete, "/customer/\\d+",
		s.AddMiddleware(handlers.DeleteCustomerHandler,
			middlewares.CheckAuth(), middlewares.Logging()))
}
