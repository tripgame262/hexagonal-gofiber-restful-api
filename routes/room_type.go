package routes

import (
	"mansion/handler"
	"mansion/repository"
	"mansion/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func roomTypeRoute(api fiber.Router, db *gorm.DB) {

	roomType := api.Group("/room-type")

	roomTypeRepo := repository.NewRoomTypeRepository(db)
	roomTypeService := service.NewRoomTypeService(roomTypeRepo)
	roomTypeHandler := handler.NewRoomTypeHandler(roomTypeService)

	roomType.Get("/", roomTypeHandler.FetchRoomTypes)
	roomType.Get("/:id", roomTypeHandler.GetRoomType)
	roomType.Post("/", roomTypeHandler.CreateRoomType)
	roomType.Put("/:id", roomTypeHandler.UpdateRoomType)
	roomType.Delete("/:id", roomTypeHandler.DeleteRoomType)
}
