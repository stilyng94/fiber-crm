package v1

import (
	v1 "stilyng94/fiber-crm/handlers/v1"

	"github.com/gofiber/fiber/v2"
)

func LeadRouter(router fiber.Router) {
	leadRouter := router.Group("/lead")
	leadRouter.Get("/:id", v1.GetLead)
	leadRouter.Delete("/:id", v1.DeleteLead)
	leadRouter.Put("/:id", v1.UpdateLead)
	leadRouter.Get("/", v1.GetLeads)
	leadRouter.Post("/", v1.CreateLead)
}
