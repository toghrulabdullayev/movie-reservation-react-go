package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/rs/zerolog"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/middleware"
)

type Server struct {
	app *fiber.App
	cfg *config.Config
}

func New(cfg *config.Config, deps *Dependencies, log *zerolog.Logger) *Server {
	app := fiber.New(fiber.Config{})
	app.Use(recover.New())
	app.Use(middleware.Logger(log)) // custom zerolog middleware
	app.Use(cors.New())

	registerRoutes(app, deps)

	return &Server{app: app, cfg: cfg}
}

func (s *Server) Start(cb func()) error {
	cb()
	return s.app.Listen(":" + s.cfg.Port)
}

func (s *Server) Shutdown(cb func()) error {
	cb()
	return s.app.Shutdown()
}
