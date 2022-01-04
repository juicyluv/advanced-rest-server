package handler

import "net/http"

// initRoutes initializes handlers for API routes.
func initRoutes(h *Handler) {
	h.router.NotFound = http.HandlerFunc(notFoundResponse)
	h.router.MethodNotAllowed = http.HandlerFunc(methodNotAllowedResponse)

	h.router.HandlerFunc(http.MethodGet, "/api/v1/health", h.health)
	h.router.HandlerFunc(http.MethodGet, "/api/v1/tracks/:id", h.getTrack)
	h.router.HandlerFunc(http.MethodPost, "/api/v1/tracks", h.createTrack)
	h.router.HandlerFunc(http.MethodPut, "/api/v1/tracks/:id", h.updateTrack)
	h.router.HandlerFunc(http.MethodDelete, "/api/v1/tracks/:id", h.deleteTrack)
}
