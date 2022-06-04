package main

import (
	"flutterexplorer/endpoints"
	"flutterexplorer/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			c.Status(code).JSON(models.FailResponse(err.Error(), code))
			return nil
		},
	})

	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group("/api") // /api
	api.Get("/widget", endpoints.GetWidgets)
	api.Get("/package", endpoints.GetPackages)

	app.Listen(":25372")
}
