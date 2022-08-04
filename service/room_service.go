package service

import (
	"mansion/repository"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/validator.v2"
)

type roomService struct {
	roomRepo repository.RoomRepository
}

func NewRoomService(roomRepo repository.RoomRepository) RoomService {
	return roomService{roomRepo: roomRepo}
}

func (s roomService) CreateRoom(room *RoomRequest) (*RoomResponse, error) {

	if errs := validator.Validate(room); errs != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Requested body incorrect form.")
	}

	newRoom := repository.Room{
		RoomID:     room.RoomID,
		RoomTypeID: room.RoomTypeID,
		RoomPrice:  room.RoomPrice,
		Floor:      room.Floor,
	}

	result, err := s.roomRepo.Create(&newRoom)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Can not add new room.")
	}

	return &RoomResponse{
		UUID:      result.UUID,
		RoomID:    result.RoomID,
		RoomPrice: result.RoomPrice,
		Floor:     result.Floor,
		RoomType: RoomTypeResponse{
			UUID:      result.RoomType.UUID,
			Name:      result.RoomType.Name,
			CreatedAt: result.RoomType.CreatedAt,
			UpdatedAt: result.RoomType.UpdatedAt,
		},
	}, nil
}

func (s roomService) FetchRooms() ([]RoomResponse, error) {

	rooms, err := s.roomRepo.GetAll()

	if err != nil {
		return nil, err
	}

	roomResponses := []RoomResponse{}

	for _, room := range rooms {
		roomResponses = append(roomResponses, RoomResponse{
			UUID:      room.UUID,
			RoomID:    room.RoomID,
			RoomPrice: room.RoomPrice,
			Floor:     room.Floor,
			RoomType: RoomTypeResponse{
				UUID:      room.RoomType.UUID,
				Name:      room.RoomType.Name,
				CreatedAt: room.RoomType.CreatedAt,
				UpdatedAt: room.RoomType.UpdatedAt,
			},
		})
	}
	return roomResponses, nil
}

func (s roomService) GetRoom(id string) (*RoomResponse, error) {
	room, err := s.roomRepo.GetById(id)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found room.")
	}

	return &RoomResponse{
		UUID:      room.UUID,
		RoomID:    room.RoomID,
		RoomPrice: room.RoomPrice,
		Floor:     room.Floor,
		RoomType: RoomTypeResponse{
			UUID:      room.RoomType.UUID,
			Name:      room.RoomType.Name,
			CreatedAt: room.RoomType.CreatedAt,
			UpdatedAt: room.RoomType.UpdatedAt,
		},
	}, nil
}

func (s roomService) UpdateRoom(id string, room *RoomRequest) error {

	if errs := validator.Validate(room); errs != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Requested body incorrect form.")
	}

	updateRoom := repository.UpdateRoom{
		RoomID:     room.RoomID,
		RoomPrice:  room.RoomPrice,
		RoomTypeID: room.RoomTypeID,
		Floor:      room.Floor,
	}

	err := s.roomRepo.Update(id, &updateRoom)

	if err != nil {
		return err
	}

	return nil
}

func (s roomService) DeleteRoom(id string) error {
	err := s.roomRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
