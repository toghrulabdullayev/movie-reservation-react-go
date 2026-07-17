package movies

import (
	"github.com/google/uuid"
)

type MovieResponse struct {
	ID          uuid.UUID `json:"id"`
	DirectorID  uuid.UUID `json:"director_id"`
	GenreID     uuid.UUID `json:"genre_id"`
	Title       string    `json:"title"`
	ReleaseYear int       `json:"release_year"`
	Rating      float32   `json:"rating"`
	Plot        string    `json:"plot"`
	Duration    int       `json:"duration"`
}

// add to dto.go
type CreateMovieRequest struct {
	DirectorID  uuid.UUID `json:"director_id" validate:"required"`
	GenreID     uuid.UUID `json:"genre_id" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	ReleaseYear int       `json:"release_year"`
	Rating      float32   `json:"rating"`
	Plot        string    `json:"plot"`
	Duration    int       `json:"duration" validate:"required,min=1"`
}

type UpdateMovieRequest struct {
	Title       string  `json:"title" validate:"required"`
	ReleaseYear int     `json:"release_year"`
	Rating      float32 `json:"rating"`
	Plot        string  `json:"plot"`
	Duration    int     `json:"duration" validate:"required,min=1"`
}

type CreateDirectorRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateGenreRequest struct {
	Name string `json:"name" validate:"required"`
}

func toMovieResponse(m Movie) MovieResponse {
	return MovieResponse{
		ID:          m.ID,
		DirectorID:  m.DirectorID,
		GenreID:     m.GenreID,
		Title:       m.Title,
		ReleaseYear: m.ReleaseYear,
		Rating:      m.Rating,
		Plot:        m.Plot,
		Duration:    m.Duration,
	}
}
