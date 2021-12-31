package model

import "time"

type Track struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"publish_date"`
	UpdatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Duration  int32     `json:"duration"`
	Genres    []string  `json:"genres,omitempty"`
	Artists   []string  `json:"artists,omitempty"`
}
