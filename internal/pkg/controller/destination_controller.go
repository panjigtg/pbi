package controller

import (
	"pbi/internal/helper"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DestinationController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
}

type destinationImpl struct {
	DesUsc usecase.DestinationUsecase
}

func NewDestinationController(Usc usecase.DestinationUsecase) DestinationController {
	return &destinationImpl{
		DesUsc: Usc,
	}
}

func(dc *destinationImpl) Create(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.Unauthorized(ctx, "invalid user")
	}

	var req models.DestinationCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest(ctx, "invalid request body")
	}

	if req.JudulAlamat == "" ||
		req.NamaPenerima == "" ||
		req.NoTelp == "" ||
		req.DetailAlamat == "" {
		return helper.BadRequest(ctx, "all fields are required")
	}

	res, herr := dc.DesUsc.Create(ctx.Context(), userID, &req)
	if herr != nil {
		// message ditentukan DI SINI
		message := "Failed to CREATE data"

		// error detail (optional, untuk debug client)
		errMsg := ""
		if herr.Err != nil {
			errMsg = herr.Err.Error()
		}

		return helper.Error(
			ctx,
			herr.Code,
			message,
			errMsg,
		)
	}

	return helper.Success(ctx, "Succeed to CREATE data", res)
}

func (dc *destinationImpl) Update(ctx *fiber.Ctx) error {
	// ambil user dari JWT
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.Unauthorized(ctx, "invalid user")
	}

	// ambil ID dari param
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return helper.BadRequest(ctx, "invalid destination id")
	}

	// parse body
	var req models.DestinationUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest(ctx, "invalid request body")
	}

	// validasi minimal (judul TIDAK ADA)
	if req.NamaPenerima == "" &&
		req.NoTelp == "" &&
		req.DetailAlamat == "" {
		return helper.BadRequest(ctx, "no fields to update")
	}

	res, herr := dc.DesUsc.Update(ctx.Context(), id, userID, &req)
	if herr != nil {
		return helper.Error(
			ctx,
			herr.Code,
			"Failed to UPDATE data",
			herr.Err.Error(),
		)
	}

	return helper.Success(ctx, "Succeed to UPDATE data", res)
}

func (dc *destinationImpl) Delete(ctx *fiber.Ctx) error {
	// ambil user dari JWT
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.Unauthorized(ctx, "invalid user")
	}

	// ambil ID dari param
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return helper.BadRequest(ctx, "invalid destination id")
	}

	herr := dc.DesUsc.Delete(ctx.Context(), id, userID)
	if herr != nil {
		return helper.Error(
			ctx,
			herr.Code,
			"Failed to DELETE data",
			herr.Err.Error(),
		)
	}

	return helper.Success(ctx, "Succeed to DELETE data", nil)
}

func (dc *destinationImpl) GetAll(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.Unauthorized(ctx, "invalid user")
	}

	res, herr := dc.DesUsc.GetAll(ctx.Context(), userID)
	if herr != nil {
		return helper.Error(ctx, herr.Code, "Failed to GET data", herr.Err.Error())
	}

	return helper.Success(ctx, "Succeed to GET data", res)
}

func (dc *destinationImpl) GetByID(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.Unauthorized(ctx, "invalid user")
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequest(ctx, "invalid destination id")
	}

	res, herr := dc.DesUsc.GetByID(ctx.Context(), id, userID)
	if herr != nil {
		return helper.Error(ctx, herr.Code, "Failed to GET data", herr.Err.Error())
	}

	return helper.Success(ctx, "Succeed to GET data", res)
}
