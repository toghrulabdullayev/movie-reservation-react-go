package showtimes

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

func (h *Handler) GetAllShowtimes(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "GetAllShowtimes endpoint",
	})
}

func (h *Handler) GetShowtimeByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid movie ID",
		})
	}

	movie, err := h.service.GetShowtimeByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Movie not found",
		})
	}

	return c.JSON(toShowtimeResponse(movie))
}

func (h *Handler) CreateShowtime(c fiber.Ctx) error {
	var req CreateShowtimeRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdMovie, err := h.service.CreateShowtime(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create movie",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(toShowtimeResponse(createdMovie))
}

func (h *Handler) UpdateShowtime(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "UpdateShowtime endpoint",
	})
}

func (h *Handler) DeleteShowtime(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DeleteShowtime endpoint",
	})
}

func (h *Handler) CreateDirector(c fiber.Ctx) error {
	var req CreateDirectorRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdDirector, err := h.service.CreateDirector(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create director",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdDirector)
}
