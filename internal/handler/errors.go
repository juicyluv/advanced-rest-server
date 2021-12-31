package handler

import (
	"fmt"
	"net/http"
)

// errorResponse logs an error and sends a JSON response with a given status code.
func errorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message interface{}) {
	msg := jsonResponse{"error": message}

	if err := sendJSON(w, msg, statusCode, nil); err != nil {
		logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// internalErrorResponse logs the error message and sends a 500 Internal Server Error by using
// errorResponse helper function.
func internalErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)

	message := "the server encountered a problem and could not process your request"
	errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse sends 404 Not Found status code and
// JSON error response to the client.
func notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponse(w, r, http.StatusNotFound, message)
}

// methodNotAllowedResponse sends a 405 Method Not Allowed
// status code and JSON response to the client.
func methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
