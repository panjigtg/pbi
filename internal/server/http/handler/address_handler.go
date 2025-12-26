package handler

import (
	"pbi/internal/pkg/controller"
	"pbi/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func AddressRoute(r fiber.Router, AddrUsc usecase.AddressUsecase){
	addressController := controller.NewAddressController(AddrUsc)

	rest := r.Group("/provcity")
	rest.Get("/listprovincies", addressController.GetProvinces)
	rest.Get("/detailprovince/:province_id", addressController.GetProvinceDetail)
	rest.Get("/listcities/:province_id", addressController.GetCity)
	rest.Get("/detailcity/:province_id/cities/:city_id", addressController.GetCityDetail)
}