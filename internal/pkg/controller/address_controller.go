package controller

import (
	"pbi/internal/helper"
	"pbi/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type AddressController interface {
	GetProvinces(ctx *fiber.Ctx) error
	GetCity(ctx *fiber.Ctx) error
	GetProvinceDetail(ctx *fiber.Ctx) error
	GetCityDetail(ctx *fiber.Ctx) error
}

type addressControllerImpl struct {
	uc usecase.AddressUsecase
}

func NewAddressController(uc usecase.AddressUsecase) AddressController {
	return &addressControllerImpl{
		uc: uc,
	}
}


func (ac *addressControllerImpl) GetProvinces(c *fiber.Ctx) error {
	res, err := ac.uc.GetProvinces(c.Context())
	if err != nil {
		return helper.BadRequest(c, "Failed get provinces", err.Err.Error())
	}
	return helper.Success(c, "OK", res)
}

func (ac *addressControllerImpl) GetCity(c *fiber.Ctx) error {
	provinceID := c.Params("province_id")

	res, err := ac.uc.GetCities(c.Context(), provinceID)
	if err != nil {
		return helper.BadRequest(c, "Failed get cities", err.Err.Error())
	}

	return helper.Success(c, "OK", res)
}

func (ac *addressControllerImpl) GetProvinceDetail(c *fiber.Ctx) error {
	provinceID := c.Params("province_id")

	res, err := ac.uc.GetProvinceDetail(c.Context(), provinceID)
	if err != nil {
		return helper.BadRequest(c, "Failed get province detail", err.Err.Error())
	}

	return helper.Success(c, "OK", res)
}

func (ac *addressControllerImpl) GetCityDetail(c *fiber.Ctx) error {
	provinceID := c.Params("province_id")
	cityID := c.Params("city_id")

	res, err := ac.uc.GetCityDetail(c.Context(), provinceID, cityID)
	if err != nil {
		return helper.BadRequest(c, "Failed get city detail", err.Err.Error())
	}

	return helper.Success(c, "OK", res)
}
