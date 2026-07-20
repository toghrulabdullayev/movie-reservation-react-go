package showtimes

import (
	"context"

	"github.com/google/uuid"
)

var _ Service = (*service)(nil)

type Service interface {
	// movies
	CreateShowtime(ctx context.Context, req CreateShowtimeRequest) (*Movie, error)
	GetShowtimeByID(ctx context.Context, id uuid.UUID) (*Movie, error)
	GetAllMovies(ctx context.Context) ([]*Movie, error)
	UpdateShowtime(ctx context.Context, movie *Movie) error
	DeleteShowtime(ctx context.Context, id uuid.UUID) error

	// directors
	CreateDirector(ctx context.Context, director CreateDirectorRequest) (*Director, error)
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

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// movies
func (s *service) CreateShowtime(ctx context.Context, req CreateShowtimeRequest) (*Movie, error) {
	movie := &Movie{
		BaseEntity: BaseEntity{
			ID: uuid.New(),
		},
		DirectorID:  req.DirectorID,
		GenreID:     req.GenreID,
		Title:       req.Title,
		ReleaseYear: req.ReleaseYear,
		Rating:      req.Rating,
		Plot:        req.Plot,
		Duration:    req.Duration,
	}

	if err := s.repo.CreateShowtime(ctx, movie); err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *service) GetShowtimeByID(ctx context.Context, id uuid.UUID) (*Movie, error) {
	return s.repo.GetShowtimeByID(ctx, id)
}

func (s *service) GetAllMovies(ctx context.Context) ([]*Movie, error) {
	return s.repo.GetAllMovies(ctx)
}

func (s *service) UpdateShowtime(ctx context.Context, movie *Movie) error {
	return s.repo.UpdateShowtime(ctx, movie)
}

func (s *service) DeleteShowtime(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteShowtime(ctx, id)
}

// directors
func (s *service) CreateDirector(ctx context.Context, req CreateDirectorRequest) (*Director, error) {
	director := &Director{
		BaseEntity: BaseEntity{
			ID: uuid.New(),
		},
		Name: req.Name,
	}

	if err := s.repo.CreateDirector(ctx, director); err != nil {
		return nil, err
	}

	return director, nil
}

func (s *service) GetDirectorByID(ctx context.Context, id uuid.UUID) (*Director, error) {
	return s.repo.GetDirectorByID(ctx, id)
}

func (s *service) GetAllDirectors(ctx context.Context) ([]*Director, error) {
	return s.repo.GetAllDirectors(ctx)
}

func (s *service) UpdateDirector(ctx context.Context, director *Director) error {
	return s.repo.UpdateDirector(ctx, director)
}

func (s *service) DeleteDirector(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteDirector(ctx, id)
}

// genres
func (s *service) CreateGenre(ctx context.Context, genre *Genre) error {
	return s.repo.CreateGenre(ctx, genre)
}

func (s *service) GetGenreByID(ctx context.Context, id uuid.UUID) (*Genre, error) {
	return s.repo.GetGenreByID(ctx, id)
}

func (s *service) GetAllGenres(ctx context.Context) ([]*Genre, error) {
	return s.repo.GetAllGenres(ctx)
}

func (s *service) UpdateGenre(ctx context.Context, genre *Genre) error {
	return s.repo.UpdateGenre(ctx, genre)
}

func (s *service) DeleteGenre(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteGenre(ctx, id)
}
