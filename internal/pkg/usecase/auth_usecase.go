package usecase

import (
	"context"
	"errors"

	// "fmt"
	"pbi/internal/pkg/entity"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/repository"
	"pbi/internal/utils"
)

type AuthUseCase interface {
	Register(ctx context.Context, req *models.RegisterRequest) (string, error)
	Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error)
}

type authUsecaseImpl struct {
	userRepo repository.UserRepository
	tokoRepo repository.TokoRepository
	AddrUsc  AddressUsecase
}

func NewAuthUseCase(userRepo repository.UserRepository, tokoRepo repository.TokoRepository, AddrUsc AddressUsecase) AuthUseCase {
	return &authUsecaseImpl{
		userRepo: userRepo,
		tokoRepo: tokoRepo,
		AddrUsc: AddrUsc,
	}
}

func (u *authUsecaseImpl) Register(ctx context.Context, req *models.RegisterRequest) (string, error) {

	existing, _ := u.userRepo.CheckEmailPhone(ctx, req.Email, req.NoTelp)
	if existing != nil {
		if existing.Email == req.Email {
			return "", errors.New("email sudah digunakan")
		}
		if existing.NoTelp == req.NoTelp {
			return "", errors.New("nomor telepon sudah digunakan")
		}
	}

	if _, err := u.AddrUsc.GetProvinceDetail(ctx, req.IDProvinsi); err != nil {
		return "", errors.New("provinsi tidak valid")
	}

	if _, err := u.AddrUsc.GetCityDetail(ctx, req.IDProvinsi, req.IDKota); err != nil {
		return "", errors.New("kota tidak valid")
	}


	hash, _ := utils.HashPassword(req.KataSandi)

	user := &entity.User{
		Nama:          req.Nama,
		Email:         req.Email,
		NoTelp:        req.NoTelp,
		KataSandi:     string(hash),
		IDProvinsi:    req.IDProvinsi,
		IDKota:        req.IDKota,
		IsAdmin:       false,
		TanggalLahir:  req.TanggalLahir,
		JenisKelamin:  req.JenisKelamin,
		Tentang:       req.Tentang,
		Pekerjaan:     req.Pekerjaan,
	}

	userCreated, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return "", err
	}

	// wajib create toko
	_, _ = u.tokoRepo.Create(ctx, &entity.Toko{
		IDUser:   userCreated.ID,
		NamaToko: userCreated.Nama + "'s Toko",
	})

	return "Register Succeed", nil
}


func (u *authUsecaseImpl) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {

	user, err := u.userRepo.CheckEmail(ctx, req.Email)
	if err != nil || user == nil {
		return nil, errors.New("email tidak ditemukan")
	}

	if err := utils.VerifyPassword(req.KataSandi, user.KataSandi); err != nil {
		return nil, errors.New("password salah")
	}

	prov, herr := u.AddrUsc.GetProvinceDetail(ctx, user.IDProvinsi)
	if herr != nil {
		return nil, herr.Err
	}

	city, herr := u.AddrUsc.GetCityDetail(ctx, user.IDProvinsi, user.IDKota)
	if herr != nil {
		return nil, herr.Err
	}

	token, err := utils.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	res := &models.LoginResponse{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: user.TanggalLahir,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IDProvinsi: models.ProvinceResponse{
			ID:   prov.ID,
			Name: prov.Name,
		},
		IDKota: models.CityResponse{
			ID:         city.ID,
			ProvinceID: city.ProvinceID,
			Name:       city.Name,
		},
		Token: token,
	}

	return res, nil
}

