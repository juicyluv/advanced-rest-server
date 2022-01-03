package interfaces

import "github.com/juicyluv/advanced-rest-server/internal/model"

type TrackRepository interface {
	Insert(track *model.Track) error
	FindAll() ([]model.Track, error)
	FindById(id int64) (*model.Track, error)
	Update(track *model.Track) error
	Delete(id int64) error
}

type GenreRepository interface {
	Insert(genre *model.Genre) error
	FindAll() ([]model.Genre, error)
	FindById(id int64) (*model.Genre, error)
	Update(track *model.Genre) error
	Delete(id int64) error
}
