package main

import (
	"encoding/json"
	"fmt"

	"pbi/internal/config/container"
	"pbi/internal/server/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cont := container.InitContainer()
	fmt.Printf("Container initialized: %+v\n", cont)
	fmt.Printf("UserUsc is nil: %v\n", cont.UserUsc == nil)

	app := fiber.New()

	http.HttpRouteInit(app, cont)

	// Di main.go
	for _, route := range app.GetRoutes() {
		fmt.Printf("Method: %s, Path: %s\n", route.Method, route.Path)
	}

	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Println(string(data))

	err := app.Listen(fmt.Sprintf(":%d", cont.Config.App.Port))
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}

}
