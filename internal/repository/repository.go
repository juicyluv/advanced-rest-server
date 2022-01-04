package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/juicyluv/advanced-rest-server/internal/model/interfaces"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Repository struct {
	Track interfaces.TrackRepository
	Genre interfaces.GenreRepository

	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db:    db,
		Track: NewTrackRepository(db),
		Genre: NewGenreRepository(db),
	}
}
