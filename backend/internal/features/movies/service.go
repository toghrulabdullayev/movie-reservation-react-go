package movies

import (
	"context"

	"github.com/google/uuid"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/shared/pagination"
)

var _ Service = (*service)(nil)

type Service interface {
	GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error)
	GetAllMovies(ctx context.Context, params pagination.PaginationParams) ([]*Movie, int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error) {
	return s.repo.GetMovieByID(ctx, id)
}

func (s *service) GetAllMovies(ctx context.Context, params pagination.PaginationParams) ([]*Movie, int, error) {
	return s.repo.GetAllMovies(ctx, params)
}
