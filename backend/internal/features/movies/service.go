package movies

import (
	"context"

	"github.com/google/uuid"
)

var _ Service = (*service)(nil)

type Service interface {
	CreateMovie(ctx context.Context, movie *Movie) error
	GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error)
	GetAllMovies(ctx context.Context) ([]*Movie, error)
	UpdateMovie(ctx context.Context, movie *Movie) error
	DeleteMovie(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateMovie(ctx context.Context, movie *Movie) error {
	return s.repo.CreateMovie(ctx, movie)
}

func (s *service) GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error) {
	return s.repo.GetMovieByID(ctx, id)
}

func (s *service) GetAllMovies(ctx context.Context) ([]*Movie, error) {
	return s.repo.GetAllMovies(ctx)
}

func (s *service) UpdateMovie(ctx context.Context, movie *Movie) error {
	return s.repo.UpdateMovie(ctx, movie)
}

func (s *service) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteMovie(ctx, id)
}
