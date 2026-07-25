package seats

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// compile-time check to ensure repository implements Repository.
var _ Repository = (*repository)(nil)

type Repository interface {
	// seats
	CreateSeat(ctx context.Context, seat *Seat) error
	GetSeatByID(ctx context.Context, id uuid.UUID) (*Seat, error)
	GetAllSeats(ctx context.Context) ([]*Seat, error)
	UpdateSeat(ctx context.Context, seat *Seat) error
	DeleteSeat(ctx context.Context, id uuid.UUID) error

	// seat types
	CreateSeatType(ctx context.Context, seatType *SeatType) error
	GetSeatTypeByID(ctx context.Context, id uuid.UUID) (*SeatType, error)
	GetAllSeatTypes(ctx context.Context) ([]*SeatType, error)
	UpdateSeatType(ctx context.Context, seatType *SeatType) error
	DeleteSeatType(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

// seats
func (r *repository) CreateSeat(ctx context.Context, seat *Seat) error {
	query := `
		INSERT INTO seats (id, room_id, seat_type_id, row, number, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())`

	_, err := r.db.Exec(
		ctx,
		query,
		seat.ID,
		seat.RoomID,
		seat.SeatTypeID,
		seat.Row,
		seat.Number,
	)

	return err
}

func (r *repository) GetSeatByID(ctx context.Context, id uuid.UUID) (*Seat, error) {
	query := `
		SELECT id, room_id, seat_type_id, row, number, created_at, updated_at 
		FROM seats 
		WHERE id = $1`

	var seat Seat

	err := r.db.QueryRow(ctx, query, id).Scan(
		&seat.ID,
		&seat.RoomID,
		&seat.SeatTypeID,
		&seat.Row,
		&seat.Number,
		&seat.CreatedAt,
		&seat.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &seat, nil
}

func (r *repository) GetAllSeats(ctx context.Context) ([]*Seat, error) {
	query := `
		SELECT id, room_id, seat_type_id, row, number, created_at, updated_at 
		FROM seats`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []*Seat

	for rows.Next() {
		var seat Seat
		err := rows.Scan(
			&seat.ID,
			&seat.RoomID,
			&seat.SeatTypeID,
			&seat.Row,
			&seat.Number,
			&seat.CreatedAt,
			&seat.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		seats = append(seats, &seat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}

func (r *repository) UpdateSeat(ctx context.Context, seat *Seat) error {
	query := `
		UPDATE seats 
		SET room_id = $1, seat_type_id = $2, row = $3, number = $4, updated_at = NOW() 
		WHERE id = $5`

	_, err := r.db.Exec(
		ctx,
		query,
		seat.RoomID,
		seat.SeatTypeID,
		seat.Row,
		seat.Number,
		seat.ID,
	)

	return err
}

func (r *repository) DeleteSeat(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM seats WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

// seat types
func (r *repository) CreateSeatType(ctx context.Context, seatType *SeatType) error {
	query := `
		INSERT INTO seat_types (id, name, multiplier, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())`

	_, err := r.db.Exec(
		ctx,
		query,
		seatType.ID,
		seatType.Name,
		seatType.Multiplier,
	)

	return err
}

func (r *repository) GetSeatTypeByID(ctx context.Context, id uuid.UUID) (*SeatType, error) {
	query := `
		SELECT id, name, multiplier, created_at, updated_at 
		FROM seat_types 
		WHERE id = $1`

	var seatType SeatType

	err := r.db.QueryRow(ctx, query, id).Scan(
		&seatType.ID,
		&seatType.Name,
		&seatType.Multiplier,
		&seatType.CreatedAt,
		&seatType.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &seatType, nil
}

func (r *repository) GetAllSeatTypes(ctx context.Context) ([]*SeatType, error) {
	query := `
		SELECT id, name, multiplier, created_at, updated_at 
		FROM seat_types`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seatTypes []*SeatType

	for rows.Next() {
		var seatType SeatType
		err := rows.Scan(
			&seatType.ID,
			&seatType.Name,
			&seatType.Multiplier,
			&seatType.CreatedAt,
			&seatType.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		seatTypes = append(seatTypes, &seatType)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return seatTypes, nil
}

func (r *repository) UpdateSeatType(ctx context.Context, seatType *SeatType) error {
	query := `
		UPDATE seat_types 
		SET name = $1, multiplier = $2, updated_at = NOW() 
		WHERE id = $3`

	_, err := r.db.Exec(
		ctx,
		query,
		seatType.Name,
		seatType.Multiplier,
		seatType.ID,
	)

	return err
}

func (r *repository) DeleteSeatType(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM seat_types WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
