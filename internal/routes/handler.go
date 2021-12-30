package routes

import (
	"net/http"

	"github.com/juicyluv/advanced-rest-server/internal/logger"
)

type handler struct {
	logger *logger.Logger
}

func newHandler() *handler {
	return &handler{
		logger: logger.New(logger.LevelDebug),
	}
}

func (h *handler) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}
