package service

import (
	"github.com/juicyluv/advanced-rest-server/internal/model"
	"github.com/juicyluv/advanced-rest-server/internal/model/interfaces"
)

type TrackService struct {
	repository   interfaces.TrackRepository
	genreService interfaces.GenreService
}

func NewTrackService(repository interfaces.TrackRepository, gs interfaces.GenreService) *TrackService {
	return &TrackService{
		repository:   repository,
		genreService: gs,
	}
}

func (s *TrackService) Create(track *model.Track) error {
	err := s.repository.Insert(track)
	if err != nil {
		return err
	}

	genres, err := s.genreService.GetTrackGenres(track.Id)
	if err != nil {
		return err
	}

	track.Genres = genres

	return nil
}
func (s *TrackService) GetAll(filters *model.TrackFilter) ([]model.Track, error) {
	return s.repository.FindAll(filters)
}

func (s *TrackService) GetById(id int64) (*model.Track, error) {
	return s.repository.FindById(id)
}

func (s *TrackService) Update(track *model.Track) error {
	err := s.repository.Update(track)
	if err != nil {
		return err
	}

	genres, err := s.genreService.GetTrackGenres(track.Id)
	if err != nil {
		return err
	}

	track.Genres = genres

	return nil
}

func (s *TrackService) Delete(trackId int64) error {
	_, err := s.GetById(trackId)
	if err != nil {
		return err
	}

	return s.repository.Delete(trackId)
}
