package controller

import (
	"pbi/internal/pkg/usecase"
	"pbi/internal/pkg/models"
	"pbi/internal/helper"	

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	UpdateProfile(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	UsrUsc usecase.UserUseCase
}

func NewUserController(usrUsc usecase.UserUseCase) *UserControllerImpl {
	return &UserControllerImpl{
		UsrUsc: usrUsc,
	}
}

func (uc *UserControllerImpl) UpdateProfile(ctx *fiber.Ctx) error {
	req := new(models.UpdateProfileRequest)
	if err := ctx.BodyParser(req); err != nil {
		return helper.BadRequest(
			ctx,
			"Failed to POST data",
			"Invalid request body",
		)
	}

	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.BadRequest(
			ctx,
			"Failed to POST data",
			"User ID not found",
		)
	}

	res, err := uc.UsrUsc.UpdateProfile(ctx.Context(), int64(userID), req)
	if err != nil {
		return helper.BadRequest(
			ctx,
			"Failed to POST data",
			err.Error(),
		)
	}

	return helper.Success(
		ctx,
		"Succeed to POST data",
		res,
	)
}
