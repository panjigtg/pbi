package controller

import (
	"github.com/gofiber/fiber/v2"

	"pbi/internal/helper"
	authmodels "pbi/internal/pkg/models"
	authusc "pbi/internal/pkg/usecase"
)

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type authControllerImpl struct {
	authUsc authusc.AuthUseCase
}

func NewAuthController(authUsc authusc.AuthUseCase) AuthController {
	return &authControllerImpl{
		authUsc: authUsc,
	}
}

func (uc *authControllerImpl) Register(ctx *fiber.Ctx) error {
	req := new(authmodels.RegisterRequest)

	if err := ctx.BodyParser(req); err != nil {
		return helper.BadRequest(
			ctx,
			"Failed to POST data",
			"Invalid request body",
		)
	}

	res, err := uc.authUsc.Register(ctx.Context(), req)
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

func (uc *authControllerImpl) Login(ctx *fiber.Ctx) error {
	req := new(authmodels.LoginRequest)

	if err := ctx.BodyParser(req); err != nil {
		return helper.BadRequest(
			ctx,
			"Failed to POST data",
			"Invalid request body",
		)
	}

	res, err := uc.authUsc.Login(ctx.Context(), req)
	if err != nil {
		return helper.BadRequest(
			ctx,
			"Failed to POST data",
			"Email atau password salah",
		)
	}

	return helper.Success(
		ctx,
		"Succeed to POST data",
		res,
	)
}

