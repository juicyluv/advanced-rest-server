package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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

// Returns an "id" url param as int and error if any.
// Example: "/api/v1/movies/32" returns (32, nil).
func readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func logError(r *http.Request, err error) {
	log.Println(err)
}
