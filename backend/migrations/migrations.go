package migrations

import (
	"context"
	"embed"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var EmbedFS embed.FS

// RunMigrations connects to the database via standard library wrapper and runs pending migrations.
func RunMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	// Set the driver dialect
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	// Supply the embedded filesystem to Goose
	goose.SetBaseFS(EmbedFS)

	// Wrap the pgx connection pool to make it compatible with database/sql
	db := stdlib.OpenDBFromPool(pool)
	defer func() {
		_ = db.Close()
	}()

	// Run all pending migrations in the root embed directory
	if err := goose.UpContext(ctx, db, "."); err != nil {
		return fmt.Errorf("goose migrations failed: %w", err)
	}

	return nil
}
