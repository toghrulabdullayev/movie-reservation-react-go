--? SQL in this section is executed when the migration is applied.

-- +goose Up
-- +goose StatementBegin

-- 1. genres table
CREATE TABLE IF NOT EXISTS genres (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. movies table (optional fields are fully nullable with no fake default placeholders)
CREATE TABLE IF NOT EXISTS movies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tmdb_id INT UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    overview TEXT,
    poster_path VARCHAR(512),
    backdrop_path VARCHAR(512),
    release_date DATE,
    runtime INT CHECK (runtime >= 0),
    vote_average NUMERIC(3, 1) CHECK (vote_average BETWEEN 0.0 AND 10.0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 3. movie genres join table (bridge tables do not require timestamp metrics)
CREATE TABLE IF NOT EXISTS movie_genres (
    movie_id UUID REFERENCES movies(id) ON DELETE CASCADE,
    genre_id UUID REFERENCES genres(id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, genre_id)
);

-- +goose StatementEnd

--? SQL in this section is executed when the migration is rolled back.

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movie_genres CASCADE;
DROP TABLE IF EXISTS movies CASCADE;
DROP TABLE IF EXISTS genres CASCADE;
-- +goose StatementEnd