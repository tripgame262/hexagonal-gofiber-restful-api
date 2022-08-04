package service

type RoomResponse struct {
	UUID      string           `json:"uuid"`
	RoomID    string           `json:"roomId"`
	RoomPrice int16            `json:"roomPrice"`
	Floor     int16            `json:"floor"`
	RoomType  RoomTypeResponse `json:"roomType"`
}

type RoomRequest struct {
	RoomID     string `json:"roomId" validate:"nonzero"`
	RoomPrice  int16  `json:"roomPrice" validate:"nonzero"`
	Floor      int16  `json:"floor" validate:"nonzero"`
	RoomTypeID string `json:"roomType" validate:"nonzero"`
}

type RoomService interface {
	FetchRooms() ([]RoomResponse, error)
	GetRoom(id string) (*RoomResponse, error)
	CreateRoom(room *RoomRequest) (*RoomResponse, error)
	UpdateRoom(id string, room *RoomRequest) error
	DeleteRoom(id string) error
}
