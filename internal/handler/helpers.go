package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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

// readJSON decodes request body to the given destination(usually model struct)
// and if an error occurred returns specific error message.
func readJSON(w http.ResponseWriter, r *http.Request, dest interface{}) error {
	// Create a new decoder and check for unknown fields
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dest)
	if err != nil {
		// If an error occurred, send an error mapped to JSON decoding error
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		// Syntax error
		case errors.As(err, &syntaxError):
			return fmt.Errorf(
				"request body contains badly-formatted JSON (at character %d)",
				syntaxError.Offset,
			)
		// Type error
		case errors.As(err, &unmarshalTypeError):
			// If there's an info for struct field, show what field contains an error
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf(
					"request body contains incorrect JSON type for field %q",
					unmarshalTypeError.Field,
				)
			}

			return fmt.Errorf(
				"request body contains incorrect JSON type (at character %d)",
				unmarshalTypeError.Offset,
			)
		// Unmarshall error
		case errors.As(err, &invalidUnmarshalError):
			// We are panicing here because this is unexpected error
			panic(err)
		// Empty JSON error
		case errors.Is(err, io.EOF):
			return errors.New("request body must not be empty")
		// Unknown field error
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("request body contains unknown key %s", fieldName)

		// Return error as-is
		default:
			return err
		}
	}

	// Decode one more time to check wheter here is another JSON object
	if err = dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("request body must only contain single JSON value")
	}

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

func (h *Handler) readQueryString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)

	// If no key exists (or the value is empty) then return the default value.
	if s == "" {
		return defaultValue
	}

	return s
}

func (h *Handler) readQueryCSV(qs url.Values, key string, defaultValue []string) []string {
	csv := qs.Get(key)

	if csv == "" {
		return defaultValue
	}

	return strings.Split(csv, ",")
}

func (h *Handler) readQueryInt(qs url.Values, key string, defaultValue int) int {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}

	return i
}
