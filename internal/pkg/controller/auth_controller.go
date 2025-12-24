package controller

import (
	"github.com/gofiber/fiber/v2"

	"pbi/internal/helper"
	authmodels "pbi/internal/pkg/models"
	authusecase "pbi/internal/pkg/usecase"
)

type AuthController interface {
	// Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type AuthControllerImpl struct {
	authUsc authusecase.AuthUsecase
}

func NewAuthController(authUsc authusecase.AuthUsecase) AuthController {
	return &AuthControllerImpl{
		authUsc: authUsc,
	}
}

func (uc *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	req := new(authmodels.RegisterRequest)

	// Parse body request
	if err := ctx.BodyParser(req); err != nil {
		return helper.BadRequest(ctx, "invalid request", err.Error())
	}

	// Panggil usecase register
	res, err := uc.authUsc.Register(ctx.Context(), req)
	if err != nil {
		return helper.BadRequest(ctx, err.Error(), nil)
	}

	return helper.Created(ctx, "user berhasil dibuat", res)
}

func (uc *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement me
}