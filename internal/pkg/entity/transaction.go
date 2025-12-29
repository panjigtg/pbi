package entity

import "time"

type (
	LogProduk struct {
		ID             int       `gorm:"primaryKey;autoIncrement"`
		IDProduk       int       `gorm:"column:id_produk;not null"`
		IDToko         int       `gorm:"column:id_toko;not null"`
		IDCategory     int       `gorm:"column:id_category;not null"`
		NamaProduk     string    `gorm:"column:nama_produk"`
		Slug           string    `gorm:"column:slug"`
		HargaReseller  float64   `gorm:"column:harga_reseller"`
		HargaKonsumen  float64   `gorm:"column:harga_konsumen"`
		Deskripsi      string    `gorm:"column:deskripsi"`
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
	Transaction struct {
		ID               int       `gorm:"primaryKey;autoIncrement"`
		IDUser           int       `gorm:"column:id_user;not null"`
		AlamatPengiriman int       `gorm:"column:alamat_pengiriman;not null"`
		HargaTotal       float64   `gorm:"column:harga_total"`
		KodeInvoice      string    `gorm:"column:kode_invoice"`
		MetodeBayar      string    `gorm:"column:metode_bayar"`
		CreatedAt        time.Time
		UpdatedAt        time.Time

		AlamatKirim *Destination `gorm:"foreignKey:AlamatPengiriman;references:ID" json:"alamat_kirim"`
	}

	DetailTransaction struct {
		ID          int       `gorm:"primaryKey;autoIncrement"`
		IDTrx       int       `gorm:"column:id_trx;not null"`
		IDLogProduk int       `gorm:"column:id_log_produk;not null"`
		IDToko      int       `gorm:"column:id_toko;not null"`
		Kuantitas   int       `gorm:"column:kuantitas"`
		HargaTotal  float64   `gorm:"column:harga_total"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	ProdukWithOwner struct {
		Produk
		IDUser int `gorm:"column:id_user"`
	}

)

func (Transaction) TableName() string {
	return "trx"
}

func (LogProduk) TableName() string {
	return "log_produk"
}

func (DetailTransaction) TableName() string {
	return "detail_trx"
}