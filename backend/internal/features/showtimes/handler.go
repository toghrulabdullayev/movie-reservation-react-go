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

// showtimes
// * Swagger annotations
// @Summary      Get all showtimes
// @Description  Get a list of all showtimes currently registered in the system.
// @Tags         showtimes
// @Accept       json
// @Produce      json
// @Success      200  {array}   ShowtimeResponse
// @Failure      500  {object}  map[string]string "Internal Server Error - Failed to retrieve showtimes"
// @Router       /showtimes [get]
func (h *Handler) GetAllShowtimes(c fiber.Ctx) error {
	showtimes, err := h.service.GetAllShowtimes(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve showtimes",
		})
	}

	showtimeResponses := make([]ShowtimeResponse, len(showtimes))
	for i, showtime := range showtimes {
		showtimeResponses[i] = toShowtimeResponse(showtime)
	}

	return c.JSON(showtimeResponses)
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
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid showtime ID",
		})
	}

	var req UpdateShowtimeRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err = h.service.UpdateShowtime(c.Context(), id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update showtime",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Showtime updated successfully",
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
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid showtime ID",
		})
	}

	err = h.service.DeleteShowtime(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete showtime",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Showtime deleted successfully",
	})
}

// showtime seats
// * Swagger annotations
// @Summary      Create a showtime seat
// @Description  Create a new showtime seat entry for a specific showtime.
// @Tags         showtimes-seats
// @Accept       json
// @Produce      json
// @Param        request  body      CreateShowtimeSeatRequest  true  "Create showtime seat payload"
// @Success      201      {object}  map[string]string "Success"
// @Failure      400      {object}  map[string]string "Bad Request - Invalid payload"
// @Failure      500      {object}  map[string]string "Internal Server Error"
// @Router       /showtimes/seats [post]
func (h *Handler) CreateShowtimeSeat(c fiber.Ctx) error {
	var req CreateShowtimeSeatRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.service.CreateShowtimeSeat(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create showtime seat",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Showtime seat created successfully",
	})
}

// * Swagger annotations
// @Summary      Get showtime seat by ID
// @Description  Retrieve the details of a single showtime seat by its UUID.
// @Tags         showtimes-seats
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Showtime Seat UUID" format(uuid)
// @Success      200  {object}  ShowtimeSeatResponse
// @Failure      400  {object}  map[string]string "Bad Request - Invalid ID format"
// @Failure      404  {object}  map[string]string "Not Found - Showtime seat not found"
// @Router       /showtimes/seats/{id} [get]
func (h *Handler) GetShowtimeSeatByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid showtime seat ID",
		})
	}

	seat, err := h.service.GetShowtimeSeatByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Showtime seat not found",
		})
	}

	return c.JSON(toShowtimeSeatResponse(seat))
}

// * Swagger annotations
// @Summary      Get all showtime seats
// @Description  Get a list of all showtime seats currently registered in the system.
// @Tags         showtimes-seats
// @Accept       json
// @Produce      json
// @Success      200  {array}   ShowtimeSeatResponse
// @Failure      500  {object}  map[string]string "Internal Server Error - Failed to retrieve showtime seats"
// @Router       /showtimes/seats [get]
func (h *Handler) GetAllShowtimeSeats(c fiber.Ctx) error {
	seats, err := h.service.GetAllShowtimeSeats(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve showtime seats",
		})
	}

	seatResponses := make([]ShowtimeSeatResponse, len(seats))
	for i, seat := range seats {
		seatResponses[i] = toShowtimeSeatResponse(seat)
	}

	return c.JSON(seatResponses)
}

// * Swagger annotations
// @Summary      Update a showtime seat
// @Description  Modify parameters of a showtime seat by its ID.
// @Tags         showtimes-seats
// @Accept       json
// @Produce      json
// @Param        id       path      string                     true  "Showtime Seat UUID" format(uuid)
// @Param        request  body      UpdateShowtimeSeatRequest  true  "Update showtime seat payload"
// @Success      200      {object}  map[string]string "Success"
// @Router       /showtimes/seats/{id} [put]
func (h *Handler) UpdateShowtimeSeat(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid showtime seat ID",
		})
	}

	var req UpdateShowtimeSeatRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err = h.service.UpdateShowtimeSeat(c.Context(), id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update showtime seat",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Showtime seat updated successfully",
	})
}

// * Swagger annotations
// @Summary      Delete a showtime seat
// @Description  Remove a showtime seat from the database by its ID.
// @Tags         showtimes-seats
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Showtime Seat UUID" format(uuid)
// @Success      200  {object}  map[string]string "Success"
// @Router       /showtimes/seats/{id} [delete]
func (h *Handler) DeleteShowtimeSeat(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid showtime seat ID",
		})
	}

	err = h.service.DeleteShowtimeSeat(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete showtime seat",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Showtime seat deleted successfully",
	})
}
