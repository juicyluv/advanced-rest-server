package service

import (
	"github.com/juicyluv/advanced-rest-server/internal/model/interfaces"
	"github.com/juicyluv/advanced-rest-server/internal/repository"
)

type Service struct {
	Track interfaces.TrackService

	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		repository: repository,
		Track:      NewTrackService(repository.Track),
	}
}
