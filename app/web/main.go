package main

import (
	"fmt"
	"pbi/docs" 
	"pbi/internal/config/container"
	"pbi/internal/server/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title           PBI API
// @version         1.0
// @description     PBI API Documentation
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@example.com


// @host      localhost:3000
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	cont := container.InitContainer()

	docs.SwaggerInfo.Title = "PBI API"
	docs.SwaggerInfo.Description = "PBI API Documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", cont.Config.App.Port)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app := fiber.New()

	http.HttpRouteInit(app, cont)

	app.Get("/swagger/*", swagger.HandlerDefault)

	err := app.Listen(fmt.Sprintf(":%d", cont.Config.App.Port))
	if err != nil {
		fmt.Println("Failed to start server yeah:s", err)
	}

}
