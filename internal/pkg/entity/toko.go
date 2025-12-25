package entity

import "time"

type Toko struct {
	ID       int
	IDUser   int
	NamaToko string
	UrlFoto  string
	CreatedAt time.Time
	UpdatedAt time.Time
}