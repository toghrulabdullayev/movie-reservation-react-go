// START OF FILE backend/cmd/api/main.go
package main

import (
	"context"
	"log" // Standard library log for bootstrap fallbacks

	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/database"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/server"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/shared/logger"
)

func main() {
	// If run fails (such as failing to load config), we log to stderr and exit
	if err := run(); err != nil {
		log.Fatalf("Application startup failed: %v", err)
	}
}

func run() error {
	ctx := context.Background()

	// 1. Load config (if this fails, it returns back to main's standard logger)
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// 2. Instantiate our structured logger
	logInstance := logger.New(cfg.Env)
	logInstance.Info().Str("env", cfg.Env).Msg("Logger initialized successfully")

	// 3. Connect to PostgreSQL
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

	// 4. Connect to Redis
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

	// 5. Build and launch server
	deps := server.BuildDependencies(pool)
	srv := server.New(cfg, deps, &logInstance)

	logInstance.Info().Str("port", cfg.Port).Msg("Starting application server...")
	return srv.Start(func() {
		logInstance.Info().Msg("Application server started successfully")
	})
}
