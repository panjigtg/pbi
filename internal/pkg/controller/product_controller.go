package controller

import (
	"pbi/internal/helper"
	"pbi/internal/pkg/entity"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
}

type productImpl struct {
	puc usecase.ProductUsecase
}

func NewProductController(uc usecase.ProductUsecase) ProductController {
	return &productImpl{
		puc: uc,
	}
}

// CreateProduct godoc
// @Summary     Create product
// @Description Create new product with images (multipart/form-data)
// @Tags        Product
// @Accept      multipart/form-data
// @Produce     json
// @Security    BearerAuth
// @Param       nama_produk      formData string true  "Product name"
// @Param       slug             formData string true  "Product slug"
// @Param       deskripsi        formData string true  "Product description"
// @Param       id_category      formData int    true  "Category ID"
// @Param       harga_reseller   formData number true  "Reseller price"
// @Param       harga_konsumen   formData number true  "Consumer price"
// @Param       stok             formData int    true  "Stock"
// @Param       photos           formData file   false "Product images (multiple)"
// @Success     200 {object} object "Success create product"
// @Failure     400 {object} object "Bad Request"
// @Failure     401 {object} object "Unauthorized"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /product [post]
func (c *productImpl) Create(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.Unauthorized(ctx, "Unauthorized")
	}

	var req models.ProductCreateReq

	req.NamaProduk = ctx.FormValue("nama_produk")
	req.Slug = ctx.FormValue("slug")
	req.Deskripsi = ctx.FormValue("deskripsi")

	if v := ctx.FormValue("id_category"); v != "" {
		id, err := strconv.Atoi(v)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid id_category", err.Error())
		}
		req.IDCategory = id
	}

	if v := ctx.FormValue("harga_reseller"); v != "" {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid harga_reseller", err.Error())
		}
		req.HargaReseller = val
	}

	if v := ctx.FormValue("harga_konsumen"); v != "" {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid harga_konsumen", err.Error())
		}
		req.HargaKonsumen = val
	}

	if v := ctx.FormValue("stok"); v != "" {
		val, err := strconv.Atoi(v)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid stok", err.Error())
		}
		req.Stok = val
	}

	form, err := ctx.MultipartForm()
	if err == nil {
		if files, ok := form.File["photos"]; ok {
			req.Photos = files
		}
	}

	id, herr := c.puc.Create(ctx.Context(), &req, userID)
	if herr != nil {
		return helper.Error(ctx, herr.Code, "Failed to POST data", herr.Err.Error())
	}

	return helper.Success(ctx, "Succeed to POST data", strconv.Itoa(id))
}

// UpdateProduct godoc
// @Summary     Update product
// @Description Update product by ID (multipart/form-data)
// @Tags        Product
// @Accept      multipart/form-data
// @Produce     json
// @Security    BearerAuth
// @Param       id               path     int    true  "Product ID"
// @Param       nama_produk      formData string false "Product name"
// @Param       slug             formData string false "Product slug"
// @Param       deskripsi        formData string false "Product description"
// @Param       id_category      formData int    false "Category ID"
// @Param       harga_reseller   formData number false "Reseller price"
// @Param       harga_konsumen   formData number false "Consumer price"
// @Param       stok             formData int    false "Stock"
// @Param       photos           formData file   false "Product images (multiple)"
// @Success     200 {object} object "Success update product"
// @Failure     400 {object} object "Bad Request"
// @Failure     401 {object} object "Unauthorized"
// @Failure     404 {object} object "Product not found"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /product/{id} [put]
func (c *productImpl) Update(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.Unauthorized(ctx, "Unauthorized")
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequest(ctx, "Invalid product id", err.Error())
	}

	var req models.ProductCreateReq
	req.NamaProduk = ctx.FormValue("nama_produk")
	req.Slug = ctx.FormValue("slug")
	req.Deskripsi = ctx.FormValue("deskripsi")

	if v := ctx.FormValue("id_category"); v != "" {
		catID, err := strconv.Atoi(v)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid id_category", err.Error())
		}
		req.IDCategory = catID
	}

	if v := ctx.FormValue("harga_reseller"); v != "" {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid harga_reseller", err.Error())
		}
		req.HargaReseller = val
	}

	if v := ctx.FormValue("harga_konsumen"); v != "" {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid harga_konsumen", err.Error())
		}
		req.HargaKonsumen = val
	}

	if v := ctx.FormValue("stok"); v != "" {
		stok, err := strconv.Atoi(v)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid stok", err.Error())
		}
		req.Stok = stok
	}

	form, err := ctx.MultipartForm()
	if err == nil {
		if files, ok := form.File["photos"]; ok {
			req.Photos = files
		}
	}

	if herr := c.puc.Update(ctx.Context(), id, &req, userID); herr != nil {
		return helper.Error(ctx, herr.Code, "Failed to UPDATE data", herr.Err.Error())
	}

	return helper.Success(ctx, "Succeed to UPDATE data", nil)
}

