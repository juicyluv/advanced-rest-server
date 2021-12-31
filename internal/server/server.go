package server

import (
	"log"
	"net/http"

	"github.com/juicyluv/advanced-rest-server/internal/handler"
)

type server struct {
	server *http.Server
	config *config
}

func New(cfg *config) *server {
	return &server{
		config: cfg,
	}
}

func (s *server) Run() error {
	router := handler.New()

	s.server = &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: router.Router(),
	}

	log.Println("Server is up and running on port " + s.config.Port)

	return s.server.ListenAndServe()
}
