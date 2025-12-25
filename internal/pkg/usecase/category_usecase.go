package usecase

import (
	"context"

	"pbi/internal/helper"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/entity"
	"pbi/internal/pkg/repository"
)

type CategoryUseCase interface {
	CreateCategory(ctx context.Context, category *models.CategoryRequest) (res uint, err *helper.ErrorStruct)
	GetAllCategories(ctx context.Context) (res []*models.CategoryResponse, err *helper.ErrorStruct)
	GetById(ctx context.Context, id int)(res *models.CategoryResponse, err *helper.ErrorStruct)
}

type categoryUsecaseImpl struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUseCase(categoryRepo repository.CategoryRepository) CategoryUseCase {
	return &categoryUsecaseImpl{
		categoryRepo: categoryRepo,
	}
}

func (u *categoryUsecaseImpl) CreateCategory(ctx context.Context, category *models.CategoryRequest) (res uint, err *helper.ErrorStruct)  {
	newCategory := &entity.Category{
		Nama: category.Nama,
	}
	
	if e := u.categoryRepo.Create(ctx, newCategory); e != nil {
        return 0, &helper.ErrorStruct{
            Err:  e,
            Code: 500,
        }
    }

	return uint(newCategory.ID), nil
}

func (u *categoryUsecaseImpl) GetAllCategories(ctx context.Context) (res []*models.CategoryResponse, err *helper.ErrorStruct) {
	categories, e := u.categoryRepo.GetAll(ctx)
	if e != nil {
		return nil, &helper.ErrorStruct{
			Err:  e,
			Code: 500,
		}
	}

	var categoryResponses []*models.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, &models.CategoryResponse{
			ID:        category.ID,
			Nama:      category.Nama,
		})
	}

	return categoryResponses, nil
}


func (u *categoryUsecaseImpl) GetById(ctx context.Context, id int)(res *models.CategoryResponse, err *helper.ErrorStruct){
	category, e := u.categoryRepo.GetById(ctx, id)
	if e != nil {
		return nil, &helper.ErrorStruct{
			Err:  e,
			Code: 500,
		}
	}

	res = &models.CategoryResponse{
		ID:   	category.ID,
        Nama: 	category.Nama,
	}

	return res, nil
}