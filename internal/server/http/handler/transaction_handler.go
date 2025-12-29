package handler

import (
	"pbi/internal/config/middleware"
	"pbi/internal/pkg/controller"
	"pbi/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoute(r fiber.Router, TrxUsc usecase.TransactionUsecase) {
	trxcontroller := controller.NewTransactionController(TrxUsc)

	rest :=r.Group("/trx")
	rest.Get("/",middleware.AuthChecker(true), trxcontroller.GetAll)
	rest.Get("/:id",middleware.AuthChecker(true), trxcontroller.GetByID)
	rest.Post("/",middleware.AuthChecker(true), trxcontroller.Create)
}