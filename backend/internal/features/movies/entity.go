package movies

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID           uuid.UUID `json:"id"`
	TMDBID       int       `json:"tmdb_id"`
	Title        string    `json:"title"`
	Overview     string    `json:"overview"`
	PosterPath   string    `json:"poster_path"`
	BackdropPath string    `json:"backdrop_path"`
	ReleaseDate  time.Time `json:"release_date"`
	Runtime      int       `json:"runtime"`
	VoteAverage  float64   `json:"vote_average"`
	Genres       []string  `json:"genres,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
