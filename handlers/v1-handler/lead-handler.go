package v1handler

import (
	"log"
	"stilyng94/fiber-crm/database"
	"stilyng94/fiber-crm/ent"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetLead(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	leadId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)

	}

	lead, err := database.Db.Lead.Get(ctx.Context(), leadId)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)

	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"lead": lead,
	})
}

func GetLeads(ctx *fiber.Ctx) error {
	leads, err := database.Db.Lead.Query().All(ctx.Context())
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"leads": leads,
	})
}

func UpdateLead(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	leadId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)

	}
	var lead ent.Lead

	if err := ctx.BodyParser(&lead); err != nil {
		log.Println(err)
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}

	updatedLead, err := database.Db.Lead.UpdateOneID(leadId).SetCompany(lead.Company).SetName(lead.Name).Save(ctx.Context())
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)

	}
	return ctx.Status(fiber.StatusOK).JSON(updatedLead)

}
func DeleteLead(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	leadId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)

	}
	deleteQuery := database.Db.Lead.DeleteOneID(leadId)

	if err := deleteQuery.Exec(ctx.Context()); err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)

	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
func CreateLead(ctx *fiber.Ctx) error {
	var lead ent.Lead

	if err := ctx.BodyParser(&lead); err != nil {
		log.Println(err)
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}

	newLead, err := database.Db.Lead.Create().SetCompany(lead.Company).SetEmail(lead.Company).SetName(lead.Name).SetPhone(lead.Phone).Save(ctx.Context())

	if err != nil {
		log.Println(err)
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}
	return ctx.Status(fiber.StatusCreated).JSON(newLead)
}
