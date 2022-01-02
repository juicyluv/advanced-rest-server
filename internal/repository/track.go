package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/juicyluv/advanced-rest-server/internal/model"
)

type TrackRepository struct {
	db *sqlx.DB
}

func NewTrackRepository(db *sqlx.DB) *TrackRepository {
	return &TrackRepository{
		db: db,
	}
}

func (r *TrackRepository) Insert(track *model.Track) error {
	return nil
}

func (r *TrackRepository) FindAll() ([]model.Track, error) {
	return nil, nil
}

func (r *TrackRepository) FindById(id int64) (*model.Track, error) {
	return nil, nil
}

func (r *TrackRepository) Update(track *model.Track) error {
	return nil
}

func (r *TrackRepository) Delete(id int64) error {
	return nil
}
