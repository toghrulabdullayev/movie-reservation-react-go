package movies

import (
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetMovies(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "GetMovies endpoint",
	})
}

func (h *Handler) GetMovieByID(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "GetMovieByID endpoint",
	})
}

func (h *Handler) CreateMovie(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "CreateMovie endpoint",
	})
}

func (h *Handler) UpdateMovie(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "UpdateMovie endpoint",
	})
}

func (h *Handler) DeleteMovie(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DeleteMovie endpoint",
	})
}
