package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/juicyluv/advanced-rest-server/internal/model"
)

type GenreRepository struct {
	db *sqlx.DB
}

func NewGenreRepository(db *sqlx.DB) *GenreRepository {
	return &GenreRepository{
		db: db,
	}
}

func (r *GenreRepository) Insert(g *model.Genre) error {
	query := `INSERT INTO genre (genre) VALUES ($1)`
	_, err := r.db.Exec(query, g.Genre)
	return err
}

func (r *GenreRepository) FindAll() ([]model.Genre, error) {
	return nil, nil
}

func (r *GenreRepository) FindById(id int64) (*model.Genre, error) {
	return nil, nil
}

func (r *GenreRepository) Update(genre *model.Genre) error {
	return nil
}

func (r *GenreRepository) Delete(id int64) error {
	return nil
}
