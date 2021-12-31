package handler

import (
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
