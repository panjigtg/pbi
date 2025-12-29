package controller

import (
	"fmt"
	"os"
	"pbi/internal/pkg/usecase"
	"pbi/internal/pkg/models" 
	"pbi/internal/helper"

	"github.com/gofiber/fiber/v2"
)

type TokoController interface {
	GetAll(ctx *fiber.Ctx) error
	GetMy(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type tokoControllerImpl struct {
	TokoUsc usecase.TokoUsecase
}

func NewTokoController(tokoUsc usecase.TokoUsecase) TokoController {
	return &tokoControllerImpl{
		TokoUsc: tokoUsc,
	}
}

// GetAll godoc
// @Summary     Get all toko
// @Description Get list of all toko (public)
// @Tags        Toko
// @Produce     json
// @Param       page  query int false "Page number"
// @Param       limit query int false "Limit per page"
// @Success     200 {object} object "Success get toko list"
// @Failure     400 {object} object "Bad Request"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /toko [get]
func (tc *tokoControllerImpl) GetAll(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	tokos, err := tc.TokoUsc.GetAll(ctx.Context())
	if err != nil {
		return helper.BadRequest(
			ctx,
			"Failed to GET data",
			err.Err.Error(),
		)
	}

	var resp []models.TokoResponse
	for _, t := range tokos {
		resp = append(resp, models.TokoResponse{
			ID:       t.ID,
			NamaToko: t.NamaToko,
			UrlFoto:  t.UrlFoto,
		})
	}

	data := map[string]interface{}{
		"page":  page,
		"limit": limit,
		"data":  resp,
	}

	return helper.Success(
		ctx,
		"Succeed to GET data",
		data,
	)
}

// GetMy godoc
// @Summary     Get my toko
// @Description Get toko owned by authenticated user
// @Tags        Toko
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} object "Success get my toko"
// @Failure     400 {object} object "Bad Request"
// @Failure     401 {object} object "Unauthorized"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /toko/my [get]
func (tc *tokoControllerImpl) GetMy(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.BadRequest(
			ctx,
			"Failed to GET data",
			"User ID not found",
		)
	}

	tokos, err := tc.TokoUsc.GetMy(ctx.Context(), userID)
	if err != nil {
		return helper.BadRequest(
			ctx,
			"Failed to GET data",
			err.Err.Error(),
		)
	}

	var resp []models.TokoResponse
	for _, t := range tokos {
		resp = append(resp, models.TokoResponse{
			ID:       t.ID,
			NamaToko: t.NamaToko,
			UrlFoto:  t.UrlFoto,
		})
	}

	data := map[string]interface{}{
		"data": resp,
	}

	return helper.Success(
		ctx,
		"Succeed to GET data",
		data,
	)
}

// GetByID godoc
// @Summary     Get toko by ID
// @Description Get detail toko by ID
// @Tags        Toko
// @Produce     json
// @Param       id path int true "Toko ID"
// @Success     200 {object} object "Success get toko detail"
// @Failure     400 {object} object "Invalid toko ID"
// @Failure     404 {object} object "Toko not found"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /toko/{id} [get]
func (tc *tokoControllerImpl) GetByID(ctx *fiber.Ctx) error {
	tokoID := ctx.Params("id")
	if tokoID == "" {
		return helper.BadRequest(ctx, "Failed to GET data", "Toko ID is required")
	}

	toko, err := tc.TokoUsc.GetByID(ctx.Context(), tokoID)
	if err != nil {
		return helper.BadRequest(ctx, "Failed to GET data", err.Err.Error())
	}

	resp := models.TokoResponse{
		ID:       toko.ID,
		NamaToko: toko.NamaToko,
		UrlFoto:  toko.UrlFoto,
	}


	return helper.Success(ctx, "Succeed to GET data", resp)
}


// Update godoc
// @Summary     Update toko
// @Description Update toko data (owner only)
// @Tags        Toko
// @Accept      multipart/form-data
// @Produce     json
// @Security    BearerAuth
// @Param       id        path     int    true  "Toko ID"
// @Param       nama_toko formData string false "Nama toko"
// @Param       url_foto  formData file   false "Foto toko"
// @Success     200 {object} object "Success update toko"
// @Failure     400 {object} object "Bad Request"
// @Failure     401 {object} object "Unauthorized"
// @Failure     403 {object} object "Forbidden"
// @Failure     404 {object} object "Toko not found"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /toko/{id} [put]
func (tc *tokoControllerImpl) Update(ctx *fiber.Ctx) error {
	tokoID := ctx.Params("id")
	if tokoID == "" {
		return helper.BadRequest(ctx, "Failed to UPDATE data", "Toko ID is required")
	}

	userIDRaw := ctx.Locals("user_id")
	var userID int
	switch v := userIDRaw.(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return helper.BadRequest(ctx, "Failed to UPDATE data", "user_id tidak valid")
	}

	namaToko := ctx.FormValue("nama_toko")

	file, fileErr := ctx.FormFile("url_foto")

	if namaToko == "" && fileErr != nil {
    	return helper.BadRequest(ctx, "Failed to UPDATE data", "key salah")
	}
	var filePath string
	if fileErr == nil && file != nil {
		// buat folder uploads/ jika belum ada
		if _, err := os.Stat("uploads"); os.IsNotExist(err) {
			os.Mkdir("uploads", 0755)
		}

		filePath = "uploads/" + file.Filename
		if err := ctx.SaveFile(file, filePath); err != nil {
			return helper.BadRequest(ctx, "Failed to UPDATE data", "Gagal upload foto")
		}
	}

	req := &models.TokoUpdateRequest{
		NamaToko: namaToko,
		UrlFoto:  filePath, // kosong jika tidak upload
	}

	updated, errUc := tc.TokoUsc.Update(ctx.Context(), tokoID, userID, req)
	if errUc != nil {
		// Logging debug untuk owner check
		fmt.Printf("DEBUG Update failed: userID token=%d, tokoID=%s, error=%s\n", userID, tokoID, errUc.Err.Error())
		return helper.BadRequest(ctx, "Failed to UPDATE data", errUc.Err.Error())
	}

	resp := models.TokoResponse{
		ID:       updated.ID,
		NamaToko: updated.NamaToko,
		UrlFoto:  updated.UrlFoto,
	}

	return helper.Success(ctx, "Succeed to UPDATE data", resp)
}



