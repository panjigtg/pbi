package main

import (
	"fmt"

	"pbi/internal/config/container"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cont := container.InitContainer()

	app := fiber.New()

	app.Listen(fmt.Sprintf(":%d", cont.Config.App.Port))
}
