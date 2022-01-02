package interfaces

import "github.com/juicyluv/advanced-rest-server/internal/model"

type TrackService interface {
	Create(track *model.Track) error
	GetAll() ([]model.Track, error)
	GetById(id int64) (*model.Track, error)
	Update(track *model.Track) error
	Delete(id int64) error
}
