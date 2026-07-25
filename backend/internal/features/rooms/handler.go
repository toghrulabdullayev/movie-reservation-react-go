package rooms

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/shared/pagination"
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
// @Summary      Create a new room
// @Description  Create a new room with the provided details.
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        room  body      CreateRoomRequest  true  "Room details"
// @Success      201   {object}  RoomResponse
// @Failure      400   {object}  map[string]string "Bad Request - Invalid input"
// @Failure      500   {object}  map[string]string "Internal Server Error - Failed to create room"
// @Router       /rooms [post]
func (h *Handler) CreateRoom(c fiber.Ctx) error {
	var req CreateRoomRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	room, err := h.service.CreateRoom(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create room",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(toRoomResponse(room))
}

// * Swagger annotations
// @Summary      Get room by ID
// @Description  Retrieve the details of a single room by its UUID.
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Room UUID" format(uuid)
// @Success      200  {object}  RoomResponse
// @Failure      400  {object}  map[string]string "Bad Request - Invalid ID format"
// @Failure      404  {object}  map[string]string "Not Found - Room not found"
// @Router       /rooms/{id} [get]
func (h *Handler) GetRoomByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid room ID",
		})
	}

	room, err := h.service.GetRoomByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Room not found",
		})
	}

	return c.JSON(toRoomResponse(room))
}

// * Swagger annotations
// @Summary      Get all rooms with pagination
// @Description  Retrieve a paginated list of all rooms.
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        page   query      int  false  "Page number (default: 1)"
// @Param        limit  query      int  false  "Number of items per page (default: 10, max: 100)"
// @Success      200    {object}   pagination.PaginatedResponse[RoomResponse]
// @Failure      500    {object}   map[string]string "Internal Server Error - Failed to retrieve rooms"
// @Router       /rooms [get]
func (h *Handler) GetAllRooms(c fiber.Ctx) error {
	params := pagination.GetPaginationParams(c, 10, 100)

	rooms, totalRecords, err := h.service.GetAllRooms(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve rooms",
		})
	}

	roomResponses := make([]RoomResponse, 0, len(rooms))
	for _, room := range rooms {
		roomResponses = append(roomResponses, toRoomResponse(room))
	}

	paginatedResponse := pagination.NewPaginatedResponse(roomResponses, totalRecords, params)

	return c.JSON(paginatedResponse)
}

// * Swagger annotations
// @Summary      Update room by ID
// @Description  Update the details of a room by its UUID.
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        id    path      string             true  "Room UUID" format(uuid)
// @Param        room  body      UpdateRoomRequest  true  "Updated room details"
// @Success      200   {object}  RoomResponse
// @Failure      400   {object}  map[string]string "Bad Request - Invalid input or ID format"
// @Failure      404   {object}  map[string]string "Not Found - Room not found"
// @Failure      500   {object}  map[string]string "Internal Server Error - Failed to update room"
// @Router       /rooms/{id} [put]
func (h *Handler) UpdateRoom(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid room ID",
		})
	}

	var req UpdateRoomRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	room := &Room{
		ID:       id,
		Name:     req.Name,
		Capacity: req.Capacity,
		Rows:     req.Rows,
		Columns:  req.Columns,
	}

	if err := h.service.UpdateRoom(c.Context(), room); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update room",
		})
	}

	return c.JSON(toRoomResponse(room))
}

// * Swagger annotations
// @Summary      Delete room by ID
// @Description  Delete a room by its UUID.
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Room UUID" format(uuid)
// @Success      204  {object}  nil "No Content - Room deleted successfully"
// @Failure      400  {object}  map[string]string "Bad Request - Invalid ID format"
// @Failure      404  {object}  map[string]string "Not Found - Room not found"
// @Failure      500  {object}  map[string]string "Internal Server Error - Failed to delete room"
// @Router       /rooms/{id} [delete]
func (h *Handler) DeleteRoom(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid room ID",
		})
	}

	if err := h.service.DeleteRoom(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete room",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
