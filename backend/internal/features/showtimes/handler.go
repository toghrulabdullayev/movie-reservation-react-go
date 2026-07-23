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

// * Swagger annotations
// GetAllShowtimes retrieves all scheduled showtimes.
// @Summary      Get all showtimes
// @Description  Get a list of all showtimes currently registered in the system.
// @Tags         showtimes
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string "Success"
// @Router       /showtimes [get]
func (h *Handler) GetAllShowtimes(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "GetAllShowtimes endpoint",
	})
}

// * Swagger annotations
// @Summary      Get showtime by ID
// @Description  Retrieve the details of a single scheduled showtime by its UUID.
// @Tags         showtimes
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Showtime UUID" format(uuid)
// @Success      200  {object}  ShowtimeResponse
// @Failure      400  {object}  map[string]string "Bad Request - Invalid ID format"
// @Failure      404  {object}  map[string]string "Not Found - Showtime not found"
// @Router       /showtimes/{id} [get]
func (h *Handler) GetShowtimeByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid showtime ID",
		})
	}

	movie, err := h.service.GetShowtimeByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Showtime not found",
		})
	}

	return c.JSON(toShowtimeResponse(movie))
}

// * Swagger annotations
// @Summary      Create a showtime
// @Description  Create a new showtime entry for a specific movie inside a specific theatre room.
// @Tags         showtimes
// @Accept       json
// @Produce      json
// @Param        request  body      CreateShowtimeRequest  true  "Create showtime payload"
// @Success      201      {object}  ShowtimeResponse
// @Failure      400      {object}  map[string]string "Bad Request - Invalid payload"
// @Failure      500      {object}  map[string]string "Internal Server Error"
// @Router       /showtimes [post]
func (h *Handler) CreateShowtime(c fiber.Ctx) error {
	var req CreateShowtimeRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdShowtime, err := h.service.CreateShowtime(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create showtime",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(toShowtimeResponse(createdShowtime))
}

// * Swagger annotations
// @Summary      Update a showtime
// @Description  Modify parameters of a scheduled showtime by its ID.
// @Tags         showtimes
// @Accept       json
// @Produce      json
// @Param        id       path      string                 true  "Showtime UUID" format(uuid)
// @Param        request  body      UpdateShowtimeRequest  true  "Update showtime payload"
// @Success      200      {object}  map[string]string "Success"
// @Router       /showtimes/{id} [put]
func (h *Handler) UpdateShowtime(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "UpdateShowtime endpoint",
	})
}

// * Swagger annotations
// @Summary      Delete a showtime
// @Description  Remove a scheduled showtime from the database by its ID.
// @Tags         showtimes
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Showtime UUID" format(uuid)
// @Success      200  {object}  map[string]string "Success"
// @Router       /showtimes/{id} [delete]
func (h *Handler) DeleteShowtime(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DeleteShowtime endpoint",
	})
}
