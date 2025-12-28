package handler

import (
	"pbi/internal/config/middleware"
	"pbi/internal/pkg/controller"
	"pbi/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func ProductRoute(ctx fiber.Router, Pusc usecase.ProductUsecase ){
	controller := controller.NewProductController(Pusc)

	rest := ctx.Group("/product")
	rest.Get("/",controller.GetAll)
	rest.Get("/:id",controller.GetByID)
	rest.Post("/",middleware.AuthChecker(true),controller.Create)
	rest.Put("/:id",middleware.AuthChecker(true),controller.Update)
	rest.Delete("/:id",middleware.AuthChecker(true),controller.Delete)
}