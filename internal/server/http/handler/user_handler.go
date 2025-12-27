package handler

import (
	"pbi/internal/config/middleware"
	usercontroller "pbi/internal/pkg/controller"
	"pbi/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router, userUsc usecase.UserUseCase) {
	controller := usercontroller.NewUserController(userUsc)

	rest := r.Group("/user")
	rest.Get("/profile", middleware.AuthChecker(true), controller.GetProfile)
	rest.Put("/profile", middleware.AuthChecker(true), controller.UpdateProfile)
}