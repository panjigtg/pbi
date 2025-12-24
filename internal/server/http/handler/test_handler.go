package handler

import (
	authcontroller "pbi/internal/pkg/controller"
	"pbi/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r fiber.Router, authUsc usecase.AuthUseCase) {
	controller := authcontroller.NewAuthController(authUsc)

	r.Post("/register", controller.Register)
	// r.Post("/login", controller.Login)
}
