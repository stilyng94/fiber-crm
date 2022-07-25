package v1router

import (
	v1handler "stilyng94/fiber-crm/handlers/v1-handler"

	"github.com/gofiber/fiber/v2"
)

func LeadRouter(router fiber.Router) {
	leadRouter := router.Group("/lead")
	leadRouter.Get("/:id", v1handler.GetLead)
	leadRouter.Delete("/:id", v1handler.DeleteLead)
	leadRouter.Put("/:id", v1handler.UpdateLead)
	leadRouter.Get("/", v1handler.GetLeads)
	leadRouter.Post("/", v1handler.CreateLead)
}
