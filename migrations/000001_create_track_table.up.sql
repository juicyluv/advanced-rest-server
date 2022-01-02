CREATE TABLE IF NOT EXISTS track(
    track_id bigserial PRIMARY KEY,
    title text,
    year int,
    duration int
);