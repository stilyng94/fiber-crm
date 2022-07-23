package v1

import "github.com/gofiber/fiber/v2"

func InitRouter(app *fiber.App) {
	router := app.Group("/api/v1")
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
}
