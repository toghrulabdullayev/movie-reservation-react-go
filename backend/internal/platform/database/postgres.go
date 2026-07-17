package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
)

var db *pgxpool.Pool

func PgConnect() {
	var err error

	host := config.GetEnv("POSTGRES_HOST")
	port := config.GetEnv("POSTGRES_PORT")
	user := config.GetEnv("POSTGRES_USER")
	password := config.GetEnv("POSTGRES_PASSWORD")
	dbname := config.GetEnv("POSTGRES_DB")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	db, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	log.Println("Connected to the PostgreSQL database successfully!")
}

func PgGetDB() *pgxpool.Pool {
	return db
}

func PgClose() {
	if db != nil {
		db.Close()
	}
}
