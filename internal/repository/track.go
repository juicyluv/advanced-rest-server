package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/juicyluv/advanced-rest-server/internal/model"
)

type TrackRepository struct {
	db *sqlx.DB
}

func NewTrackRepository(db *sqlx.DB) *TrackRepository {
	return &TrackRepository{
		db: db,
	}
}

func (r *TrackRepository) Insert(t *model.Track) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
	INSERT INTO track (title, year, duration, track_url) 
	VALUES ($1, $2, $3, $4) 
	RETURNING track_id`

	err = tx.QueryRow(query, t.Title, t.Year, t.Duration, t.TrackURL).Scan(&t.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = "INSERT INTO track_genre(track_id, genre_id) VALUES"
	genresLen := len(t.Genres)
	args := make([]interface{}, genresLen)
	valuesStr := " (%d, $%d),"
	for i, v := range t.Genres {
		if i == genresLen-1 {
			valuesStr = " (%d, $%d)"
		}
		query += fmt.Sprintf(valuesStr, t.Id, i+1)
		args[i] = v.Id
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *TrackRepository) FindAll() ([]model.Track, error) {
	return nil, nil
}

func (r *TrackRepository) FindById(trackId int64) (*model.Track, error) {
	var genres []model.Genre

	query := `
	SELECT g.* 
	FROM track_genre tg 
	INNER JOIN genre g 
	ON g.genre_id = tg.genre_id 
	WHERE track_id = $1`

	err := r.db.Select(&genres, query, trackId)
	if err != nil {
		return nil, err
	}

	var track model.Track
	query = `SELECT * FROM track WHERE track_id = $1`

	err = r.db.Get(&track, query, trackId)
	if err != nil {
		return nil, err
	}

	track.Genres = genres

	return &track, nil
}

func (r *TrackRepository) Update(track *model.Track) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
	UPDATE track 
	SET title = $1, year = $2, duration = $3, track_url = $4
	WHERE track_id = $5`
	args := []interface{}{track.Title, track.Year, track.Duration, track.TrackURL, track.Id}
	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Update track genres
	query = `DELETE FROM track_genre WHERE track_id = $1`
	_, err = tx.Exec(query, track.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = "INSERT INTO track_genre(track_id, genre_id) VALUES"
	genresLen := len(track.Genres)
	args = make([]interface{}, genresLen)
	valuesStr := " (%d, $%d),"
	for i, v := range track.Genres {
		if i == genresLen-1 {
			valuesStr = " (%d, $%d)"
		}
		query += fmt.Sprintf(valuesStr, track.Id, i+1)
		args[i] = v.Id
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *TrackRepository) Delete(trackId int64) error {
	query := `DELETE FROM track WHERE track_id = $1`
	_, err := r.db.Exec(query, trackId)
	return err
}
