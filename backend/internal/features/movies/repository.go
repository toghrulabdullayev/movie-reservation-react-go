package movies

import (
	"context"

	"github.com/google/uuid"
)

type MovieRepository interface {
	CreateMovie(ctx context.Context, movie *Movie) error
	GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error)
	GetAllMovies(ctx context.Context) ([]*Movie, error)
	UpdateMovie(ctx context.Context, movie *Movie) error
	DeleteMovie(ctx context.Context, id uuid.UUID) error
}

type DirectorRepository interface {
	CreateDirector(ctx context.Context, director *Director) error
	GetDirectorByID(ctx context.Context, id uuid.UUID) (*Director, error)
	GetAllDirectors(ctx context.Context) ([]*Director, error)
	UpdateDirector(ctx context.Context, director *Director) error
	DeleteDirector(ctx context.Context, id uuid.UUID) error
}

type GenreRepository interface {
	CreateGenre(ctx context.Context, genre *Genre) error
	GetGenreByID(ctx context.Context, id uuid.UUID) (*Genre, error)
	GetAllGenres(ctx context.Context) ([]*Genre, error)
	UpdateGenre(ctx context.Context, genre *Genre) error
	DeleteGenre(ctx context.Context, id uuid.UUID) error
}