// DeleteProduct godoc
// @Summary     Delete product
// @Description Delete product by ID
// @Tags        Product
// @Produce     json
// @Security    BearerAuth
// @Param       id path int true "Product ID"
// @Success     200 {object} object "Success delete product"
// @Failure     400 {object} object "Invalid product ID"
// @Failure     401 {object} object "Unauthorized"
// @Failure     404 {object} object "Product not found"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /product/{id} [delete]
func (c *productImpl) Delete(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return helper.Unauthorized(ctx, "Unauthorized")
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequest(ctx, "Invalid product id", err.Error())
	}

	if herr := c.puc.Delete(ctx.Context(), id, userID); herr != nil {
		return helper.Error(ctx, herr.Code, "Failed to DELETE data", herr.Err.Error())
	}

	return helper.Success(ctx, "Succeed to DELETE data", nil)
}

// GetAllProducts godoc
// @Summary     Get all products
// @Description Get product list with filters & pagination
// @Tags        Product
// @Produce     json
// @Param       nama_produk query string false "Product name"
// @Param       category_id query int    false "Category ID"
// @Param       toko_id     query int    false "Toko ID"
// @Param       min_harga   query number false "Minimum price"
// @Param       max_harga   query number false "Maximum price"
// @Param       page        query int    false "Page number"
// @Param       limit       query int    false "Limit per page"
// @Success     200 {object} object "Success get product list"
// @Failure     400 {object} object "Bad Request"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /product [get]
func (c *productImpl) GetAll(ctx *fiber.Ctx) error {
	// Parse query parameters
	filter := &entity.ProductFilter{
		NamaProduk: ctx.Query("nama_produk"),
	}

	if categoryID := ctx.Query("category_id"); categoryID != "" {
		id, err := strconv.Atoi(categoryID)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid category_id", err.Error())
		}
		filter.CategoryID = id
	}

	if tokoID := ctx.Query("toko_id"); tokoID != "" {
		id, err := strconv.Atoi(tokoID)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid toko_id", err.Error())
		}
		filter.TokoID = id
	}

	if minHarga := ctx.Query("min_harga"); minHarga != "" {
		val, err := strconv.ParseFloat(minHarga, 64)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid min_harga", err.Error())
		}
		filter.MinHarga = val
	}

	if maxHarga := ctx.Query("max_harga"); maxHarga != "" {
		val, err := strconv.ParseFloat(maxHarga, 64)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid max_harga", err.Error())
		}
		filter.MaxHarga = val
	}

	if page := ctx.Query("page"); page != "" {
		p, err := strconv.Atoi(page)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid page", err.Error())
		}
		filter.Page = p
	}

	if limit := ctx.Query("limit"); limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return helper.BadRequest(ctx, "Invalid limit", err.Error())
		}
		filter.Limit = l
	}

	// Call usecase
	result, herr := c.puc.GetAll(ctx.Context(), filter)
	if herr != nil {
		return helper.Error(ctx, herr.Code, "Failed to GET data", herr.Err.Error())
	}

	return helper.Success(ctx, "Succeed to GET data", result)
}

// GetProductByID godoc
// @Summary     Get product by ID
// @Description Get product detail
// @Tags        Product
// @Produce     json
// @Param       id path int true "Product ID"
// @Success     200 {object} object "Success get product detail"
// @Failure     400 {object} object "Invalid product ID"
// @Failure     404 {object} object "Product not found"
// @Failure     500 {object} object "Internal Server Error"
// @Router      /product/{id} [get]
func (c *productImpl) GetByID(ctx *fiber.Ctx) error {
	// Parse product_id from params
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequest(ctx, "Invalid product id", err.Error())
	}

	// Call usecase
	product, herr := c.puc.GetByID(ctx.Context(), id)
	if herr != nil {
		return helper.Error(ctx, herr.Code, "Failed to GET data", herr.Err.Error())
	}

	return helper.Success(ctx, "Succeed to GET data", product)
}