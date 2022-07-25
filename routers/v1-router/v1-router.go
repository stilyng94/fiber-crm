package v1router

import (
	"stilyng94/fiber-crm/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	router := app.Group("/api/v1")
	LeadRouter(router)
	// /api/v1
	router.Get("/", handlers.HealthCheck)
}
