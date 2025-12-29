package entity

import "time"

type Produk struct {
	ID            int       `gorm:"column:id;primaryKey;autoIncrement"`
	IDToko        int       `gorm:"column:id_toko;not null"`
	IDCategory    int       `gorm:"column:id_category;not null"`
	NamaProduk    string    `gorm:"column:nama_produk;type:varchar(255)"`
	Slug          string    `gorm:"column:slug;type:varchar(255)"`
	HargaReseller float64   `gorm:"column:harga_reseller;type:decimal(15,2);not null"`
	HargaKonsumen float64   `gorm:"column:harga_konsumen;type:decimal(15,2);not null"`
	Stok          int       `gorm:"column:stok"`
	Deskripsi     string    `gorm:"column:deskripsi;type:text"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime"`


	Toko     *Toko          `gorm:"foreignKey:IDToko;references:ID"`
	Category *Category      `gorm:"foreignKey:IDCategory;references:ID"`
	Photos   []ProductPhoto `gorm:"foreignKey:ProductID;references:ID"`
}

type ProductPhoto struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement"`
	ProductID int    `gorm:"column:id_produk;not null"`
	Url       string `gorm:"column:url;type:varchar(255)" json:"url"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type ProductFilter struct {
	NamaProduk string
	CategoryID int
	TokoID     int
	MinHarga   float64
	MaxHarga   float64
	Page       int
	Limit      int
}

func (Produk) TableName() string { 
	return "produk" 
}
func (ProductPhoto) TableName() string { 
	return "foto_produk" 
}
func (Toko) TableName() string {
	return "toko"
}

func (Category) TableName() string {
	return "category"
}