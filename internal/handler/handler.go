package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handler is used as http.Server handler for handling HTTP requests.
type Handler struct {
	router *httprouter.Router
}

// jsonResponse is used to send response with the data in JSON format.
type jsonResponse map[string]interface{}

// New returns a new Handler instance.
func New() *Handler {
	handler := &Handler{
		router: httprouter.New(),
	}

	initRoutes(handler)

	return handler
}

// Router returns handler's router.
func (h *Handler) Router() *httprouter.Router {
	return h.router
}

// initRoutes initializes handlers for API routes.
func initRoutes(h *Handler) {
	h.router.NotFound = http.HandlerFunc(notFoundResponse)
	h.router.MethodNotAllowed = http.HandlerFunc(methodNotAllowedResponse)

	h.router.HandlerFunc(http.MethodGet, "/api/v1/health", h.health)
	h.router.HandlerFunc(http.MethodGet, "/api/v1/tracks/:id", h.getTrack)

}
