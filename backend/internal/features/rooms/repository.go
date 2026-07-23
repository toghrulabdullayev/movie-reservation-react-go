package rooms

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/shared/pagination"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	CreateRoom(ctx context.Context, room *Room) error
	GetRoomByID(ctx context.Context, id uuid.UUID) (*Room, error)
	GetAllRooms(ctx context.Context, params pagination.PaginationParams) ([]*Room, int, error)
	UpdateRoom(ctx context.Context, room *Room) error
	DeleteRoom(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateRoom(ctx context.Context, room *Room) error {
	query := `
		INSERT INTO rooms (id, name, capacity, rows, columns, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())`

	_, err := r.db.Exec(
		ctx,
		query,
		room.ID,
		room.Name,
		room.Capacity,
		room.Rows,
		room.Columns,
	)

	return err
}

func (r *repository) GetRoomByID(ctx context.Context, id uuid.UUID) (*Room, error) {
	query := `
		SELECT id, name, capacity, rows, columns, created_at, updated_at 
		FROM rooms 
		WHERE id = $1`

	var room Room
	err := r.db.QueryRow(ctx, query, id).Scan(
		&room.ID,
		&room.Name,
		&room.Capacity,
		&room.Rows,
		&room.Columns,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *repository) GetAllRooms(ctx context.Context, params pagination.PaginationParams) ([]*Room, int, error) {
	totalRecordsQuery := `SELECT COUNT(*) FROM rooms`
	query := `SELECT id, name, capacity, rows, columns, created_at, updated_at FROM rooms ORDER BY name LIMIT $1 OFFSET $2`

	totalRecords := 0
	err := r.db.QueryRow(ctx, totalRecordsQuery).Scan(&totalRecords)
	if err != nil {
		return nil, 0, err
	}

	if totalRecords == 0 {
		return []*Room{}, 0, nil
	}

	rows, err := r.db.Query(ctx, query, params.Limit, params.Offset())
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []*Room

	for rows.Next() {
		var room Room
		err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.Capacity,
			&room.Rows,
			&room.Columns,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, &room)
	}

	return list, totalRecords, nil
}

func (r *repository) UpdateRoom(ctx context.Context, room *Room) error {
	query := `
		UPDATE rooms 
		SET name = $1, capacity = $2, rows = $3, columns = $4, updated_at = NOW() 
		WHERE id = $5`

	_, err := r.db.Exec(
		ctx,
		query,
		room.Name,
		room.Capacity,
		room.Rows,
		room.Columns,
		room.ID,
	)

	return err
}

func (r *repository) DeleteRoom(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM rooms WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	return err
}
