package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
)

type Server struct {
	app *fiber.App
	cfg *config.Config
}

func New(cfg *config.Config, deps *Dependencies) *Server {
	app := fiber.New(fiber.Config{})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	registerRoutes(app, deps)

	return &Server{app: app, cfg: cfg}
}

func (s *Server) Start() error {
	return s.app.Listen(":" + s.cfg.Port)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
