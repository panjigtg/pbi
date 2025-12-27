package usecase

import (
	"context"
	"errors"
	"pbi/internal/helper"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/repository"
)

type UserUseCase interface {
	UpdateProfile(ctx context.Context, userID int64, req *models.UpdateProfileRequest) (*models.UserResponse, *helper.ErrorStruct)
	GetProfile(ctx context.Context, userID int64) (*models.UserResponse, *helper.ErrorStruct)
}

type userUsecaseImpl struct {
	userRepo repository.UserRepository
	addrUsc  AddressUsecase

}

func NewUserUseCase(userRepo repository.UserRepository, addrUsc AddressUsecase) UserUseCase {
	return &userUsecaseImpl{
		userRepo: userRepo,
		addrUsc: addrUsc,
	}
}

func (u *userUsecaseImpl) UpdateProfile(ctx context.Context,userID int64,req *models.UpdateProfileRequest,) (*models.UserResponse, *helper.ErrorStruct) {
	if req.IDProvinsi != nil && req.IDKota != nil {
		if herr := u.addrUsc.ValidateRegion(ctx, *req.IDProvinsi, *req.IDKota); herr != nil {
			return nil, herr
		}
	}

	if err := u.userRepo.UpdateProfile(
		ctx,
		userID,
		req.Nama,
		req.NoTelp,
		req.Email,
		req.IDProvinsi,
		req.IDKota,
		req.Tentang,
		req.Pekerjaan,
		req.KataSandi,
		req.TanggalLahir,
		req.JenisKelamin,
	); err != nil {
		return nil, &helper.ErrorStruct{Code: 400, Err: err}
	}

	return u.GetProfile(ctx, userID)
}

func (u *userUsecaseImpl) GetProfile(ctx context.Context, userID int64) (*models.UserResponse, *helper.ErrorStruct) {

	user, err := u.userRepo.FindById(ctx, userID)
	if err != nil {
		return nil, &helper.ErrorStruct{Code: 500, Err: err}
	}
	if user == nil {
		return nil, &helper.ErrorStruct{Code: 404, Err: errors.New("user tidak ditemukan")}
	}

	var prov models.ProvinceResponse
	var city models.CityResponse

	if user.IDProvinsi != "" {
		provStruct, herr := u.addrUsc.GetProvinceDetail(ctx, user.IDProvinsi)
		if herr != nil {
			return nil, herr
		}
		prov = models.ProvinceResponse{
			ID:   provStruct.ID,
			Name: provStruct.Name,
		}
	}

	if user.IDKota != "" && user.IDProvinsi != "" {
		cityStruct, herr := u.addrUsc.GetCityDetail(ctx, user.IDProvinsi, user.IDKota)
		if herr != nil {
			return nil, herr
		}
		city = models.CityResponse{
			ID:         cityStruct.ID,
			ProvinceID: cityStruct.ProvinceID,
			Name:       cityStruct.Name,
		}
	}

	resp := &models.UserResponse{
		Nama:         user.Nama,
		Email:        user.Email,
		NoTelp:       user.NoTelp,
		TanggalLahir: user.TanggalLahir,
		JenisKelamin: user.JenisKelamin,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		IDProvinsi:   prov,
		IDKota:       city,
	}

	return resp, nil
}
