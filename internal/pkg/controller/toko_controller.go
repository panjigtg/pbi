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

type TokoControllerImpl struct {
	TokoUsc usecase.TokoUsecase
}

func NewTokoController(tokoUsc usecase.TokoUsecase) *TokoControllerImpl {
	return &TokoControllerImpl{
		TokoUsc: tokoUsc,
	}
}

func (tc *TokoControllerImpl) GetAll(ctx *fiber.Ctx) error {
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

func (tc *TokoControllerImpl) GetMy(ctx *fiber.Ctx) error {
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

func (tc *TokoControllerImpl) GetByID(ctx *fiber.Ctx) error {
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


func (tc *TokoControllerImpl) Update(ctx *fiber.Ctx) error {
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



