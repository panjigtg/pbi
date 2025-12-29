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

// GetProvinces godoc
// @Summary      Get All Provinces
// @Description  Get list of all provinces in Indonesia
// @Tags         Address
// @Accept       json
// @Produce      json
// @Success      200 {object} object "Success get provinces"
// @Failure      400 {object} object "Bad Request"
// @Failure      500 {object} object "Internal Server Error"
// @Router       /provcity/listprovincies [get]
func (ac *addressControllerImpl) GetProvinces(c *fiber.Ctx) error {
	res, err := ac.uc.GetProvinces(c.Context())
	if err != nil {
		return helper.BadRequest(c, "Failed get provinces", err.Err.Error())
	}
	return helper.Success(c, "OK", res)
}

// GetCity godoc
// @Summary      Get Cities by Province
// @Description  Get list of cities in a specific province
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        province_id path string true "Province ID"
// @Success      200 {object} object "Success get cities"
// @Failure      400 {object} object "Bad Request"
// @Failure      500 {object} object "Internal Server Error"
// @Router       /provcity/listcities/{province_id} [get]
func (ac *addressControllerImpl) GetCity(c *fiber.Ctx) error {
	provinceID := c.Params("province_id")

	res, err := ac.uc.GetCities(c.Context(), provinceID)
	if err != nil {
		return helper.BadRequest(c, "Failed get cities", err.Err.Error())
	}

	return helper.Success(c, "OK", res)
}

// GetProvinceDetail godoc
// @Summary      Get Province Detail
// @Description  Get detailed information of a specific province
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        province_id path string true "Province ID"
// @Success      200 {object} object "Success get province detail"
// @Failure      400 {object} object "Bad Request"
// @Failure      404 {object} object "Province not found"
// @Failure      500 {object} object "Internal Server Error"
// @Router       /provcity/detailprovince/{province_id} [get]
func (ac *addressControllerImpl) GetProvinceDetail(c *fiber.Ctx) error {
	provinceID := c.Params("province_id")

	res, err := ac.uc.GetProvinceDetail(c.Context(), provinceID)
	if err != nil {
		return helper.BadRequest(c, "Failed get province detail", err.Err.Error())
	}

	return helper.Success(c, "OK", res)
}

// GetCityDetail godoc
// @Summary      Get City Detail
// @Description  Get detailed information of a specific city
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        province_id path string true "Province ID"
// @Param        city_id path string true "City ID"
// @Success      200 {object} object "Success get city detail"
// @Failure      400 {object} object "Bad Request"
// @Failure      404 {object} object "City not found"
// @Failure      500 {object} object "Internal Server Error"
// @Router       /provcity/detailcity/{province_id}/cities/{city_id} [get]
func (ac *addressControllerImpl) GetCityDetail(c *fiber.Ctx) error {
	provinceID := c.Params("province_id")
	cityID := c.Params("city_id")

	res, err := ac.uc.GetCityDetail(c.Context(), provinceID, cityID)
	if err != nil {
		return helper.BadRequest(c, "Failed get city detail", err.Err.Error())
	}

	return helper.Success(c, "OK", res)
}
