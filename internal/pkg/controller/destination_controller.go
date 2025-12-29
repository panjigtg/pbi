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

// CreateDestination godoc
// @Summary     Create destination address
// @Description Create new destination address for logged-in user
// @Tags        Destination
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       body body models.DestinationCreateRequest true "Destination payload"
// @Success     200 {object} object "Success create destination"
// @Failure     400 {object} object "Bad Request"
// @Failure     401 {object} object "Unauthorized"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /user/alamat [post]
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

// UpdateDestination godoc
// @Summary     Update destination
// @Description Update destination address by ID (owned by user)
// @Tags        Destination
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path int true "Destination ID"
// @Param       body body models.DestinationUpdateRequest true "Update payload"
// @Success     200 {object} object "Success update destination"
// @Failure     400 {object} object "Bad Request"
// @Failure     401 {object} object "Unauthorized"
// @Failure     404 {object} object "Destination not found"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /user/alamat/{id} [put]
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

// DeleteDestination godoc
// @Summary     Delete destination
// @Description Delete destination address by ID (owned by user)
// @Tags        Destination
// @Produce     json
// @Security    BearerAuth
// @Param       id path int true "Destination ID"
// @Success     200 {object} object "Success delete destination"
// @Failure     400 {object} object "Invalid destination ID"
// @Failure     401 {object} object "Unauthorized"
// @Failure     404 {object} object "Destination not found"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /user/alamat/{id} [delete]
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

// GetAllDestinations godoc
// @Summary     Get all destinations
// @Description Get all destination addresses owned by logged-in user
// @Tags        Destination
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} object "Success get destination list"
// @Failure     401 {object} object "Unauthorized"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /user/alamat [get]
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

// GetDestinationByID godoc
// @Summary     Get destination by ID
// @Description Get destination address detail by ID (owned by user)
// @Tags        Destination
// @Produce     json
// @Security    BearerAuth
// @Param       id path int true "Destination ID"
// @Success     200 {object} object "Success get destination detail"
// @Failure     400 {object} object "Invalid destination ID"
// @Failure     401 {object} object "Unauthorized"
// @Failure     404 {object} object "Destination not found"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /user/alamat/{id} [get]
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
