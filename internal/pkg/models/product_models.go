package models

import (
	"mime/multipart"
)

type (
	ProductCreateReq struct {
		IDCategory    int                   `form:"id_category" validate:"required"`
		NamaProduk    string                `form:"nama_produk" validate:"required"`
		Slug          string                `form:"slug"`
		HargaReseller float64               `form:"harga_reseller" validate:"required"`
		HargaKonsumen float64               `form:"harga_konsumen" validate:"required"`
		Stok          int                   `form:"stok" validate:"required"`
		Deskripsi     string                `form:"deskripsi"`
		Photos        []*multipart.FileHeader `form:"photos"`
	}

	ProductGetAllReq struct {
		NamaProduk string  `json:"nama_produk"`
		CategoryID int     `json:"category_id"`
		TokoID     int     `json:"toko_id"`
		MinHarga   float64 `json:"min_harga"`
		MaxHarga   float64 `json:"max_harga"`
		Page       int     `json:"page"`
		Limit      int     `json:"limit"`
	}

	ProductResponse struct {
		ID            int                    `json:"id"`
		NamaProduk    string                 `json:"nama_produk"`
		Slug          string                 `json:"slug"`
		HargaReseller float64                `json:"harga_reseler"`
		HargaKonsumen float64                `json:"harga_konsumen"`
		Stok          int                    `json:"stok"`
		Deskripsi     string                 `json:"deskripsi"`
		Toko          *TokoResponse          `json:"toko"`
		Category      *CategoryResponse      `json:"category"`
		Photos        []ProductPhotoResponse `json:"photos"`
	}

	ProductPhotoResponse struct {
		ID        int    `json:"id"`
		ProductID int    `json:"product_id"`
		Url       string `json:"url"`
	}

	ProductListResponse struct {
		Data   	[]ProductResponse 	`json:"data"`
		Total      int64           	`json:"total"`
		Page       int             	`json:"page"`
		Limit      int             	`json:"limit"`
		TotalPages int             	`json:"total_pages"`
	}
)