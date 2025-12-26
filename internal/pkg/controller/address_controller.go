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

type AddressControllerImpl struct {
	uc usecase.AddressUsecase
}

func NewAddressController(uc usecase.AddressUsecase) AddressController {
	return &AddressControllerImpl{
		uc: uc,
	}
}

func (ac *AddressControllerImpl) GetProvinces(c *fiber.Ctx) error {
	res, err := ac.uc.GetProvinces(c.Context())
	if err != nil {
		return helper.BadRequest(c, "Failed get provinces", err.Err.Error())
	}
	return helper.Success(c, "OK", res)
}

func (ac *AddressControllerImpl) GetCity(c *fiber.Ctx) error {
	provinceID := c.Params("province_id")

	res, err := ac.uc.GetCities(c.Context(), provinceID)
	if err != nil {
		return helper.BadRequest(c, "Failed get cities", err.Err.Error())
	}

	return helper.Success(c, "OK", res)
}

func (ac *AddressControllerImpl) GetProvinceDetail(c *fiber.Ctx) error {
	provinceID := c.Params("province_id")

	res, err := ac.uc.GetProvinceDetail(c.Context(), provinceID)
	if err != nil {
		return helper.BadRequest(c, "Failed get province detail", err.Err.Error())
	}

	return helper.Success(c, "OK", res)
}

func (ac *AddressControllerImpl) GetCityDetail(c *fiber.Ctx) error {
	provinceID := c.Params("province_id")
	cityID := c.Params("city_id")

	res, err := ac.uc.GetCityDetail(c.Context(), provinceID, cityID)
	if err != nil {
		return helper.BadRequest(c, "Failed get city detail", err.Err.Error())
	}

	return helper.Success(c, "OK", res)
}
