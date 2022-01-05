package handler

import (
	"errors"
	"net/http"

	"github.com/juicyluv/advanced-rest-server/internal/model"
	"github.com/juicyluv/advanced-rest-server/internal/repository"
	"github.com/juicyluv/advanced-rest-server/internal/validator"
)

func (h *Handler) getTrack(w http.ResponseWriter, r *http.Request) {
	id, err := readIdParam(r)
	if err != nil {
		return
	}

	track, err := h.service.Track.GetById(id)
	if err != nil {
		switch {
		case errors.Is(err, errNoRows):
			badRequestResponse(w, r, errNoRowsResponse)
		default:
			internalErrorResponse(w, r, err)
		}
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

	if err := sendJSON(w, input, http.StatusOK, nil); err != nil {
		internalErrorResponse(w, r, err)
	}
}

func (h *Handler) updateTrack(w http.ResponseWriter, r *http.Request) {
	id, err := readIdParam(r)
	if err != nil {
		return
	}

	track, err := h.service.Track.GetById(id)
	if err != nil {
		switch {
		case errors.Is(err, errNoRows):
			badRequestResponse(w, r, errNoRowsResponse)
		case errors.Is(err, repository.ErrEditConflict):
			editConflictResponse(w, r)
		default:
			internalErrorResponse(w, r, err)
		}
		return
	}

	var input model.UpdateTrack
	err = readJSON(w, r, &input)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	input.Copy(track)
	v := validator.New()
	track.Validate(v)

	if !v.Valid() {
		failedValidationResponse(w, r, v.Errors)
		return
	}

	err = h.service.Track.Update(track)
	if err != nil {
		internalErrorResponse(w, r, err)
		return
	}

	err = sendJSON(w, track, http.StatusOK, nil)
	if err != nil {
		internalErrorResponse(w, r, err)
	}
}

func (h *Handler) deleteTrack(w http.ResponseWriter, r *http.Request) {
	id, err := readIdParam(r)
	if err != nil {
		return
	}

	if err := h.service.Track.Delete(id); err != nil {
		switch {
		case errors.Is(err, errNoRows):
			badRequestResponse(w, r, errNoRowsResponse)
		default:
			internalErrorResponse(w, r, err)
		}
		return
	}

	if err := sendJSON(w, jsonResponse{"message": "record has been deleted"}, http.StatusOK, nil); err != nil {
		internalErrorResponse(w, r, err)
	}
}
