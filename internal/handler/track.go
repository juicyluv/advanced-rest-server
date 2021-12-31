package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/juicyluv/advanced-rest-server/internal/model"
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
	}

	if err := readJSON(w, r, &input); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	log.Println(input)
}
