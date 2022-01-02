ALTER TABLE track ADD CONSTRAINT track_duration_check CHECK (duration > 0);

ALTER TABLE track ADD CONSTRAINT track_year_check CHECK (year BETWEEN 1888 AND date_part('year', now()));