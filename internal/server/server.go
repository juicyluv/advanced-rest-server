package server

import (
	"log"
	"net/http"

	"github.com/juicyluv/advanced-rest-server/internal/handler"
)

type Server struct {
	server *http.Server
	config *config
}

func New(cfg *config) *Server {
	return &Server{
		config: cfg,
	}
}

func (s *Server) Run() error {
	router := handler.NewHttpRouter()

	s.server = &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: router,
	}

	log.Println("Server is up and running on port " + s.config.Port)

	return s.server.ListenAndServe()
}
