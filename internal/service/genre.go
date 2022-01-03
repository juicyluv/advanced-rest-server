package service

import (
	"github.com/juicyluv/advanced-rest-server/internal/model"
	"github.com/juicyluv/advanced-rest-server/internal/model/interfaces"
)

type GenreService struct {
	repository interfaces.GenreRepository
}

func NewGenreService(repository interfaces.GenreRepository) *GenreService {
	return &GenreService{
		repository: repository,
	}
}

func (s *GenreService) Create(g *model.Genre) error {
	return s.repository.Insert(g)
}
func (s *GenreService) GetAll() ([]model.Genre, error) {
	return nil, nil
}

func (s *GenreService) GetById(id int64) (*model.Genre, error) {
	return nil, nil
}

func (s *GenreService) Update(track *model.Genre) error {
	return nil
}

func (s *GenreService) Delete(id int64) error {
	return nil
}
