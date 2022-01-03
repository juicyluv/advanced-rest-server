CREATE TABLE IF NOT EXISTS track_genre (
    track_id BIGINT NOT NULL,
    genre_id SMALLINT NOT NULL,

    FOREIGN KEY(track_id) REFERENCES track(track_id),
    FOREIGN KEY(genre_id) REFERENCES genre(genre_id)
);