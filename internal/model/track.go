package model

import (
	"time"

	"github.com/juicyluv/advanced-rest-server/internal/validator"
)

type Track struct {
	Id       int64    `json:"id"`
	Title    string   `json:"title"`
	Year     int32    `json:"year,omitempty"`
	Duration int32    `json:"duration"`
	Genres   []string `json:"genres,omitempty"`
	Artists  []string `json:"artists"`
	Album    string   `json:"album,omitempty"`
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

	// Genres, Artists
	v.UniqueStrings(t.Genres)
	v.UniqueStrings(t.Artists)

	// Album
	v.NotEmpty(t.Album, "album")
}
