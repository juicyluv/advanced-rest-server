package service

import (
	"github.com/juicyluv/advanced-rest-server/internal/model/interfaces"
	"github.com/juicyluv/advanced-rest-server/internal/repository"
)

type Service struct {
	Track interfaces.TrackService
	Genre interfaces.GenreService

	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	genreService := NewGenreService(repository.Genre)
	return &Service{
		repository: repository,
		Genre:      genreService,
		Track:      NewTrackService(repository.Track, genreService),
	}
}
