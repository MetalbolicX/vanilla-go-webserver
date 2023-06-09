package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MetalbolicX/vanilla-go-webserver/database"
	"github.com/MetalbolicX/vanilla-go-webserver/repository"
)

type Server struct {
	port         string
	router       *router
	staticFolder string
}

func NewServer(port, staticFolder string) *Server {
	return &Server{
		port:         port,
		router:       NewRouter(),
		staticFolder: staticFolder,
	}
}

func (s *Server) String() string {
	return fmt.Sprintf("Port listening selected is %s", s.port)
}

func (s *Server) Handle(method, path string, handlerFc http.HandlerFunc) {
	_, existsThisPath := s.router.rules[path]
	if !existsThisPath {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handlerFc
}

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

func (s *Server) setupStaticFileServer() {
	if s.staticFolder != "" {
		fs := http.FileServer(http.Dir(fmt.Sprintf("./%s", s.staticFolder)))
		folder := fmt.Sprintf("/%s/", s.staticFolder)
		http.Handle(folder, http.StripPrefix(folder, fs))
	}
}
