package repository

import (
	"context"
	"pbi/internal/pkg/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, tx *gorm.DB, trx *entity.Transaction) error
	CreateLogProduk(ctx context.Context, tx *gorm.DB, log *entity.LogProduk) error
	CreateDetailTransaction(ctx context.Context, tx *gorm.DB, detail *entity.DetailTransaction) error
	UpdateTotalHarga(ctx context.Context,tx *gorm.DB,trxID int,total float64,) error
	UpdateStokProduk(ctx context.Context, tx *gorm.DB, produkID int, qty int) error
	GetProdukByID(ctx context.Context, tx *gorm.DB, productID int) (*entity.ProdukWithOwner, error)
	GetTransactionsByUser(ctx context.Context, userID int) ([]entity.Transaction, error)
	GetTransactionDetails(ctx context.Context, trxID int) ([]entity.DetailTrxWithJoin, error)
	GetProductPhotosByLogProdukID(ctx context.Context, logProdukID int) ([]entity.ProductPhoto, error)
	GetTransactionByID(ctx context.Context, trxID int, userID int) (*entity.Transaction, error)
}

type transactionImpl struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepository {
	return &transactionImpl{
		db: db,
	}
}

func (t *transactionImpl) CreateTransaction(ctx context.Context, tx *gorm.DB, trx *entity.Transaction) error{
	return tx.WithContext(ctx).Create(trx).Error
}

// func (r *transactionImpl) GetProdukByID(ctx context.Context,tx *gorm.DB,productID int,) (*entity.Produk, error) {

// 	var produk entity.Produk
// 	err := tx.WithContext(ctx).
// 		Where("id = ?", productID).
// 		First(&produk).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &produk, nil
// }

func (r *transactionImpl) CreateLogProduk(ctx context.Context,tx *gorm.DB,log *entity.LogProduk,) error {
	return tx.WithContext(ctx).Create(log).Error
}

func (r *transactionImpl) CreateDetailTransaction(ctx context.Context,tx *gorm.DB,detail *entity.DetailTransaction,) error {
	return tx.WithContext(ctx).Create(detail).Error
}

func (r *transactionImpl) UpdateTotalHarga(ctx context.Context,tx *gorm.DB,trxID int,total float64,) error {
	return tx.WithContext(ctx).
		Model(&entity.Transaction{}).
		Where("id = ?", trxID).
		Update("harga_total", total).Error
}

func (r *transactionImpl) GetProdukByID(ctx context.Context,tx *gorm.DB,produkID int,) (*entity.ProdukWithOwner, error) {

	var res entity.ProdukWithOwner

	err := tx.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Table("produk").
		Select("produk.*, toko.id_user").
		Joins("JOIN toko ON toko.id = produk.id_toko").
		Where("produk.id = ?", produkID).
		First(&res).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *transactionImpl) UpdateStokProduk(ctx context.Context,tx *gorm.DB,produkID int,qty int,) error {
	return tx.WithContext(ctx).
		Model(&entity.Produk{}).
		Where("id = ?", produkID).
		Update("stok", gorm.Expr("stok - ?", qty)).
		Error
}

func (r *transactionImpl) GetTransactionsByUser(ctx context.Context,userID int,) ([]entity.Transaction, error) {

	var trxs []entity.Transaction

	err := r.db.WithContext(ctx).
		Preload("AlamatKirim").
		Where("id_user = ?", userID).
		Order("id DESC").
		Find(&trxs).Error

	return trxs, err
}

func (r *transactionImpl) GetTransactionDetails(ctx context.Context, trxID int) ([]entity.DetailTrxWithJoin, error) {
    var res []entity.DetailTrxWithJoin

    err := r.db.WithContext(ctx).
        Table("detail_trx").
        Select(`
            detail_trx.id_trx,
            detail_trx.kuantitas,
            detail_trx.harga_total,
            log_produk.id as id_log_produk,
            log_produk.nama_produk,
            log_produk.slug,
            log_produk.harga_reseller,
            log_produk.harga_konsumen,
            log_produk.deskripsi,
            log_produk.id_category,
            toko.id as id_toko,
            toko.nama_toko,
            toko.url_foto,
            category.nama_category
        `).
        Joins("JOIN log_produk ON log_produk.id = detail_trx.id_log_produk").
        Joins("JOIN toko ON toko.id = detail_trx.id_toko").
        Joins("LEFT JOIN category ON category.id = log_produk.id_category"). // Tambahkan join category
        Where("detail_trx.id_trx = ?", trxID).
        Find(&res).Error

    return res, err
}

func (r *transactionImpl) GetProductPhotosByLogProdukID(ctx context.Context, logProdukID int) ([]entity.ProductPhoto, error) {
    var photos []entity.ProductPhoto
    
    // Ambil id_produk dari log_produk
    var logProduk entity.LogProduk
    if err := r.db.WithContext(ctx).
        Select("id_produk").
        Where("id = ?", logProdukID).
        First(&logProduk).Error; err != nil {
        return nil, err
    }
    
    // Ambil photos berdasarkan id_produk
    err := r.db.WithContext(ctx).
        Where("id_produk = ?", logProduk.IDProduk).
        Find(&photos).Error
    
    return photos, err
}

func (r *transactionImpl) GetTransactionByID(ctx context.Context, trxID int, userID int) (*entity.Transaction, error) {
    var trx entity.Transaction
    
    err := r.db.WithContext(ctx).
        Preload("AlamatKirim").
        Where("id = ? AND id_user = ?", trxID, userID).
        First(&trx).Error
    
    if err != nil {
        return nil, err
    }
    
    return &trx, nil
}