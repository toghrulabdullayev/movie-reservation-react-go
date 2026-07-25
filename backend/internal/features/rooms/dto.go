package rooms

import "github.com/google/uuid"

type RoomResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Capacity int       `json:"capacity"`
	Rows     int       `json:"rows"`
	Columns  int       `json:"columns"`
}

type CreateRoomRequest struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Rows     int    `json:"rows"`
	Columns  int    `json:"columns"`
}

type UpdateRoomRequest struct {
	Name     string    `json:"name"`
	Capacity int       `json:"capacity"`
	Rows     int       `json:"rows"`
	Columns  int       `json:"columns"`
}

func toRoomResponse(r *Room) RoomResponse {
	if r == nil {
		return RoomResponse{}
	}

	return RoomResponse{
		ID:       r.ID,
		Name:     r.Name,
		Capacity: r.Capacity,
		Rows:     r.Rows,
		Columns:  r.Columns,
	}
}
