package entity

import "time"

type Toko struct {
	ID        int       `gorm:"primaryKey"`
	IDUser    int       `gorm:"column:id_user"`
	NamaToko  string    `gorm:"column:nama_toko"`
	UrlFoto   string    `gorm:"column:url_foto"`
	CreatedAt time.Time
	UpdatedAt time.Time
}