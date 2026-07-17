package server

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/movies"
)

type Dependencies struct {
	MovieHandler *movies.Handler
}

func BuildDependencies(pgPool *pgxpool.Pool) *Dependencies {
	movieRepository := movies.NewRepository(pgPool)
	movieService := movies.NewService(movieRepository)
	movieHandler := movies.NewHandler(movieService)

	return &Dependencies{
		MovieHandler: movieHandler,
	}
}
