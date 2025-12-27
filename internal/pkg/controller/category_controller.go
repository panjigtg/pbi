package controller

import (
	"pbi/internal/helper"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	CreateCategory(ctx *fiber.Ctx) error
	GetAllCategories(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type categoryControllerImpl struct {
	cUsc usecase.CategoryUseCase
}

func NewCategoryController(categoryUseCase usecase.CategoryUseCase) CategoryController {
	return &categoryControllerImpl{
		cUsc: categoryUseCase,
	}
}

func (cc *categoryControllerImpl) CreateCategory(ctx *fiber.Ctx) error {
	req := new(models.CategoryRequest)
	if err := ctx.BodyParser(req); err != nil {
		return helper.BadRequest(
			ctx,
			"Failed to POST data",
			"Invalid request body",
		)
	}


	res, err := cc.cUsc.CreateCategory(ctx.Context(), req)
	if err != nil {
		return helper.BadRequest(
			ctx,
			"Failed to POST data",
			err.Err.Error(),
		)
	}
	return helper.Success(ctx, "Succeed to POST data", res)
}

func (cc *categoryControllerImpl) GetAllCategories(ctx *fiber.Ctx) error {
	res, err := cc.cUsc.GetAllCategories(ctx.Context())
	if err != nil {
		return helper.BadRequest(
			ctx,
			"Failed to GET data",
			err.Err.Error(),
		)
	}
	return helper.Success(ctx, "Succeed to GET data", res)
}

func (cc *categoryControllerImpl) GetById(ctx *fiber.Ctx) error {

	idParam := ctx.Params("id")
    id, convErr := strconv.Atoi(idParam)
    if convErr != nil {
        return helper.BadRequest(ctx, "Invalid ID", convErr.Error())
    }

	res, err := cc.cUsc.GetById(ctx.Context(), id)
	if err != nil {
		return helper.BadRequest(
			ctx,
			"Failed to GET data",
			err.Err.Error(),
		)
	}
	return helper.Success(ctx, "Succeed to GET data", res)
}

func (cc *categoryControllerImpl) Update(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return helper.BadRequest(ctx, "Invalid ID", err.Error())
    }

	req := &models.UpdateRequest{}
	if err := ctx.BodyParser(req); err != nil {
        return helper.BadRequest(ctx, "Invalid request body", err.Error())
    }

	// if req.Nama == "" {
    // return helper.BadRequest(ctx, "Nama category tidak boleh kosong", "Invalid request body")
	// }


	 if errStruct := cc.cUsc.Update(ctx.Context(), id, req); errStruct != nil {
        return helper.BadRequest(ctx, "Failed to update category", errStruct.Err.Error())
    }

    return helper.Success(ctx, "Category updated successfully", nil)
}

func (cc *categoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return helper.BadRequest(ctx, "Invalid ID", err.Error())
	}

	if errStruct := cc.cUsc.Delete(ctx.Context(), id); errStruct != nil {
		return helper.BadRequest(ctx, "Failed to delete category", errStruct.Err.Error())
	}

	return helper.Success(ctx, "Category deleted successfully", nil)
}
