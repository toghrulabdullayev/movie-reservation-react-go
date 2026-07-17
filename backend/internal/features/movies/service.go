package movies

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateMovie(ctx context.Context, movie *Movie) error {
	return s.repo.CreateMovie(ctx, movie)
}

func (s *Service) GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error) {
	return s.repo.GetMovieByID(ctx, id)
}

func (s *Service) GetAllMovies(ctx context.Context) ([]*Movie, error) {
	return s.repo.GetAllMovies(ctx)
}

func (s *Service) UpdateMovie(ctx context.Context, movie *Movie) error {
	return s.repo.UpdateMovie(ctx, movie)
}

func (s *Service) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteMovie(ctx, id)
}
