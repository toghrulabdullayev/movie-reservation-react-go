package movies

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid movie ID",
		})
	}

	movie, err := h.service.GetMovieByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Movie not found",
		})
	}

	return c.JSON(toMovieResponse(movie))
}

func (h *Handler) CreateMovie(c fiber.Ctx) error {
	var req CreateMovieRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdMovie, err := h.service.CreateMovie(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create movie",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(toMovieResponse(createdMovie))
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
