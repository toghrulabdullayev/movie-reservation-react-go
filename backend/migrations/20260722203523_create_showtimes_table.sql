-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS showtimes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    movie_id UUID NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    room_id UUID NOT NULL REFERENCES rooms(id) ON DELETE RESTRICT,

    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE NOT NULL,

    base_price NUMERIC(10,2) NOT NULL,

    status VARCHAR(20) NOT NULL DEFAULT 'scheduled'
        CHECK (status IN (
            'scheduled',
            'ongoing',
            'completed',
            'canceled'
        )),

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    CHECK (end_time > start_time)
);

-- CREATE INDEX IF NOT EXISTS idx_showtimes_movie_id
--     ON showtimes(movie_id);

-- CREATE INDEX IF NOT EXISTS idx_showtimes_room_id
--     ON showtimes(room_id);

-- CREATE INDEX IF NOT EXISTS idx_showtimes_start_time
--     ON showtimes(start_time);

-- CREATE INDEX IF NOT EXISTS idx_showtimes_status
--     ON showtimes(status);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS showtimes;