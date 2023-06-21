package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MetalbolicX/vanilla-go-webserver/database"
	"github.com/MetalbolicX/vanilla-go-webserver/repository"
	"github.com/MetalbolicX/vanilla-go-webserver/types"
)

// The Server struct represents the server configuration.
// It has fields for the listening port, a router instance
// and the static file folder path.
type Server struct {
	port         string
	router       *router
	staticFolder string
}

// The NewServer function creates a new instance of
// the Server. It takes the port and static folder path
// as parameters and initializes the server with
// the provided values. It also creates a new router
// using the NewRouter function.
func NewServer(port, staticFolder string) *Server {
	return &Server{
		port:         port,
		router:       NewRouter(),
		staticFolder: staticFolder,
	}
}

// The String method provides a string representation
// of the server, indicating the selected port for
// listening.
func (s *Server) String() string {
	return fmt.Sprintf("Port listening selected is %s", s.port)
}

// The Handle method allows you to define a routing rule
// for the server. It takes an HTTP method, URL path
// and a handler function. It checks if the method
// already exists in the router's rules. If not,
// it creates a new map to store the rules for that method.
// It then assigns the provided handler function to
// the specified method and path.
func (s *Server) Handle(method, path string, handlerLogic http.HandlerFunc) {
	if _, methodExists := s.router.rules[method]; !methodExists {
		s.router.rules[method] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[method][path] = handlerLogic
}

// The Listen method starts the server and listens for
// incoming requests. It takes the database management
// system (e.g., MySQL, PostgreSQL) and the database URL
// as parameters. It creates a new repository using
// the provided parameters and sets it as the
// implementation for the repository using repository.
// SetRepository. It registers the router with the
// root path ("/") as the default handler for all requests.
// It then calls s.setupStaticFileServer to configure
// the server to serve static files if a static folder
// is specified. Finally, it starts the server by calling
// http.ListenAndServe with the specified port
// and it logs the server's listening port.
func (s *Server) Listen(dbManagmentSystem, dbUrl string) error {

	repo, err := database.NewRelationalDBRepo(dbManagmentSystem, dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)

	http.Handle("/", s.router)
	s.setupStaticFileServer()

	log.Println(s.String())
	if err := http.ListenAndServe(s.port, nil); err != nil {
		return err
	}

	return nil
}

// The setupStaticFileServer method configures the
// server to serve static files if a static folder is
// specified. It creates a file server using
// http.FileServer with the provided static folder path.
// It then registers a handler for the folder path,
// stripping the folder prefix from the URL path
// before serving the static files.
func (s *Server) setupStaticFileServer() {
	if s.staticFolder != "" {
		fs := http.FileServer(http.Dir(fmt.Sprintf("./%s", s.staticFolder)))
		folder := fmt.Sprintf("/%s/", s.staticFolder)
		http.Handle(folder, http.StripPrefix(folder, fs))
	}
}

// The function applies the provided middlewares to the
// handler logic in a sequential manner and returns the
// resulting handler function is then used for routing
// or serving HTTP requests. When a request is received,
// the middlewares will be executed in the order they
// were added, allowing you to perform additional
// operations before and after calling the original
// handler logic.
func (s *Server) AddMiddleware(handlerLogic http.HandlerFunc, middlewares ...types.Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handlerLogic = middleware(handlerLogic)
	}
	return handlerLogic
}
