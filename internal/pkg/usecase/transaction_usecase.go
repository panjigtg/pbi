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

type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, userID int, req *models.CreateTrxRequest) (int, *helper.ErrorStruct)
    GetAll(ctx context.Context, userID int) (*models.TransactionListResponseWrapper, *helper.ErrorStruct)
	GetByID(ctx context.Context, trxID int, userID int) (*models.TransactionDetailByIDResponse, *helper.ErrorStruct) // Tambahkan ini
}

type transactionImpl struct {
	db        	*gorm.DB
	repo 		repository.TransactionRepository
	destRepo repository.DestinationRepository
}

func NewTransactionUsecase(db *gorm.DB,trxRepo repository.TransactionRepository,destRepo repository.DestinationRepository,) TransactionUsecase {
	return &transactionImpl{
		db:       db,
		repo:  trxRepo,
		destRepo: destRepo,
	}
}

func (t *transactionImpl) CreateTransaction(ctx context.Context, userID int, req *models.CreateTrxRequest) (int, *helper.ErrorStruct) {
	if err := helper.Validate.Struct(req); err != nil {
		return 0, &helper.ErrorStruct{Err: err, Code: 400}
	}

	if _, err := t.destRepo.FindByID(ctx, req.AlamatKirim, userID); err != nil {
		return 0, &helper.ErrorStruct{
			Err:  errors.New("alamat pengiriman tidak valid"),
			Code: 404,
		}
	}

	var trxID int

	err := t.db.Transaction(func(tx *gorm.DB) error {

		trx := &entity.Transaction{
			IDUser:           userID,
			AlamatPengiriman: req.AlamatKirim,
			MetodeBayar:      req.MethodBayar,
			KodeInvoice:      helper.GenerateInvoice(),
		}

		if err := t.repo.CreateTransaction(ctx, tx, trx); err != nil {
			return err
		}

		totalHarga := float64(0)

		for _, item := range req.DetailTrx {

			produk, err := t.repo.GetProdukByID(ctx, tx, item.ProductID)
			if err != nil {
				return err
			}

			if produk.IDUser == userID {
				return errors.New("tidak boleh membeli produk milik sendiri")
			}

			if produk.Stok < item.Kuantitas {
				return errors.New("stok produk tidak mencukupi")
			}

			if err := t.repo.UpdateStokProduk(ctx, tx, produk.ID, item.Kuantitas); err != nil {
				return err
			}

			logProduk := &entity.LogProduk{
				IDProduk:      produk.ID,
				IDToko:        produk.IDToko,
				IDCategory:    produk.IDCategory,
				NamaProduk:    produk.NamaProduk,
				Slug:          produk.Slug,
				HargaReseller: produk.HargaReseller,
				HargaKonsumen: produk.HargaKonsumen,
				Deskripsi:     produk.Deskripsi,
			}

			if err := t.repo.CreateLogProduk(ctx, tx, logProduk); err != nil {
				return err
			}

			subTotal := produk.HargaKonsumen * float64(item.Kuantitas)
			totalHarga += subTotal

			detail := &entity.DetailTransaction{
				IDTrx: 			trx.ID,
				IDLogProduk:   	logProduk.ID,
				IDToko:        	produk.IDToko,
				Kuantitas:     	item.Kuantitas,
				HargaTotal:   	subTotal,
			}

			if err := t.repo.CreateDetailTransaction(ctx, tx, detail); err != nil {
				return err
			}
		}

		if err := t.repo.UpdateTotalHarga(ctx, tx, trx.ID, totalHarga); err != nil {
			return err
		}

		trxID = trx.ID
		return nil
	})

	if err != nil {
		return 0, &helper.ErrorStruct{Err: err, Code: 500}
	}

	return trxID, nil
}

