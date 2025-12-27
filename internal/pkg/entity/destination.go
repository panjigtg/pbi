package entity

import "time"

type Destination struct {
	ID           int       `gorm:"primaryKey;column:id"`
	UserID       int       `gorm:"column:id_user;not null"`
	JudulAlamat  string    `gorm:"column:judul_alamat"`
	NamaPenerima string    `gorm:"column:nama_penerima"`
	NoTelp       string    `gorm:"column:no_telp"`
	DetailAlamat string    `gorm:"column:detail_alamat"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Destination) TableName() string {
	return "alamat"
}