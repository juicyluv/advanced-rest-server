package server

import (
	"log"
	"net/http"

	"github.com/juicyluv/advanced-rest-server/internal/routes"
)

type Server struct {
	server *http.Server
	config *config
}

func New(cfg *config) *Server {
	return &Server{
		config: cfg,
		server: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: routes.New().Handler(),
		},
	}
}

func (s *Server) Run() error {
	log.Println("Server is running on port " + s.config.Port)
	return s.server.ListenAndServe()
}
