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
