package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	h      *handler
	router *httprouter.Router
}

func New() *Router {
	router := &Router{
		h:      newHandler(),
		router: httprouter.New(),
	}

	router.initRoutes()

	return router
}

func (r *Router) initRoutes() {
	r.router.HandlerFunc(http.MethodGet, "/api/v1", r.h.index)
	r.router.HandlerFunc(http.MethodGet, "/api/v1/health", r.h.health)
}

func (r *Router) Handler() *httprouter.Router {
	return r.router
}
