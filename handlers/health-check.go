package handlers

import "github.com/gofiber/fiber/v2"

// @title       HealthCheck
// @Summary     Show the status of server.
// @Description get the status of server.
// @Tags        root
// @Accept      */*
// @Produce     json
// @Success     200 {string} string
// @Router      /api/ [get]
func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
