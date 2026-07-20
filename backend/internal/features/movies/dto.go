package movies

import (
	"time"

	"github.com/google/uuid"
)

type MovieResponse struct {
	ID           uuid.UUID `json:"id"`
	TMDBID       int       `json:"tmdb_id"`
	Title        string    `json:"title"`
	Overview     string    `json:"overview"`
	PosterPath   string    `json:"poster_path"`
	BackdropPath string    `json:"backdrop_path"`
	ReleaseDate  time.Time `json:"release_date"`
	Runtime      int       `json:"runtime"`
	VoteAverage  float64   `json:"vote_average"`
	Genres       []string  `json:"genres"`
}

func toMovieResponse(movie *Movie) *MovieResponse {
	if movie == nil {
		return nil
	}
	return &MovieResponse{
		ID:           movie.ID,
		TMDBID:       movie.TMDBID,
		Title:        movie.Title,
		Overview:     movie.Overview,
		PosterPath:   movie.PosterPath,
		BackdropPath: movie.BackdropPath,
		ReleaseDate:  movie.ReleaseDate,
		Runtime:      movie.Runtime,
		VoteAverage:  movie.VoteAverage,
		Genres:       movie.Genres,
	}
}
