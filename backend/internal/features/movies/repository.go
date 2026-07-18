package movies

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// compile-time check to ensure repository implements Repository.
var _ Repository = (*repository)(nil)

// todo: use TMDB API or seed your database with 500-1000 movies from TMDB API in cmd/seed/main.go (ignore DI)
type Repository interface {
	// movies
	CreateMovie(ctx context.Context, movie *Movie) error
	GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error)
	GetAllMovies(ctx context.Context) ([]*Movie, error)
	UpdateMovie(ctx context.Context, movie *Movie) error
	DeleteMovie(ctx context.Context, id uuid.UUID) error

	// directors
	CreateDirector(ctx context.Context, director *Director) error
	GetDirectorByID(ctx context.Context, id uuid.UUID) (*Director, error)
	GetAllDirectors(ctx context.Context) ([]*Director, error)
	UpdateDirector(ctx context.Context, director *Director) error
	DeleteDirector(ctx context.Context, id uuid.UUID) error

	// genres
	CreateGenre(ctx context.Context, genre *Genre) error
	GetGenreByID(ctx context.Context, id uuid.UUID) (*Genre, error)
	GetAllGenres(ctx context.Context) ([]*Genre, error)
	UpdateGenre(ctx context.Context, genre *Genre) error
	DeleteGenre(ctx context.Context, id uuid.UUID) error
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

// movies
func (r *repository) CreateMovie(ctx context.Context, movie *Movie) error {
	query := `
		INSERT INTO movies (id, director_id, genre_id, title, release_year, rating, plot, duration, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())`

	_, err := r.db.Exec(
		ctx,
		query,
		movie.ID,
		movie.DirectorID,
		movie.GenreID,
		movie.Title,
		movie.ReleaseYear,
		movie.Rating,
		movie.Plot,
		movie.Duration,
	)

	return err
}

func (r *repository) GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error) {
	query := `
		SELECT id, director_id, genre_id, title, release_year, rating, plot, duration, created_at, updated_at 
		FROM movies 
		WHERE id = $1`

	var m Movie
	err := r.db.QueryRow(ctx, query, id).Scan(
		&m.ID, &m.DirectorID, &m.GenreID, &m.Title,
		&m.ReleaseYear, &m.Rating, &m.Plot, &m.Duration,
		&m.CreatedAt, &m.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &m, nil
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

// directors
func (r *repository) CreateDirector(ctx context.Context, director *Director) error {
	query := `
		INSERT INTO directors (id, name, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())`

	_, err := r.db.Exec(
		ctx,
		query,
		director.ID,
		director.Name,
	)

	return err
}

func (r *repository) GetDirectorByID(ctx context.Context, id uuid.UUID) (*Director, error) {
	panic("not implemented")
}

func (r *repository) GetAllDirectors(ctx context.Context) ([]*Director, error) {
	panic("not implemented")
}

func (r *repository) UpdateDirector(ctx context.Context, director *Director) error {
	panic("not implemented")
}

func (r *repository) DeleteDirector(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
}

// genres
func (r *repository) CreateGenre(ctx context.Context, genre *Genre) error {
	panic("not implemented")
}

func (r *repository) GetGenreByID(ctx context.Context, id uuid.UUID) (*Genre, error) {
	panic("not implemented")
}

func (r *repository) GetAllGenres(ctx context.Context) ([]*Genre, error) {
	panic("not implemented")
}

func (r *repository) UpdateGenre(ctx context.Context, genre *Genre) error {
	panic("not implemented")
}

func (r *repository) DeleteGenre(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
}
