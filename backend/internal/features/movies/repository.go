package movies

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// compile-time check to ensure repository implements Repository.
var _ Repository = (*repository)(nil)

type Repository interface {
	CreateMovie(ctx context.Context, movie *Movie) error
	GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error)
	GetAllMovies(ctx context.Context) ([]*Movie, error)
	UpdateMovie(ctx context.Context, movie *Movie) error
	DeleteMovie(ctx context.Context, id uuid.UUID) error
}

// concrete implementation — lowercase, unexported
type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateMovie(ctx context.Context, movie *Movie) error {
	// TODO: actual pgx query, e.g.
	// _, err := r.db.Exec(ctx, "INSERT INTO movies (...) VALUES (...)", ...)
	panic("not implemented")
}

func (r *repository) GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error) {
	panic("not implemented")
}

func (r *repository) GetAllMovies(ctx context.Context) ([]*Movie, error) {
	panic("not implemented")
}

func (r *repository) UpdateMovie(ctx context.Context, movie *Movie) error {
	panic("not implemented")
}

func (r *repository) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
}

// type DirectorRepository interface {
// 	CreateDirector(ctx context.Context, director *Director) error
// 	GetDirectorByID(ctx context.Context, id uuid.UUID) (*Director, error)
// 	GetAllDirectors(ctx context.Context) ([]*Director, error)
// 	UpdateDirector(ctx context.Context, director *Director) error
// 	DeleteDirector(ctx context.Context, id uuid.UUID) error
// }

// type GenreRepository interface {
// 	CreateGenre(ctx context.Context, genre *Genre) error
// 	GetGenreByID(ctx context.Context, id uuid.UUID) (*Genre, error)
// 	GetAllGenres(ctx context.Context) ([]*Genre, error)
// 	UpdateGenre(ctx context.Context, genre *Genre) error
// 	DeleteGenre(ctx context.Context, id uuid.UUID) error
// }
