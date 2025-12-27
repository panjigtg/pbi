package usecase

import (
	"context"
	"errors"
	"pbi/internal/helper"
	"pbi/internal/pkg/entity"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/repository"

	"gorm.io/gorm"
)

type DestinationUsecase interface {
	Create(ctx context.Context,userID int, req *models.DestinationCreateRequest)(*models.DestinationResponse,*helper.ErrorStruct)
	Update(ctx context.Context,id int,userID int,req *models.DestinationUpdateRequest,) (*models.DestinationResponse, *helper.ErrorStruct)
	Delete(ctx context.Context,id int,userID int,) *helper.ErrorStruct
	GetByID(ctx context.Context, id int, userID int) (*models.DestinationResponse, *helper.ErrorStruct)
	GetAll(ctx context.Context, userID int) ([]models.DestinationResponse, *helper.ErrorStruct)
}

type destinationImpl struct {
	repo repository.DestinationRepository
}

func NewDestinationUsecase(repo repository.DestinationRepository) DestinationUsecase{
	return &destinationImpl{
		repo: repo,
	}
}


func (d *destinationImpl) Create(ctx context.Context, userID int, req *models.DestinationCreateRequest)(*models.DestinationResponse,*helper.ErrorStruct){
	dest := &entity.Destination{
		UserID:        userID,
		JudulAlamat:   req.JudulAlamat,
		NamaPenerima:  req.NamaPenerima,
		NoTelp:        req.NoTelp,
		DetailAlamat: req.DetailAlamat,
	}

	if err := d.repo.Create(ctx, dest); err != nil {
		return nil, &helper.ErrorStruct{
			Code: 500,
			Err:  err,
		}
	}

	res := &models.DestinationResponse{
		ID:           dest.ID,
		JudulAlamat:  dest.JudulAlamat,
		NamaPenerima: dest.NamaPenerima,
		NoTelp:       dest.NoTelp,
		DetailAlamat: dest.DetailAlamat,
	}

	return res, nil

}

func (d *destinationImpl) Update(ctx context.Context,id int,userID int,req *models.DestinationUpdateRequest,) (*models.DestinationResponse, *helper.ErrorStruct) {

	existing, err := d.repo.FindByID(ctx, id, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &helper.ErrorStruct{
				Code: 404,
				Err:  errors.New("destination not found"),
			}
		}
		return nil, &helper.ErrorStruct{
			Code: 500,
			Err:  err,
		}
	}

	if req.NamaPenerima != "" {
		existing.NamaPenerima = req.NamaPenerima
	}
	if req.NoTelp != "" {
		existing.NoTelp = req.NoTelp
	}
	if req.DetailAlamat != "" {
		existing.DetailAlamat = req.DetailAlamat
	}

	if err := d.repo.Update(ctx, existing); err != nil {
		return nil, &helper.ErrorStruct{
			Code: 500,
			Err:  err,
		}
	}

	res := &models.DestinationResponse{
		ID:           existing.ID,
		JudulAlamat:  existing.JudulAlamat,
		NamaPenerima: existing.NamaPenerima,
		NoTelp:       existing.NoTelp,
		DetailAlamat: existing.DetailAlamat,
	}

	return res, nil
}


func (d *destinationImpl) Delete(ctx context.Context,id int,userID int,) *helper.ErrorStruct {

	if err := d.repo.Delete(ctx, id, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &helper.ErrorStruct{
				Code: 404,
				Err:  errors.New("destination not found"),
			}
		}

		return &helper.ErrorStruct{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

func (d *destinationImpl) GetByID(ctx context.Context,id int,userID int,) (*models.DestinationResponse, *helper.ErrorStruct) {

	dest, err := d.repo.FindByID(ctx, id, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &helper.ErrorStruct{
				Code: 404,
				Err:  errors.New("destination not found"),
			}
		}
		return nil, &helper.ErrorStruct{
			Code: 500,
			Err:  err,
		}
	}

	return &models.DestinationResponse{
		ID:           dest.ID,
		JudulAlamat:  dest.JudulAlamat,
		NamaPenerima: dest.NamaPenerima,
		NoTelp:       dest.NoTelp,
		DetailAlamat: dest.DetailAlamat,
	}, nil
}

func (d *destinationImpl) GetAll(ctx context.Context,userID int,) ([]models.DestinationResponse, *helper.ErrorStruct) {

	dests, err := d.repo.FindAllByUserID(ctx, userID)
	if err != nil {
		return nil, &helper.ErrorStruct{
			Code: 500,
			Err:  err,
		}
	}

	res := make([]models.DestinationResponse, 0, len(dests))
	for _, dest := range dests {
		res = append(res, models.DestinationResponse{
			ID:           dest.ID,
			JudulAlamat:  dest.JudulAlamat,
			NamaPenerima: dest.NamaPenerima,
			NoTelp:       dest.NoTelp,
			DetailAlamat: dest.DetailAlamat,
		})
	}

	return res, nil
}
