package usecase

import (
	"context"
	"errors"
	"fmt"
	"pbi/internal/pkg/entity"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/repository"
	"pbi/internal/utils"
	"time"
)

type AuthUseCase interface {
	Register(ctx context.Context, req *models.RegisterRequest) (*models.UserResponse, error)
	Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error)
}

type authUsecaseImpl struct {
	userRepo repository.UserRepository
	tokoRepo repository.TokoRepository
}

func NewAuthUseCase(userRepo repository.UserRepository, tokoRepo repository.TokoRepository) AuthUseCase {
	return &authUsecaseImpl{
		userRepo: userRepo,
		tokoRepo: tokoRepo,
	}
}

func (u *authUsecaseImpl) Register(ctx context.Context, req *models.RegisterRequest) (*models.UserResponse, error) {
	existing, _ := u.userRepo.CheckEmailPhone(ctx, req.Email, req.NoTelp)

	if existing != nil {
		if existing.Email == req.Email && existing.NoTelp == req.NoTelp {
			return nil, errors.New("email dan nomor telepon sudah digunakan")
		}
		if existing.Email == req.Email {
			return nil, errors.New("email sudah digunakan")
		}
		if existing.NoTelp == req.NoTelp {
			return nil, errors.New("nomor telepon sudah digunakan")
		}
	}

	if ok, err := utils.ValidateProvince(req.IDProvinsi); err != nil || !ok {
		return nil, errors.New("provinsi tidak valid")
	}
	if ok, err := utils.ValidateCity(req.IDProvinsi, req.IDKota); err != nil || !ok {
		return nil, errors.New("kota tidak valid")
	}

	hash, _ := utils.HashPassword(req.KataSandi)

	userEntity := &entity.User{
		Nama:        req.Nama,
		Email:       req.Email,
		NoTelp:      req.NoTelp,
		KataSandi:   string(hash),
		IDProvinsi:  req.IDProvinsi,
		IDKota:      req.IDKota,
		IsAdmin:     false,
		TanggalLahir: req.TanggalLahir,
		JenisKelamin: req.JenisKelamin,
		Tentang:      req.Tentang,
		Pekerjaan:    req.Pekerjaan,
	}
	userCreated, err := u.userRepo.Create(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	tokoEntity := &entity.Toko{
		IDUser:   userCreated.ID,
		NamaToko: userCreated.Nama + "'s Toko",
		UrlFoto:  "",
	}

	_, _ = u.tokoRepo.CreateToko(ctx, tokoEntity)

	res := &models.UserResponse{
		ID:           userCreated.ID,
		Nama:         userCreated.Nama,
		Email:        userCreated.Email,
		NoTelp:       userCreated.NoTelp,
		TanggalLahir: userCreated.TanggalLahir,
		JenisKelamin: userCreated.JenisKelamin,
		Tentang:      userCreated.Tentang,
		Pekerjaan:    userCreated.Pekerjaan,
	}

	return res, nil
}

func (u *authUsecaseImpl) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	start := time.Now()
	user, err := u.userRepo.CheckEmail(ctx, req.Email)
	fmt.Println("CheckEmail:", time.Since(start))

	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("email tidak ditemukan")
	}

	start = time.Now()
	if err := utils.VerifyPassword(req.KataSandi, user.KataSandi); err != nil {
		return nil, errors.New("password salah")
	}
	fmt.Println("Bcrypt verify:", time.Since(start))


	token, err := utils.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	res := &models.LoginResponse{
		ID:           user.ID,
		Nama:         user.Nama,
		Token: token,
	}
	return res, nil
}