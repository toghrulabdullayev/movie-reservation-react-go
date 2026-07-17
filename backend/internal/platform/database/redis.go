package database

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
)

var rdb *redis.Client

func RedisConnect() {
	host := config.GetEnv("REDIS_HOST")
	port := config.GetEnv("REDIS_PORT")

	addr := fmt.Sprintf("%s:%s", host, port)

	opts := &redis.Options{
		Addr: addr,
	}

	rdb = redis.NewClient(opts)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Unable to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully!")
}

func RedisGet() *redis.Client {
	return rdb
}

func RedisClose() {
	if rdb != nil {
		if err := rdb.Close(); err != nil {
			log.Printf("Failed to close Redis: %v", err)
		}
	}
}
