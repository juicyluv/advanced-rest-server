package handler

import (
	"net/http"

	"github.com/juicyluv/advanced-rest-server/internal/model"
	"github.com/juicyluv/advanced-rest-server/internal/validator"
)

func (h *Handler) getTrack(w http.ResponseWriter, r *http.Request) {
	id, err := readIdParam(r)
	if err != nil {
		return
	}

	track := &model.Track{
		Id:       id,
		Title:    "Casual Title",
		Year:     2002,
		Duration: 201,
		Genres:   []string{"This", "And", "This"},
		Artists:  []string{"Yeah", "Boi"},
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

	if err := sendJSON(w, input, http.StatusOK, nil); err != nil {
		internalErrorResponse(w, r, err)
	}
}
