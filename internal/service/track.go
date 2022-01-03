package service

import (
	"github.com/juicyluv/advanced-rest-server/internal/model"
	"github.com/juicyluv/advanced-rest-server/internal/model/interfaces"
)

type TrackService struct {
	repository interfaces.TrackRepository
}

func NewTrackService(repository interfaces.TrackRepository) *TrackService {
	return &TrackService{
		repository: repository,
	}
}

func (s *TrackService) Create(track *model.Track) error {
	return s.repository.Insert(track)
}
func (s *TrackService) GetAll() ([]model.Track, error) {
	return nil, nil
}

func (s *TrackService) GetById(id int64) (*model.Track, error) {
	return nil, nil
}

func (s *TrackService) Update(track *model.Track) error {
	return nil
}

func (s *TrackService) Delete(id int64) error {
	return nil
}
