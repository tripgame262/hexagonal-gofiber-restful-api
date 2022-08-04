package routes

import (
	"mansion/handler"
	"mansion/repository"
	"mansion/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func roomRoute(api fiber.Router, db *gorm.DB) {

	room := api.Group("/room")

	roomRepo := repository.NewRoomRepository(db)
	roomService := service.NewRoomService(roomRepo)
	roomHandler := handler.NewRoomHandler(roomService)

	room.Get("/:id", roomHandler.GetRoom)
	room.Post("/", roomHandler.CreateRoom)
	room.Get("/", roomHandler.FetchRooms)
	room.Put("/:id", roomHandler.UpdateRoom)
	room.Delete("/:id", roomHandler.DeleteRoom)
}
