package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// sendJSON sends given object in JSON format with given headers.
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

// readIdParam returns an id of the URL path on success.
func readIdParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

// logError logs an error to logger output.
func logError(r *http.Request, err error) {
	log.Println(err)
}
