package handler

import (
	"fmt"
	"net/http"

	"github.com/juicyluv/advanced-rest-server/internal/validator"
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

// badRequestResponse sends a 404 Bad Request status code and JSON error message.
func badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

// failedValidationResponse sends a 422 UnpocessableEntity response with field errors if validation fails.
func failedValidationResponse(w http.ResponseWriter, r *http.Request, errors validator.ValidatorErrors) {
	errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
