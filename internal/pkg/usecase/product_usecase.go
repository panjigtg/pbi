package usecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"pbi/internal/helper"
	"pbi/internal/pkg/entity"
	"pbi/internal/pkg/models"
	"pbi/internal/pkg/repository"
	"time"

	"gorm.io/gorm"
	"github.com/gosimple/slug"
)

type ProductUsecase interface {
	Create(ctx context.Context, req *models.ProductCreateReq, UserID int)(int, *helper.ErrorStruct)
	Update(ctx context.Context, productID int, req *models.ProductCreateReq, userID int) *helper.ErrorStruct
	Delete(ctx context.Context, productID int, userID int) *helper.ErrorStruct
	GetAll(ctx context.Context, filter *entity.ProductFilter) (*models.ProductListResponse, *helper.ErrorStruct)
	GetByID(ctx context.Context, productID int) (*models.ProductResponse, *helper.ErrorStruct)
}

type produkusecaseImpl struct {
	db   *gorm.DB
	prepo repository.ProductRepository
}



func NewProductUsecase(db *gorm.DB,prepo repository.ProductRepository) ProductUsecase {
	return &produkusecaseImpl{
		prepo: prepo,
		db: db,
	}
}

func uploadFile(file *multipart.FileHeader) (string, error) {
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
	return filename, nil
}

func (p *produkusecaseImpl) Create(ctx context.Context,req *models.ProductCreateReq,userID int,) (int, *helper.ErrorStruct) {

	tx := p.db.Begin()
	if tx.Error != nil {
		return 0, &helper.ErrorStruct{Err: tx.Error, Code: 500}
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	slugText := req.Slug
	if slugText == "" {
		slugText = slug.Make(req.NamaProduk)
	}

	produk := &entity.Produk{
		IDToko:        userID,
		IDCategory:    req.IDCategory,
		NamaProduk:    req.NamaProduk,
		Slug:          slugText,
		HargaReseller: req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
	}

	if err := p.prepo.Create(ctx, tx, produk); err != nil {
		tx.Rollback()
		return 0, &helper.ErrorStruct{Err: err, Code: 500}
	}

	var photos []entity.ProductPhoto
	for _, file := range req.Photos {
		url, err := uploadFile(file)
		if err != nil {
			tx.Rollback()
			return 0, &helper.ErrorStruct{Err: err, Code: 400}
		}

		photos = append(photos, entity.ProductPhoto{
			ProductID: produk.ID,
			Url:       url,
		})
	}

	if err := p.prepo.CreatePhotos(ctx, tx, photos); err != nil {
		tx.Rollback()
		return 0, &helper.ErrorStruct{Err: err, Code: 500}
	}

	if err := tx.Commit().Error; err != nil {
		return 0, &helper.ErrorStruct{Err: err, Code: 500}
	}

	return produk.ID, nil
}

func (p *produkusecaseImpl) Update(ctx context.Context,productID int,req *models.ProductCreateReq,userID int,) *helper.ErrorStruct {

	tx := p.db.Begin()
	if tx.Error != nil {
		return &helper.ErrorStruct{Err: tx.Error, Code: 500}
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	slugText := req.Slug
	if slugText == "" {
		slugText = slug.Make(req.NamaProduk)
	}

	produk := &entity.Produk{
		ID:            productID,
		IDCategory:    req.IDCategory,
		NamaProduk:    req.NamaProduk,
		Slug:          slugText,
		HargaReseller: req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
	}

	if err := p.prepo.Update(ctx, tx, produk, userID); err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return &helper.ErrorStruct{Err: err, Code: 404}
		}
		return &helper.ErrorStruct{Err: err, Code: 500}
	}

	if len(req.Photos) > 0 {
		if err := p.prepo.DeletePhotosByProduct(ctx, tx, productID); err != nil {
			tx.Rollback()
			return &helper.ErrorStruct{Err: err, Code: 500}
		}

		var photos []entity.ProductPhoto
		for _, file := range req.Photos {
			url, err := uploadFile(file)
			if err != nil {
				tx.Rollback()
				return &helper.ErrorStruct{Err: err, Code: 400}
			}
			photos = append(photos, entity.ProductPhoto{
				ProductID: productID,
				Url:       url,
			})
		}

		if err := p.prepo.CreatePhotos(ctx, tx, photos); err != nil {
			tx.Rollback()
			return &helper.ErrorStruct{Err: err, Code: 500}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return &helper.ErrorStruct{Err: err, Code: 500}
	}

	return nil
}

func (p *produkusecaseImpl) Delete(ctx context.Context,productID int,userID int,) *helper.ErrorStruct {

	tx := p.db.Begin()
	if tx.Error != nil {
		return &helper.ErrorStruct{Err: tx.Error, Code: 500}
	}

	if err := p.prepo.DeletePhotosByProduct(ctx, tx, productID); err != nil {
		tx.Rollback()
		return &helper.ErrorStruct{Err: err, Code: 500}
	}

	if err := p.prepo.Delete(ctx, tx, productID, userID); err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return &helper.ErrorStruct{Err: err, Code: 404}
		}
		return &helper.ErrorStruct{Err: err, Code: 500}
	}

	if err := tx.Commit().Error; err != nil {
		return &helper.ErrorStruct{Err: err, Code: 500}
	}

	return nil
}

