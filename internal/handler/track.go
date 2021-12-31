package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/juicyluv/advanced-rest-server/internal/model"
	"github.com/juicyluv/advanced-rest-server/internal/validator"
)

func (h *Handler) getTrack(w http.ResponseWriter, r *http.Request) {
	id, err := readIdParam(r)
	if err != nil {
		return
	}

	track := &model.Track{
		Id:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     "Casual Title",
		Year:      2002,
		Duration:  201,
		Genres:    []string{"This", "And", "This"},
		Artists:   []string{"Yeah", "Boi"},
	}

	err = sendJSON(w, track, http.StatusOK, nil)
	if err != nil {
		internalErrorResponse(w, r, err)
	}
}

func (h *Handler) createTrack(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Year  int32  `json:"year"`
	}

	if err := readJSON(w, r, &input); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.Title != "", "title", "must not be empty")
	v.Check(len(input.Title) >= 1, "title", "the length must be at least 1 character")
	v.Check(len(input.Title) <= 50, "title", "the length must be not greater than 50 characters")

	v.Check(input.Year != 0, "year", "must be provided")
	v.Check(input.Year > 1888, "year", "must be greater than 1888")
	v.Check(input.Year <= int32(time.Now().Year()), "year", "must be not greater than current year")

	if !v.Valid() {
		failedValidationResponse(w, r, v.Errors)
		return
	}

	log.Println(input)
}
