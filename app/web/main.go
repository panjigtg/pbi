package main

import (
	"fmt"

	"pbi/internal/config/container"
	"pbi/internal/server/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cont := container.InitContainer()

	app := fiber.New()

	http.HttpRouteInit(app, cont)

	err := app.Listen(fmt.Sprintf(":%d", cont.Config.App.Port))
	if err != nil {
		fmt.Println("Failed to start server yeah:s", err)
	}

}
