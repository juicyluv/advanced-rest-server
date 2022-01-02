package interfaces

import "github.com/juicyluv/advanced-rest-server/internal/model"

type TrackRepository interface {
	Insert(track *model.Track) error
	FindAll() ([]model.Track, error)
	FindById(id int64) (*model.Track, error)
	Update(track *model.Track) error
	Delete(id int64) error
}
