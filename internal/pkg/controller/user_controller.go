package controller

import (
	"pbi/internal/pkg/usecase"
	"pbi/internal/pkg/models"
	"pbi/internal/helper"	

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	UpdateProfile(ctx *fiber.Ctx) error
	GetProfile(ctx *fiber.Ctx) error
}

type userControllerImpl struct {
	UsrUsc usecase.UserUseCase
}

func NewUserController(usrUsc usecase.UserUseCase) UserController {
	return &userControllerImpl{
		UsrUsc: usrUsc,
	}
}

// UpdateProfile godoc
// @Summary     Update user profile
// @Description Update authenticated user profile
// @Tags        User
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       body body object true "Update profile payload"
// @Success     200 {object} object "Success update profile"
// @Failure     400 {object} object "Bad Request"
// @Failure     401 {object} object "Unauthorized"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /user/profile [put]
func (uc *userControllerImpl) UpdateProfile(ctx *fiber.Ctx) error {
	req := new(models.UpdateProfileRequest)
	if err := ctx.BodyParser(req); err != nil {
		return helper.BadRequest(
			ctx,
			"Failed to UPDATE data",
			"Invalid request body",
		)
	}

	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.BadRequest(
			ctx,
			"Failed to UPDATE data",
			"User ID not found",
		)
	}

	res, herr := uc.UsrUsc.UpdateProfile(
		ctx.Context(),
		int64(userID),
		req,
	)
	if herr != nil {
		return helper.Error(
			ctx,
			herr.Code,        
			"Failed to UPDATE data",
			herr.Err.Error(), 
		)
	}

	return helper.Success(
		ctx,
		"Succeed to UPDATE data",
		res,
	)
}

// GetProfile godoc
// @Summary     Get user profile
// @Description Get authenticated user profile
// @Tags        User
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} object "Success get profile"
// @Failure     400 {object} object "Bad Request"
// @Failure     401 {object} object "Unauthorized"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /user/profile [get]
func (uc *userControllerImpl) GetProfile(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.BadRequest(
			ctx,
			"Failed to GET data",
			"User ID not found",
		)
	}

	res, herr := uc.UsrUsc.GetProfile(
		ctx.Context(),
		int64(userID),
	)
	if herr != nil {
		return helper.Error(
			ctx,
			herr.Code,
			"Failed to GET data",
			herr.Err.Error(),
		)
	}

	return helper.Success(
		ctx,
		"Succeed to GET data",
		res,
	)
}
