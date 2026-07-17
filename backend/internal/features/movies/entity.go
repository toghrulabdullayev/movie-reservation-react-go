package movies

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Movie struct {
	BaseEntity
	DirectorID  uuid.UUID
	GenreID     uuid.UUID
	Title       string
	ReleaseYear int
	Rating      float32
	Plot        string
	Duration    int
}

type Director struct {
	BaseEntity
	Name string
}

type Genre struct {
	BaseEntity
	Name string
}
