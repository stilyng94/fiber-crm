package main

import (
	"stilyng94/fiber-crm/docs"
	_ "stilyng94/fiber-crm/docs"
	"stilyng94/fiber-crm/handlers"
	v1router "stilyng94/fiber-crm/routers/v1-router"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

// @title          Fiber Swagger Example API
// @version        2.0
// @description    This is a sample server.
// @termsOfService http://swagger.io/terms/
// @contact.name  stilyng94
// @contact.email oseipaulkofi@gmail.com

// @BasePath /
func createServer() {
	port := ":" + viper.GetString("PORT")
	docs.SwaggerInfo.Host = "localhost" + port
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app := fiber.New(fiber.Config{AppName: "Fiber CRM", EnablePrintRoutes: true})
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(cors.New(cors.ConfigDefault))
	app.Use(limiter.New(limiter.Config{Max: 15}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	// api/v1
	v1router.InitRouter(app)

	//api
	app.Get("/api", handlers.HealthCheck)

	app.Listen(port)
}
