package service

import "time"

type RoomTypeResponse struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RoomTypeRequest struct {
	Name string `validate:"nonzero"`
}

type RoomTypeService interface {
	FetchRoomTypes() ([]RoomTypeResponse, error)
	GetRoomType(id string) (*RoomTypeResponse, error)
	CreateRoomType(roomType *RoomTypeRequest) (*RoomTypeResponse, error)
	UpdateRoomType(id string, room *RoomTypeRequest) error
	DeleteRoomType(id string) error
}
