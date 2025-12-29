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

// CreateCategory godoc
// @Summary     Create category
// @Description Create new category (Admin only)
// @Tags        Category
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       body body models.CategoryRequest true "Category payload"
// @Success     200 {object} object "Success create category"
// @Failure     400 {object} object "Bad Request"
// @Failure     401 {object} object "Unauthorized"
// @Failure     403 {object} object "Forbidden"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /category [post]
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


// GetAllCategories godoc
// @Summary     Get all categories
// @Description Get list of categories (Admin only)
// @Tags        Category
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} object "Success get category list"
// @Failure     401 {object} object "Unauthorized"
// @Failure     403 {object} object "Forbidden"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /category [get]
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

// GetCategoryByID godoc
// @Summary     Get category by ID
// @Description Get category detail (Admin only)
// @Tags        Category
// @Produce     json
// @Security    BearerAuth
// @Param       id path int true "Category ID"
// @Success     200 {object} object "Success get category detail"
// @Failure     400 {object} object "Invalid ID"
// @Failure     401 {object} object "Unauthorized"
// @Failure     403 {object} object "Forbidden"
// @Failure     404 {object} object "Category not found"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /category/{id} [get]
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

// UpdateCategory godoc
// @Summary     Update category
// @Description Update category by ID (Admin only)
// @Tags        Category
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path int true "Category ID"
// @Param       body body models.UpdateRequest true "Update payload"
// @Success     200 {object} object "Success update category"
// @Failure     400 {object} object "Bad Request"
// @Failure     401 {object} object "Unauthorized"
// @Failure     403 {object} object "Forbidden"
// @Failure     404 {object} object "Category not found"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /category/{id} [put]
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

// DeleteCategory godoc
// @Summary     Delete category
// @Description Delete category by ID (Admin only)
// @Tags        Category
// @Produce     json
// @Security    BearerAuth
// @Param       id path int true "Category ID"
// @Success     200 {object} object "Success delete category"
// @Failure     400 {object} object "Invalid ID"
// @Failure     401 {object} object "Unauthorized"
// @Failure     403 {object} object "Forbidden"
// @Failure     404 {object} object "Category not found"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /category/{id} [delete]
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
