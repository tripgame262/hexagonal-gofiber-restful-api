package handler

import (
	"mansion/service"

	"github.com/gofiber/fiber/v2"
)

type roomTypeHandler struct {
	roomTypeSrv service.RoomTypeService
}

func NewRoomTypeHandler(roomTypeSrv service.RoomTypeService) roomTypeHandler {
	return roomTypeHandler{roomTypeSrv: roomTypeSrv}
}

func (h roomTypeHandler) FetchRoomTypes(c *fiber.Ctx) error {
	result, err := h.roomTypeSrv.FetchRoomTypes()

	if err != nil {
		return handleError(c, err)
	}

	return handlerOk(c, result)
}

func (h roomTypeHandler) GetRoomType(c *fiber.Ctx) error {
	id := c.Params("id")

	room, err := h.roomTypeSrv.GetRoomType(id)

	if err != nil {
		return handleError(c, err)
	}

	return handlerOk(c, *room)
}

func (h roomTypeHandler) CreateRoomType(c *fiber.Ctx) error {
	newRoomType := service.RoomTypeRequest{}

	err := c.BodyParser(&newRoomType)

	if err != nil {
		return handleError(c, fiber.NewError(fiber.StatusBadRequest, "Requested body incorrect form"))
	}

	result, err := h.roomTypeSrv.CreateRoomType(&newRoomType)

	if err != nil {
		return handleError(c, err)
	}

	return handlerCreatedOk(c, *result)
}

func (h roomTypeHandler) UpdateRoomType(c *fiber.Ctx) error {
	updateRoomType := service.RoomTypeRequest{}

	id := c.Params("id")

	err := c.BodyParser(&updateRoomType)

	if err != nil {
		return handleError(c, fiber.NewError(fiber.StatusBadRequest, "Requested body incorrect form"))
	}

	err = h.roomTypeSrv.UpdateRoomType(id, &updateRoomType)

	if err != nil {
		return handleError(c, err)
	}

	return handlerUpdateOk(c)
}

func (h roomTypeHandler) DeleteRoomType(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.roomTypeSrv.DeleteRoomType(id)

	if err != nil {
		return handleError(c, err)
	}

	return handlerDeleteOk(c)
}
