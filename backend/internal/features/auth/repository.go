package auth

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByIdentity(ctx context.Context, identity string) (*User, error)
	GetUserByGoogleID(ctx context.Context, googleID string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, u *User) error {
	query := `
		INSERT INTO users (id, username, email, password_hash, role, google_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
	`
	_, err := r.db.Exec(ctx, query, u.ID, u.Username, u.Email, u.PasswordHash, u.Role, u.GoogleID)
	return err
}

func (r *repository) GetUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	query := `SELECT id, username, email, password_hash, role, google_id, created_at, updated_at FROM users WHERE id = $1`
	var u User
	err := r.db.QueryRow(ctx, query, id).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Role, &u.GoogleID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := `SELECT id, username, email, password_hash, role, google_id, created_at, updated_at FROM users WHERE email = $1`
	var u User
	err := r.db.QueryRow(ctx, query, email).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Role, &u.GoogleID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	query := `SELECT id, username, email, password_hash, role, google_id, created_at, updated_at FROM users WHERE username = $1`
	var u User
	err := r.db.QueryRow(ctx, query, username).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Role, &u.GoogleID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *repository) GetUserByIdentity(ctx context.Context, identity string) (*User, error) {
	query := `SELECT id, username, email, password_hash, role, google_id, created_at, updated_at FROM users WHERE email = $1 OR username = $2`
	var u User
	err := r.db.QueryRow(ctx, query, identity, identity).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Role, &u.GoogleID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *repository) GetUserByGoogleID(ctx context.Context, googleID string) (*User, error) {
	query := `SELECT id, username, email, password_hash, role, google_id, created_at, updated_at FROM users WHERE google_id = $1`
	var u User
	err := r.db.QueryRow(ctx, query, googleID).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Role, &u.GoogleID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *repository) UpdateUser(ctx context.Context, u *User) error {
	query := `
		UPDATE users
		SET username = $1, email = $2, password_hash = $3, role = $4, google_id = $5, updated_at = NOW()
		WHERE id = $6
	`
	_, err := r.db.Exec(ctx, query, u.Username, u.Email, u.PasswordHash, u.Role, u.GoogleID, u.ID)
	return err
}
