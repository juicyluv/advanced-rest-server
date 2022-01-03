package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/juicyluv/advanced-rest-server/internal/handler"
	"github.com/juicyluv/advanced-rest-server/internal/logger"
	"github.com/juicyluv/advanced-rest-server/internal/repository"
)

type server struct {
	server *http.Server
	config *Config
	logger logger.Logging
	db     *sqlx.DB
}

func New(cfg *Config, logger logger.Logging, db *sqlx.DB) *server {
	return &server{
		config: cfg,
		logger: logger,
		db:     db,
	}
}

func (s *server) Run() error {
	r := repository.NewRepository(s.db)
	router := handler.New(r)

	s.server = &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: router.Router(),
	}

	s.logger.Printf("Server is up and running on port %s\n", s.config.Port)

	return s.server.ListenAndServe()
}
