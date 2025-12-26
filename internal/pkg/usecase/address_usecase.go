package usecase

import (
	"context"
	"errors"
	"pbi/internal/helper"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/repository"
)

type AddressUsecase interface {
	GetProvinces(ctx context.Context) ([]*models.Province, *helper.ErrorStruct)
	GetCities(ctx context.Context, provinceID string) ([]*models.City, *helper.ErrorStruct)
	ValidateRegion(ctx context.Context, provinceID, cityID string) *helper.ErrorStruct
	GetProvinceDetail(ctx context.Context, provinceID string) (*models.Province, *helper.ErrorStruct)
	GetCityDetail(ctx context.Context, provinceID, cityID string) (*models.City, *helper.ErrorStruct)
}

type addressUsecaseImpl struct {
	Addrrepo repository.AddressRepository
}

func NewAddressUsecase(repo repository.AddressRepository) AddressUsecase {
	return &addressUsecaseImpl{
		Addrrepo: repo,
	}
}

func (u *addressUsecaseImpl) GetProvinces(ctx context.Context) ([]*models.Province, *helper.ErrorStruct) {
	res, err := u.Addrrepo.GetProvinces()
	if err != nil {
		return nil, &helper.ErrorStruct{
			Err:  err,
			Code: 500,
		}
	}
	return res, nil
}

func (u *addressUsecaseImpl) GetCities(ctx context.Context, provinceID string) ([]*models.City, *helper.ErrorStruct) {
	if provinceID == "" {
		return nil, &helper.ErrorStruct{
			Err:  errors.New("province_id wajib diisi"),
			Code: 400,
		}
	}

	res, err := u.Addrrepo.GetCities(provinceID)
	if err != nil {
		return nil, &helper.ErrorStruct{
			Err:  err,
			Code: 500,
		}
	}

	return res, nil
}

func (u *addressUsecaseImpl) ValidateRegion(ctx context.Context, provinceID, cityID string,) *helper.ErrorStruct {

	if provinceID == "" || cityID == "" {
		return &helper.ErrorStruct{
			Err:  errors.New("province_id dan city_id wajib diisi"),
			Code: 400,
		}
	}

	cities, err := u.Addrrepo.GetCities(provinceID)
	if err != nil {
		return &helper.ErrorStruct{
			Err:  err,
			Code: 500,
		}
	}

	for _, c := range cities {
		if c.ID == cityID {
			return nil
		}
	}

	return &helper.ErrorStruct{
		Err:  errors.New("city tidak valid untuk province tersebut"),
		Code: 400,
	}
}

func (u *addressUsecaseImpl) GetProvinceDetail(
	ctx context.Context,
	provinceID string,
) (*models.Province, *helper.ErrorStruct) {

	if provinceID == "" {
		return nil, &helper.ErrorStruct{
			Err:  errors.New("province_id wajib diisi"),
			Code: 400,
		}
	}

	res, err := u.Addrrepo.GetProvinceByID(provinceID)
	if err != nil {
		return nil, &helper.ErrorStruct{
			Err:  err,
			Code: 404,
		}
	}

	return res, nil
}

func (u *addressUsecaseImpl) GetCityDetail(
	ctx context.Context,
	provinceID, cityID string,
) (*models.City, *helper.ErrorStruct) {

	if provinceID == "" || cityID == "" {
		return nil, &helper.ErrorStruct{
			Err:  errors.New("province_id dan city_id wajib diisi"),
			Code: 400,
		}
	}

	res, err := u.Addrrepo.GetCityByID(provinceID, cityID)
	if err != nil {
		return nil, &helper.ErrorStruct{
			Err:  err,
			Code: 404,
		}
	}

	return res, nil
}
