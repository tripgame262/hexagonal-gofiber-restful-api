package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, err error) error {

	if e, ok := err.(*fiber.Error); ok {
		return c.Status(e.Code).JSON(fiber.Map{
			"status":  false,
			"message": e.Message,
		})
	}

	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"status":  false,
		"message": "Unexpected error",
	})
}

func handlerOk(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Get result success",
		"data":    data,
	})
}

func handlerCreatedOk(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Add data successfully",
		"data":    data,
	})
}

func handlerUpdateOk(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Update data successfully",
	})
}

func handlerDeleteOk(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Delete data successfully",
	})
}
