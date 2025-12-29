package controller

import (
	"pbi/internal/helper"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TransactionController interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
}

type transactionImpl struct {
	trxUsc usecase.TransactionUsecase
}

func NewTransactionController(trxUsc usecase.TransactionUsecase) TransactionController {
	return &transactionImpl{
		trxUsc: trxUsc,
	}
}

// Create godoc
// @Summary      Create Transaction
// @Description  Create a new transaction with multiple products
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body models.CreateTrxRequest true "Transaction Request"
// @Success      200 {object} helper.Response{data=int} "Transaction created successfully"
// @Failure      400 {object} helper.Response "Bad Request"
// @Failure      401 {object} helper.Response "Unauthorized"
// @Failure      404 {object} helper.Response "Alamat pengiriman tidak valid"
// @Failure      500 {object} helper.Response "Internal Server Error"
// @Security     BearerAuth
// @Router       /transactions [post]
func (c *transactionImpl) Create(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.Unauthorized(ctx, "Unauthorized")
	}

	var req models.CreateTrxRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest(ctx, "Failed to POST data", err.Error())
	}

	trxID, herr := c.trxUsc.CreateTransaction(ctx.Context(), userID, &req)
	if herr != nil {
		return helper.Error(ctx, herr.Code, "Failed to POST data", herr.Err.Error())
	}

	return helper.Success(ctx, "Succeed to POST data", trxID)
}

// GetAll godoc
// @Summary      Get All Transactions
// @Description  Get all transactions for authenticated user with full details
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Success      200 {object} helper.Response{data=models.TransactionListResponseWrapper} "Success get all transactions"
// @Failure      401 {object} helper.Response "Unauthorized"
// @Failure      500 {object} helper.Response "Internal Server Error"
// @Security     BearerAuth
// @Router       /transactions [get]
func (c *transactionImpl) GetAll(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.Unauthorized(ctx, "Unauthorized")
	}

	data, herr := c.trxUsc.GetAll(ctx.Context(), userID)
	if herr != nil {
		return helper.Error(ctx, herr.Code, "Failed to GET data", herr.Err.Error())
	}

	return helper.Success(ctx, "Succeed to GET data", data)
}

// GetByID godoc
// @Summary      Get Transaction by ID
// @Description  Get detailed transaction information by transaction ID
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        id path int true "Transaction ID"
// @Success      200 {object} helper.Response{data=models.TransactionDetailByIDResponse} "Success get transaction detail"
// @Failure      400 {object} helper.Response "Invalid transaction ID"
// @Failure      401 {object} helper.Response "Unauthorized"
// @Failure      404 {object} helper.Response "Transaction not found"
// @Failure      500 {object} helper.Response "Internal Server Error"
// @Security     BearerAuth
// @Router       /transactions/{id} [get]
func (c *transactionImpl) GetByID(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.Unauthorized(ctx, "Unauthorized")
	}

	trxID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequest(ctx, "Failed to GET data", "Invalid transaction ID")
	}

	data, herr := c.trxUsc.GetByID(ctx.Context(), trxID, userID)
	if herr != nil {
		return helper.Error(ctx, herr.Code, "Failed to GET data", herr.Err.Error())
	}

	return helper.Success(ctx, "Succeed to GET transaction detail", data)
}