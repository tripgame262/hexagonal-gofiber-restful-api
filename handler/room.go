package handler

import (
	"mansion/service"

	"github.com/gofiber/fiber/v2"
)

type roomHandler struct {
	roomSrv service.RoomService
}

func NewRoomHandler(roomSrv service.RoomService) roomHandler {
	return roomHandler{roomSrv: roomSrv}
}

func (h roomHandler) FetchRooms(c *fiber.Ctx) error {
	result, err := h.roomSrv.FetchRooms()

	if err != nil {
		return handleError(c, err)
	}

	return handlerOk(c, result)
}

func (h roomHandler) GetRoom(c *fiber.Ctx) error {
	id := c.Params("id")

	room, err := h.roomSrv.GetRoom(id)

	if err != nil {
		return handleError(c, err)
	}

	return handlerOk(c, *room)
}

func (h roomHandler) CreateRoom(c *fiber.Ctx) error {
	newRoom := service.RoomRequest{}

	err := c.BodyParser(&newRoom)

	if err != nil {
		return handleError(c, fiber.NewError(fiber.StatusBadRequest, "Requested body incorrect form"))
	}

	result, err := h.roomSrv.CreateRoom(&newRoom)

	if err != nil {
		return handleError(c, err)
	}

	return handlerCreatedOk(c, *result)
}

func (h roomHandler) UpdateRoom(c *fiber.Ctx) error {
	updateRoom := service.RoomRequest{}

	id := c.Params("id")

	err := c.BodyParser(&updateRoom)

	if err != nil {
		return handleError(c, fiber.NewError(fiber.StatusBadRequest, "Requested body incorrect form"))
	}

	err = h.roomSrv.UpdateRoom(id, &updateRoom)

	if err != nil {
		return handleError(c, err)
	}

	return handlerUpdateOk(c)
}

func (h roomHandler) DeleteRoom(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.roomSrv.DeleteRoom(id)

	if err != nil {
		return handleError(c, err)
	}

	return handlerDeleteOk(c)
}
