package handler

import (
	authcontroller "pbi/internal/pkg/controller"
	"pbi/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r fiber.Router, authUsc usecase.AuthUseCase) {
	controller := authcontroller.NewAuthController(authUsc)
	rest := r.Group("/auth")

	rest.Post("/register", controller.Register)
	rest.Post("/login", controller.Login)
}
