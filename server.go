package main

import (
	v1 "stilyng94/fiber-crm/routers/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func createServer() {
	app := fiber.New(fiber.Config{AppName: "Fiber CRM", EnablePrintRoutes: true})

	port := ":" + viper.GetString("PORT")

	// api/v1
	v1.InitRouter(app)

	app.Listen(port)
}
