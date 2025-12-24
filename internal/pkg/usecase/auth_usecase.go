package usecase

import (
	"context"
	"errors"
	"pbi/internal/pkg/entity"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/repository"
	"pbi/internal/utils"
)

type AuthUseCase interface {
	Register(ctx context.Context, req *models.RegisterRequest) (*models.UserResponse, error)
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

	hash, _ := utils.HashPassword(req.KataSandi)

	userEntity := &entity.User{
		Nama:       req.Nama,
		Email:      req.Email,
		NoTelp:     req.NoTelp,
		KataSandi:  string(hash),
		IDProvinsi: req.IDProvinsi,
		IDKota:     req.IDKota,
		IsAdmin:    false,
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
		ID:     userCreated.ID,
		Nama:   userCreated.Nama,
		Email:  userCreated.Email,
		NoTelp: userCreated.NoTelp,
	}

	return res, nil
}
