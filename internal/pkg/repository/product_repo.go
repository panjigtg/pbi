package repository

import (
	"context"
	"pbi/internal/pkg/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, tx *gorm.DB, produk *entity.Produk) error
	CreatePhotos(ctx context.Context, tx *gorm.DB, photos []entity.ProductPhoto) error
	Update(ctx context.Context, tx *gorm.DB, produk *entity.Produk, userID int) error
	Delete(ctx context.Context, tx *gorm.DB, productID int, userID int) error
	DeletePhotosByProduct(ctx context.Context, tx *gorm.DB, productID int) error
	GetAll(ctx context.Context, query *gorm.DB) ([]entity.Produk, error)
	Count(ctx context.Context, query *gorm.DB) (int64, error)
	GetByID(ctx context.Context, productID int) (*entity.Produk, error)
}


type produkrepoImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &produkrepoImpl{
		db: db,
	}
}

func (p *produkrepoImpl) Create(ctx context.Context,tx *gorm.DB,produk *entity.Produk,) error {
	return tx.WithContext(ctx).
		Table("produk").
		Create(produk).
		Error
}

func (p *produkrepoImpl) CreatePhotos(ctx context.Context,tx *gorm.DB,photos []entity.ProductPhoto,) error {
	if len(photos) < 0 {
		return nil
	}

	return tx.WithContext(ctx).
		Table("foto_produk").
		Create(&photos).
		Error
}

func (p *produkrepoImpl) Update(ctx context.Context,tx *gorm.DB,produk *entity.Produk,userID int,) error {
	res := tx.WithContext(ctx).
		Table("produk").
		Where("id = ? AND id_toko = ?", produk.ID, userID).
		Updates(produk)

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}

func (p *produkrepoImpl) Delete(ctx context.Context,tx *gorm.DB,productID int,userID int,) error {
	res := tx.WithContext(ctx).
		Table("produk").
		Where("id = ? AND id_toko = ?", productID, userID).
		Delete(nil)

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}

func (p *produkrepoImpl) DeletePhotosByProduct(ctx context.Context,tx *gorm.DB,productID int,) error {
	return tx.WithContext(ctx).
		Table("foto_produk").
		Where("id_produk = ?", productID).
		Delete(nil).
		Error
}

func (p *produkrepoImpl) GetAll(ctx context.Context, query *gorm.DB) ([]entity.Produk, error) {
	var products []entity.Produk

	err := query.WithContext(ctx).
		Preload("Toko").
		Preload("Category").
		Preload("Photos").
		Find(&products).
		Error

	return products, err
}

func (p *produkrepoImpl) Count(ctx context.Context, query *gorm.DB) (int64, error) {
	var total int64
	err := query.WithContext(ctx).
		Model(&entity.Produk{}).
		Count(&total).
		Error
	return total, err
}

func (p *produkrepoImpl) GetByID(ctx context.Context, productID int) (*entity.Produk, error) {
	var product entity.Produk

	err := p.db.WithContext(ctx).
		Table("produk").
		Preload("Toko").
		Preload("Category").
		Preload("Photos").
		Where("id = ?", productID).
		First(&product).
		Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}