func (p *produkusecaseImpl) GetAll(ctx context.Context, filter *entity.ProductFilter) (*models.ProductListResponse, *helper.ErrorStruct) {
	// Set default pagination
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	if filter.Page <= 0 {
		filter.Page = 1
	}

	// Build query with filters
	query := p.db.Table("produk")

	if filter.NamaProduk != "" {
		query = query.Where("nama_produk LIKE ?", "%"+filter.NamaProduk+"%")
	}

	if filter.CategoryID > 0 {
		query = query.Where("id_category = ?", filter.CategoryID)
	}

	if filter.TokoID > 0 {
		query = query.Where("id_toko = ?", filter.TokoID)
	}

	if filter.MinHarga > 0 {
		query = query.Where("harga_konsumen >= ?", filter.MinHarga)
	}

	if filter.MaxHarga > 0 {
		query = query.Where("harga_konsumen <= ?", filter.MaxHarga)
	}

	// Count total records
	total, err := p.prepo.Count(ctx, query)
	if err != nil {
		return nil, &helper.ErrorStruct{Err: err, Code: 500}
	}

	// Apply pagination
	offset := (filter.Page - 1) * filter.Limit
	query = query.Limit(filter.Limit).Offset(offset)

	// Get products
	products, err := p.prepo.GetAll(ctx, query)
	if err != nil {
		return nil, &helper.ErrorStruct{Err: err, Code: 500}
	}

	// Map entity to response langsung di sini
	productResponses := make([]models.ProductResponse, 0, len(products))
	for _, product := range products {
		productResp := models.ProductResponse{
			ID:            product.ID,
			NamaProduk:    product.NamaProduk,
			Slug:          product.Slug,
			HargaReseller: product.HargaReseller,
			HargaKonsumen: product.HargaKonsumen,
			Stok:          product.Stok,
			Deskripsi:     product.Deskripsi,
		}

		// Map Toko
		if product.Toko != nil {
			productResp.Toko = &models.TokoResponse{
				ID:       product.Toko.ID,
				NamaToko: product.Toko.NamaToko,
				UrlFoto:  product.Toko.UrlFoto,
			}
		}

		// Map Category
		if product.Category != nil {
			productResp.Category = &models.CategoryResponse{
				ID:           product.Category.ID,
				Nama: product.Category.Nama,
			}
		}

		// Map Photos
		productResp.Photos = make([]models.ProductPhotoResponse, 0, len(product.Photos))
		for _, photo := range product.Photos {
			productResp.Photos = append(productResp.Photos, models.ProductPhotoResponse{
				ID:        photo.ID,
				ProductID: photo.ProductID,
				Url:       photo.Url,
			})
		}

		productResponses = append(productResponses, productResp)
	}

	// Calculate total pages
	totalPages := int((total + int64(filter.Limit) - 1) / int64(filter.Limit))

	response := &models.ProductListResponse{
		Data:       productResponses,
		Total:      total,
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalPages: totalPages,
	}

	return response, nil
}

func (p *produkusecaseImpl) GetByID(ctx context.Context, productID int) (*models.ProductResponse, *helper.ErrorStruct) {
	product, err := p.prepo.GetByID(ctx, productID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &helper.ErrorStruct{Err: err, Code: 404}
		}
		return nil, &helper.ErrorStruct{Err: err, Code: 500}
	}

	// Map entity to response langsung di sini
	response := &models.ProductResponse{
		ID:            product.ID,
		NamaProduk:    product.NamaProduk,
		Slug:          product.Slug,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     product.Deskripsi,
	}

	// Map Toko
	if product.Toko != nil {
		response.Toko = &models.TokoResponse{
			ID:       product.Toko.ID,
			NamaToko: product.Toko.NamaToko,
			UrlFoto:  product.Toko.UrlFoto,
		}
	}

	// Map Category
	if product.Category != nil {
		response.Category = &models.CategoryResponse{
			ID:           	product.Category.ID,
			Nama: 			product.Category.Nama,
		}
	}

	// Map Photos
	response.Photos = make([]models.ProductPhotoResponse, 0, len(product.Photos))
	for _, photo := range product.Photos {
		response.Photos = append(response.Photos, models.ProductPhotoResponse{
			ID:        photo.ID,
			ProductID: photo.ProductID,
			Url:       photo.Url,
		})
	}

	return response, nil
}