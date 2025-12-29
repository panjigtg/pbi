package controller

import (
	"github.com/gofiber/fiber/v2"

	"pbi/internal/helper"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/usecase"
)

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}


type authControllerImpl struct {
	authUsc usecase.AuthUseCase
}

func NewAuthController(authUsc usecase.AuthUseCase) AuthController {
	return &authControllerImpl{
		authUsc: authUsc,
	}
}

// Register godoc
// @Summary     Register user
// @Description Create new user account
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param 		body body models.RegisterRequest true "Register payload"
// @Success     200 {object} object "Success register user"
// @Failure     400 {object} object "Bad Request"
// @Failure     409 {object} object "Email already exists"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /auth/register [post]
func (uc *authControllerImpl) Register(ctx *fiber.Ctx) error {
	req := new(models.RegisterRequest)

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

// Login godoc
// @Summary     Login user
// @Description Login using email and password
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Success     200 {object} object "Success login"
// @Failure     400 {object} object "Bad Request"
// @Failure     401 {object} object "Invalid email or password"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /auth/login [post]
func (uc *authControllerImpl) Login(ctx *fiber.Ctx) error {
	req := new(models.LoginRequest)

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

