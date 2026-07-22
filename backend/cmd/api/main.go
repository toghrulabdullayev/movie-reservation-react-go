package main

import (
	"context"
	// "fmt"
	"log"

	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/database"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/server"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/shared/logger"
	// "github.com/toghrulabdullayev/movie-reservation-react-go/backend/migrations"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Application startup failed: %v", err)
	}
}

func run() error {
	ctx := context.Background()

	// 1. load .env configuration
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// 2. initialize zerolog logger
	logInstance := logger.New(cfg.Env)
	logInstance.Info().Str("env", cfg.Env).Msg("Logger initialized successfully")

	// 3. postgres
	logInstance.Info().Msg("Connecting to PostgreSQL...")
	pool, err := database.NewPostgresPool(ctx, cfg.PostgresDSN())
	if err != nil {
		return err
	}

	logInstance.Info().Msg("Successfully connected to PostgreSQL")
	defer func() {
		logInstance.Info().Msg("Closing PostgreSQL connection pool...")
		pool.Close()
	}()

	// // run migrations
	// logInstance.Info().Msg("Running database migrations...")
	// if err := migrations.RunMigrations(ctx, pool); err != nil {
	// 	return fmt.Errorf("database migrations failed: %w", err)
	// }
	// logInstance.Info().Msg("Database migrations applied successfully")

	// 4. redis
	logInstance.Info().Msg("Connecting to Redis...")
	rdb, err := database.NewRedisClient(ctx, cfg.RedisAddr())
	if err != nil {
		return err
	}

	logInstance.Info().Msg("Successfully connected to Redis")
	defer func() {
		logInstance.Info().Msg("Closing Redis connection client...")
		if err := rdb.Close(); err != nil {
			logInstance.Error().Err(err).Msg("Failed to close Redis client gracefully")
		}
	}()

	// 5. build and launch the server
	deps := server.BuildDependencies(cfg, pool, rdb)
	srv := server.New(cfg, deps, &logInstance)

	logInstance.Info().Str("port", cfg.Port).Msg("Starting application server...")
	return srv.Start()
}
