package server

import (
	"net/http"

	"github.com/juicyluv/advanced-rest-server/internal/handler"
	"github.com/juicyluv/advanced-rest-server/internal/logger"
)

type server struct {
	server *http.Server
	config *config
	logger logger.Logging
}

func New(cfg *config, logger logger.Logging) *server {
	return &server{
		config: cfg,
		logger: logger,
	}
}

func (s *server) Run() error {
	router := handler.New()

	s.server = &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: router.Router(),
	}

	s.logger.Printf("Server is up and running on port %s\n", s.config.Port)

	return s.server.ListenAndServe()
}
