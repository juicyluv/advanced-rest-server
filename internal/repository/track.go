package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/juicyluv/advanced-rest-server/internal/model"
)

var (
	queryTimeout = 3 * time.Second
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

	query, args := insertTrackGenresQuery(t.Genres, t.Id)
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

	ctxGenres, cancelGenres := context.WithTimeout(context.Background(), queryTimeout)
	defer cancelGenres()

	err := r.db.SelectContext(ctxGenres, &genres, query, trackId)
	if err != nil {
		return nil, err
	}

	var track model.Track
	query = `SELECT * FROM track WHERE track_id = $1`

	ctxTrack, cancelTrack := context.WithTimeout(context.Background(), queryTimeout)
	defer cancelTrack()

	err = r.db.GetContext(ctxTrack, &track, query, trackId)
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
	SET title = $1, year = $2, duration = $3, track_url = $4, version = version + 1
	WHERE track_id = $5 AND version = $6
	RETURNING version`
	args := []interface{}{
		track.Title,
		track.Year,
		track.Duration,
		track.TrackURL,
		track.Id,
		track.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	err = tx.QueryRowContext(ctx, query, args...).Scan(&track.Version)
	if err != nil {
		tx.Rollback()
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	// Update track genres
	query = `DELETE FROM track_genre WHERE track_id = $1`

	ctxDelete, cancelDelete := context.WithTimeout(context.Background(), queryTimeout)
	defer cancelDelete()

	_, err = tx.ExecContext(ctxDelete, query, track.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	query, args = insertTrackGenresQuery(track.Genres, track.Id)

	ctxInsert, cancelInsert := context.WithTimeout(context.Background(), queryTimeout)
	defer cancelInsert()

	_, err = tx.ExecContext(ctxInsert, query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *TrackRepository) Delete(trackId int64) error {
	query := `DELETE FROM track WHERE track_id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	_, err := r.db.ExecContext(ctx, query, trackId)
	return err
}

func insertTrackGenresQuery(genres []model.Genre, trackId int64) (string, []interface{}) {
	query := "INSERT INTO track_genre(track_id, genre_id) VALUES"
	genresLen := len(genres)
	args := make([]interface{}, genresLen)
	valuesStr := " (%d, $%d),"
	for i, v := range genres {
		if i == genresLen-1 {
			valuesStr = " (%d, $%d)"
		}
		query += fmt.Sprintf(valuesStr, trackId, i+1)
		args[i] = v.Id
	}

	return query, args
}
