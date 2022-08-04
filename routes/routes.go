package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func InitRoute(app *fiber.App, db *gorm.DB) {

	api := app.Group("/api", logger.New())

	roomTypeRoute(api, db)
	roomRoute(api, db)

}
