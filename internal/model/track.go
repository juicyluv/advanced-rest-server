package model

import (
	"time"

	"github.com/juicyluv/advanced-rest-server/internal/validator"
)

type Track struct {
	Id       int64   `json:"id" db:"track_id"`
	Title    string  `json:"title"`
	Year     int32   `json:"year"`
	Duration int32   `json:"duration"`
	Genres   []Genre `json:"genres"`
	Artists  []int   `json:"artists"`
	TrackURL string  `json:"trackUrl" db:"track_url"`
}

type UpdateTrack struct {
	Title    *string `json:"title"`
	Year     *int32  `json:"year"`
	Genres   []Genre `json:"genres"`
	TrackURL *string `json:"trackUrl"`
}

// Validate checks whether track instance is valid. If it's not, it
// will fill up given validator Error map with appropriate error messages.
func (t *Track) Validate(v *validator.Validator) {
	// Title
	v.NotEmpty(t.Title, "title")
	v.MaxLength(t.Title, 100, "title")

	// Year
	v.Min(int(t.Year), 1888, "year")
	v.Max(int(t.Year), time.Now().Year(), "year")

	// Duration
	v.Min(int(t.Duration), 1, "duration")
}

func (t *UpdateTrack) Copy(track *Track) {
	if t.Title != nil {
		track.Title = *t.Title
	}

	if t.Year != nil {
		track.Year = *t.Year
	}

	if t.Genres != nil {
		track.Genres = t.Genres
	}

	if t.TrackURL != nil {
		track.TrackURL = *t.TrackURL
	}
}
