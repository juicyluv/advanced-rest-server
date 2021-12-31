package handler

import "net/http"

// health allows to check information about the server:
// status, version and build states.
func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"status":  "available",
		"version": "1.0.0",
		"build":   "development",
	}

	if err := sendJSON(w, &data, http.StatusOK, nil); err != nil {
		internalErrorResponse(w, r, err)
	}
}
