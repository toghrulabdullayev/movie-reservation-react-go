package main

import (
	"context"
	"log"

	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/database"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	pool, err := database.NewPostgresPool(ctx, cfg.PostgresDSN())
	if err != nil {
		return err
	}
	defer pool.Close()

	rdb, err := database.NewRedisClient(ctx, cfg.RedisAddr())
	if err != nil {
		return err
	}
	defer func() {
		if err := rdb.Close(); err != nil {
			log.Printf("failed to close redis client: %v", err)
		}
	}()

	deps := server.BuildDependencies(pool)
	srv := server.New(cfg, deps)

	return srv.Start()
}
