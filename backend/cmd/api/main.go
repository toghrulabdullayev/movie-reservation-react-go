package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/database"
	// "github.com/gofiber/fiber/v3/middleware/static"
)

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

func main() {
	// load .env file
	config.LoadEnv()

	port := fmt.Sprintf(":%s", config.GetEnv("PORT"))
	isProd := config.GetEnv("ENV") == EnvProduction

	// todo: connect to database
	database.PgConnect()
	defer database.PgClose()

	database.RedisConnect()
	defer database.RedisClose()

	// Create fiber app
	app := fiber.New(fiber.Config{})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	v1.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "hello world",
		})
	})

	// // Bind handlers
	// v1.Get("/users", handlers.UserList)
	// v1.Post("/users", handlers.UserCreate)

	// // Setup static files
	// app.Get("/*", static.New("./static/public"))

	// // Handle not founds
	// app.Use(handlers.NotFound)

	// Listen on port 3000
	if err := app.Listen(port, fiber.ListenConfig{EnablePrefork: isProd}); err != nil {
		log.Printf("server error: %v", err)
	}
}
