package entity

import "time"

type Toko struct {
	ID       int		`gorm:"primaryKey"`
	IDUser   int
	NamaToko string
	UrlFoto  string
	CreatedAt time.Time
	UpdatedAt time.Time
}