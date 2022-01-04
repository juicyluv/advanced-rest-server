package handler

import (
	"fmt"
	"net/http"

	"github.com/juicyluv/advanced-rest-server/internal/model"
	"github.com/juicyluv/advanced-rest-server/internal/validator"
)

func (h *Handler) getTrack(w http.ResponseWriter, r *http.Request) {
	id, err := readIdParam(r)
	if err != nil {
		return
	}

	track, err := h.service.Track.GetById(id)
	if err != nil {
		internalErrorResponse(w, r, err)
		return
	}

	err = sendJSON(w, track, http.StatusOK, nil)
	if err != nil {
		internalErrorResponse(w, r, err)
	}
}

func (h *Handler) createTrack(w http.ResponseWriter, r *http.Request) {
	input := &model.Track{}

	if err := readJSON(w, r, &input); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	input.Validate(v)

	if !v.Valid() {
		failedValidationResponse(w, r, v.Errors)
		return
	}

	if err := h.service.Track.Create(input); err != nil {
		internalErrorResponse(w, r, err)
		return
	}

	genres, err := h.service.Genre.GetTrackGenres(input.Id)
	if err != nil {
		internalErrorResponse(w, r, err)
		return
	}

	input.Genres = genres

	fmt.Printf("%v", input)

	if err := sendJSON(w, input, http.StatusOK, nil); err != nil {
		internalErrorResponse(w, r, err)
	}
}
