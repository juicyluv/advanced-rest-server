CREATE TABLE IF NOT EXISTS track(
    track_id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    year INT NOT NULL,
    duration INT NOT NULL
);