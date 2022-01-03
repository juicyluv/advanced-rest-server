package handler

import (
	"github.com/juicyluv/advanced-rest-server/internal/repository"
	"github.com/juicyluv/advanced-rest-server/internal/service"
	"github.com/julienschmidt/httprouter"
)

// Handler is used as http.Server handler for handling HTTP requests.
type Handler struct {
	router  *httprouter.Router
	service *service.Service
}

// jsonResponse is used to send response with the data in JSON format.
type jsonResponse map[string]interface{}

// New returns a new Handler instance.
func New(repository *repository.Repository) *Handler {
	handler := &Handler{
		router:  httprouter.New(),
		service: service.NewService(repository),
	}

	initRoutes(handler)

	return handler
}

// Router returns handler's router.
func (h *Handler) Router() *httprouter.Router {
	return h.router
}
