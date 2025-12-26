package handler

import (
	"pbi/internal/config/middleware"
	"pbi/internal/pkg/controller"
	"pbi/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func TokoRoute(r fiber.Router, tokoUsc usecase.TokoUsecase) {
	tokoController := controller.NewTokoController(tokoUsc)

	rest := r.Group("/toko")
	rest.Get("/", tokoController.GetAll)
	rest.Get("/my", middleware.AuthChecker(true), tokoController.GetMy)
	rest.Get("/:id", tokoController.GetByID)
	rest.Put("/:id", middleware.AuthChecker(true), tokoController.Update)
}