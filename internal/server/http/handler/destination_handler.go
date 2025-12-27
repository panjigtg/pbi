package handler

import (
	"pbi/internal/config/middleware"
	"pbi/internal/pkg/controller"
	"pbi/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)


func DestinationRoute(r fiber.Router, DestUsc usecase.DestinationUsecase){
	DestController := controller.NewDestinationController(DestUsc)

	rest := r.Group("user")
	rest.Get("/alamat", middleware.AuthChecker(true), DestController.GetAll)
	rest.Get("/alamat/:id", middleware.AuthChecker(true), DestController.GetByID)
	rest.Post("/alamat",middleware.AuthChecker(true), DestController.Create)
	rest.Put("/alamat/:id",middleware.AuthChecker(true), DestController.Update)
	rest.Delete("/alamat/:id",middleware.AuthChecker(true), DestController.Delete)
}