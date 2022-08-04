package service

import (
	"mansion/repository"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/validator.v2"
)

type roomTypeService struct {
	roomTypeRepo repository.RoomTypeRepository
}

func NewRoomTypeService(roomTypeRepo repository.RoomTypeRepository) RoomTypeService {
	return roomTypeService{roomTypeRepo: roomTypeRepo}
}

func (s roomTypeService) CreateRoomType(roomType *RoomTypeRequest) (*RoomTypeResponse, error) {

	if errs := validator.Validate(roomType); errs != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Requested body incorrect form.")
	}

	newRoom := repository.RoomType{
		Name: roomType.Name,
	}

	result, err := s.roomTypeRepo.Create(&newRoom)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Can not add new room.")
	}

	return &RoomTypeResponse{
		UUID:      result.UUID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (s roomTypeService) FetchRoomTypes() ([]RoomTypeResponse, error) {

	roomTypes, err := s.roomTypeRepo.GetAll()

	if err != nil {
		return nil, err
	}

	roomTypeResponses := []RoomTypeResponse{}

	for _, room := range roomTypes {
		roomTypeResponses = append(roomTypeResponses, RoomTypeResponse{
			UUID:      room.UUID,
			Name:      room.Name,
			CreatedAt: room.CreatedAt,
			UpdatedAt: room.UpdatedAt,
		})
	}
	return roomTypeResponses, nil
}

func (s roomTypeService) GetRoomType(id string) (*RoomTypeResponse, error) {
	room, err := s.roomTypeRepo.GetById(id)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found room.")
	}

	return &RoomTypeResponse{
		UUID:      room.UUID,
		Name:      room.Name,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}, nil
}

func (s roomTypeService) UpdateRoomType(id string, roomType *RoomTypeRequest) error {

	if errs := validator.Validate(roomType); errs != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Requested body incorrect form.")
	}

	updateRoom := repository.UpdateRoomType{
		Name: roomType.Name,
	}

	err := s.roomTypeRepo.Update(id, &updateRoom)

	if err != nil {
		return err
	}

	return nil
}

func (s roomTypeService) DeleteRoomType(id string) error {
	err := s.roomTypeRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
