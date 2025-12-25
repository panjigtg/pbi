package handler

import (
	"pbi/internal/config/middleware"
	cc "pbi/internal/pkg/controller"
	"pbi/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoute(r fiber.Router, cUsc usecase.CategoryUseCase) {
	categoryController := cc.NewCategoryController(cUsc)

	rest := r.Group("/category")

	rest.Use(middleware.AuthChecker(true))
	rest.Get("/",middleware.AdminChecker(true), categoryController.GetAllCategories)
	rest.Get("/:id",middleware.AdminChecker(true), categoryController.GetById)
	rest.Post("/",middleware.AdminChecker(true), categoryController.CreateCategory)
	rest.Put("/:id",middleware.AdminChecker(true), categoryController.Update)
	rest.Delete("/:id",middleware.AdminChecker(true), categoryController.Delete)
}