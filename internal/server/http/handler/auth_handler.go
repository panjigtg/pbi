package handler

import (
	"fmt"
	authcontroller "pbi/internal/pkg/controller"
	"pbi/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r fiber.Router, authUsc usecase.AuthUseCase) {
	fmt.Println("=== Registering Auth Routes ===")
	controller := authcontroller.NewAuthController(authUsc)

	r.Post("/register", controller.Register)
	fmt.Println("Route registered: POST /api/v1/register")

	// r.Post("/login", controller.Login)
}
