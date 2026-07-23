package rooms

import (
	"context"

	"github.com/google/uuid"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/shared/pagination"
)

var _ Service = (*service)(nil)

type Service interface {
	// rooms
	CreateRoom(ctx context.Context, req CreateRoomRequest) (*Room, error)
	GetRoomByID(ctx context.Context, id uuid.UUID) (*Room, error)
	GetAllRooms(ctx context.Context, params pagination.PaginationParams) ([]*Room, int, error)
	UpdateRoom(ctx context.Context, room *Room) error
	DeleteRoom(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// rooms
func (s *service) CreateRoom(ctx context.Context, req CreateRoomRequest) (*Room, error) {
	room := &Room{
		ID:       uuid.New(),
		Name:     req.Name,
		Capacity: req.Capacity,
		Rows:     req.Rows,
		Columns:  req.Columns,
	}

	if err := s.repo.CreateRoom(ctx, room); err != nil {
		return nil, err
	}

	return room, nil
}

func (s *service) GetRoomByID(ctx context.Context, id uuid.UUID) (*Room, error) {
	return s.repo.GetRoomByID(ctx, id)
}

func (s *service) GetAllRooms(ctx context.Context, params pagination.PaginationParams) ([]*Room, int, error) {
	return s.repo.GetAllRooms(ctx, params)
}

func (s *service) UpdateRoom(ctx context.Context, room *Room) error {
	return s.repo.UpdateRoom(ctx, room)
}

func (s *service) DeleteRoom(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteRoom(ctx, id)
}
