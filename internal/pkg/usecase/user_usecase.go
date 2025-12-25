package usecase

import (
	"context"
	"errors"
	"pbi/internal/pkg/entity"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/repository"
	"pbi/internal/utils"
)

type UserUseCase interface {
	UpdateProfile(ctx context.Context, userID int64, req *models.UpdateProfileRequest) (*entity.User, error)
}

type userUsecaseImpl struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUsecaseImpl{
		userRepo: userRepo,
	}
}

func (u *userUsecaseImpl) UpdateProfile(ctx context.Context, userID int64, req *models.UpdateProfileRequest) (*entity.User, error) {
	if req.IDProvinsi != nil && req.IDKota != nil {
		ok, err := utils.ValidateCity(*req.IDProvinsi, *req.IDKota)
		if err != nil || !ok {
			return nil, errors.New("invalid location")
		}
	}

	err := u.userRepo.UpdateProfile(
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
	)

	if err != nil {
		return nil, err
	}

	return u.userRepo.FindById(ctx, userID)
}