func (u *transactionImpl) GetAll(ctx context.Context, userID int) (*models.TransactionListResponseWrapper, *helper.ErrorStruct) {

    trxs, err := u.repo.GetTransactionsByUser(ctx, userID)
    if err != nil {
        return nil, &helper.ErrorStruct{Err: err, Code: 500}
    }

    var res []models.TransactionListResponse

    for _, trx := range trxs {
        details, err := u.repo.GetTransactionDetails(ctx, trx.ID)
        if err != nil {
            return nil, &helper.ErrorStruct{Err: err, Code: 500}
        }

        var detailRes []models.TransactionDetailResponse
        for _, d := range details {
            photos, _ := u.repo.GetProductPhotosByLogProdukID(ctx, d.IDLogProduk)
            
            var photoRes []models.PhotoInfo
            for _, p := range photos {
                photoRes = append(photoRes, models.PhotoInfo{
                    ID:        p.ID,
                    ProductID: p.ProductID,
                    URL:       p.Url,
                })
            }

            item := models.TransactionDetailResponse{
                Product: models.ProductDetailInTransaction{
                    ID:            d.IDLogProduk,
                    NamaProduk:    d.NamaProduk,
                    Slug:          d.Slug,
                    HargaReseller: d.HargaReseller,
                    HargaKonsumen: d.HargaKonsumen,
                    Deskripsi:     d.Deskripsi,
                    Toko: models.TokoBasicInfo{
                        NamaToko: d.NamaToko, 
                        URLFoto:  d.URLFoto,  
                    },
                    Category: models.CategoryInfo{
                        ID:           d.IDCategory,
                        NamaCategory: d.NamaCategory,
                    },
                    Photos: photoRes,
                },
                Toko: models.TokoInfo{
                    ID:       d.IDToko,
                    NamaToko: d.NamaToko,
                    URLFoto:  d.URLFoto,
                },
                Kuantitas:  d.Kuantitas,
                HargaTotal: d.HargaTotal,
            }

            detailRes = append(detailRes, item)
        }

        trxResp := models.TransactionListResponse{
            ID:          trx.ID,
            HargaTotal:  trx.HargaTotal,
            KodeInvoice: trx.KodeInvoice,
            MethodBayar: trx.MetodeBayar,
            DetailTrx:   detailRes,
        }

        if trx.AlamatKirim != nil {
            trxResp.AlamatKirim = models.DestinationResponse{
                ID:           trx.AlamatKirim.ID,
                JudulAlamat:  trx.AlamatKirim.JudulAlamat,
                NamaPenerima: trx.AlamatKirim.NamaPenerima,
                NoTelp:       trx.AlamatKirim.NoTelp,
                DetailAlamat: trx.AlamatKirim.DetailAlamat,
            }
        }

        res = append(res, trxResp)
    }

    return &models.TransactionListResponseWrapper{
        Data:  res,
        Page:  0,
        Limit: 0,
    }, nil
}

func (u *transactionImpl) GetByID(ctx context.Context, trxID int, userID int) (*models.TransactionDetailByIDResponse, *helper.ErrorStruct) {
    
    // Get transaction
    trx, err := u.repo.GetTransactionByID(ctx, trxID, userID)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, &helper.ErrorStruct{
                Err:  errors.New("transaksi tidak ditemukan"),
                Code: 404,
            }
        }
        return nil, &helper.ErrorStruct{Err: err, Code: 500}
    }
    
    // Get transaction details
    details, err := u.repo.GetTransactionDetails(ctx, trx.ID)
    if err != nil {
        return nil, &helper.ErrorStruct{Err: err, Code: 500}
    }
    
    var detailRes []models.TransactionDetailResponse
    for _, d := range details {
        photos, _ := u.repo.GetProductPhotosByLogProdukID(ctx, d.IDLogProduk)
        
        var photoRes []models.PhotoInfo
        for _, p := range photos {
            photoRes = append(photoRes, models.PhotoInfo{
                ID:        p.ID,
                ProductID: p.ProductID,
                URL:       p.Url,
            })
        }
        
        item := models.TransactionDetailResponse{
            Product: models.ProductDetailInTransaction{
                ID:            d.IDLogProduk,
                NamaProduk:    d.NamaProduk,
                Slug:          d.Slug,
                HargaReseller: d.HargaReseller,
                HargaKonsumen: d.HargaKonsumen,
                Deskripsi:     d.Deskripsi,
                Toko: models.TokoBasicInfo{
                    NamaToko: d.NamaToko,
                    URLFoto:  d.URLFoto,
                },
                Category: models.CategoryInfo{
                    ID:           d.IDCategory,
                    NamaCategory: d.NamaCategory,
                },
                Photos: photoRes,
            },
            Toko: models.TokoInfo{
                ID:       d.IDToko,
                NamaToko: d.NamaToko,
                URLFoto:  d.URLFoto,
            },
            Kuantitas:  d.Kuantitas,
            HargaTotal: d.HargaTotal,
        }
        
        detailRes = append(detailRes, item)
    }
    
    response := &models.TransactionDetailByIDResponse{
        ID:          trx.ID,
        HargaTotal:  trx.HargaTotal,
        KodeInvoice: trx.KodeInvoice,
        MethodBayar: trx.MetodeBayar,
        DetailTrx:   detailRes,
    }
    
    if trx.AlamatKirim != nil {
        response.AlamatKirim = models.DestinationResponse{
            ID:           trx.AlamatKirim.ID,
            JudulAlamat:  trx.AlamatKirim.JudulAlamat,
            NamaPenerima: trx.AlamatKirim.NamaPenerima,
            NoTelp:       trx.AlamatKirim.NoTelp,
            DetailAlamat: trx.AlamatKirim.DetailAlamat,
        }
    }
    
    return response, nil
}