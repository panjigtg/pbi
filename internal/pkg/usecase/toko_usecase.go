package usecase

import (
	"context"
	"errors"
	"fmt"
	"pbi/internal/helper"
	"pbi/internal/pkg/entity"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)


type TokoUsecase interface {
	GetAll(ctx context.Context) ([]*entity.Toko, *helper.ErrorStruct)
	GetMy(ctx context.Context, userID int) ([]*entity.Toko, *helper.ErrorStruct)
	GetByID(ctx context.Context, tokoID string) (*entity.Toko, *helper.ErrorStruct)
	Update(ctx context.Context, tokoID string, userID int, req *models.TokoUpdateRequest) (*entity.Toko, *helper.ErrorStruct)
}

type tokoUsecaseImpl struct {
	tokoRepo repository.TokoRepository
}

func NewTokoUsecase(repo repository.TokoRepository) TokoUsecase {
	return &tokoUsecaseImpl{tokoRepo: repo}
}

// GetAll mengambil semua toko
func (uc *tokoUsecaseImpl) GetAll(ctx context.Context) ([]*entity.Toko, *helper.ErrorStruct) {
	tokos, errRepo := uc.tokoRepo.GetAll(ctx)
	if errors.Is(errRepo, fiber.ErrNotFound) {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("tidak ada data toko"),
		}
	}
	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	return tokos, nil
}

// GetMy mengambil toko user tertentu
func (uc *tokoUsecaseImpl) GetMy(ctx context.Context, userID int) ([]*entity.Toko, *helper.ErrorStruct) {
	if userID <= 0 {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("unauthorized"),
		}
	}

	tokos, errRepo := uc.tokoRepo.GetByUserID(ctx, userID)
	if errors.Is(errRepo, fiber.ErrNotFound) {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("toko tidak ditemukan"),
		}
	}
	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	if len(tokos) == 0 {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("toko tidak ditemukan"),
		}
	}

	return tokos, nil
}

func (uc *tokoUsecaseImpl) GetByID(ctx context.Context, tokoID string) (*entity.Toko, *helper.ErrorStruct) {
	id, err := strconv.Atoi(tokoID)
	if err != nil {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errors.New("ID toko tidak valid"),
		}
	}

	data, errRepo := uc.tokoRepo.GetByID(ctx, id)
	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("toko tidak ditemukan"),
		}
	}

	return data, nil
}

// Update toko
func (uc *tokoUsecaseImpl) Update(ctx context.Context, tokoID string, userID int, req *models.TokoUpdateRequest) (*entity.Toko, *helper.ErrorStruct) {
	id, err := strconv.Atoi(tokoID)
	if err != nil {
		return nil, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New("ID toko tidak valid")}
	}

	// Ambil toko
	toko, errRepo := uc.tokoRepo.GetByID(ctx, id)
	if errRepo != nil {
		return nil, &helper.ErrorStruct{Code: fiber.StatusInternalServerError, Err: errRepo}
	}
	if toko == nil {
		return nil, &helper.ErrorStruct{Code: fiber.StatusNotFound, Err: errors.New("toko tidak ditemukan")}
	}
	fmt.Printf("DEBUG toko.IDUser: %v, tokoID: %v\n", toko.IDUser, toko.ID)


	// Ownership
	if toko.IDUser != userID {
		return nil, &helper.ErrorStruct{Code: fiber.StatusForbidden, Err: errors.New("akses ditolak")}
	}

	// Field fallback: jika kosong, gunakan field lama
	if req.NamaToko != "" {
		toko.NamaToko = req.NamaToko
	}
	if req.UrlFoto != "" {
		toko.UrlFoto = req.UrlFoto
	}

	updated, errRepo := uc.tokoRepo.Update(ctx, id, toko)
	if errRepo != nil {
		return nil, &helper.ErrorStruct{Code: fiber.StatusInternalServerError, Err: errRepo}
	}

	return updated, nil
}

