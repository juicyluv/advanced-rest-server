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

func (r *GenreRepository) GetTrackGenres(trackId int64) ([]model.Genre, error) {
	genres := []model.Genre{}

	query := `
	SELECT g.genre_id, g.genre 
	FROM track_genre tr 
	INNER JOIN genre g 
	ON  g.genre_id = tr.genre_id
	WHERE track_id=$1`

	err := r.db.Select(&genres, query, trackId)
	if err != nil {
		return nil, err
	}

	return genres, nil
}
