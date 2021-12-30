package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	obj := map[string]interface{}{
		"status":  "available",
		"version": "1.0.0",
		"build":   "development",
	}

	if err := sendJSON(w, &obj, http.StatusOK, nil); err != nil {
		message := err.Error()
		log.Println("An error occurred: " + message)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
}

// sendJSON sends given object in JSON format
func sendJSON(w http.ResponseWriter, data interface{}, statusCode int, headers http.Header) error {
	obj, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(obj)

	return nil
}
