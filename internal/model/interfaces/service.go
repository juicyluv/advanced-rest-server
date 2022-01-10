package interfaces

import "github.com/juicyluv/advanced-rest-server/internal/model"

type TrackService interface {
	Create(t *model.Track) error
	GetAll(filters *model.TrackFilter) ([]model.Track, error)
	GetById(id int64) (*model.Track, error)
	Update(t *model.Track) error
	Delete(id int64) error
}

type GenreService interface {
	Create(g *model.Genre) error
	GetAll() ([]model.Genre, error)
	GetById(id int64) (*model.Genre, error)
	Update(g *model.Genre) error
	Delete(id int64) error
	GetTrackGenres(trackId int64) ([]model.Genre, error)
}
