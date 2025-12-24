package entity

import "time"

type User struct {
	ID           	int			`gorm:"primaryKey"`
	Nama         	string
	Email        	string
	NoTelp       	string
	KataSandi     	string		`gorm:"column:kata_sandi"`
	IDProvinsi string   `gorm:"column:id_provinsi"`
	IDKota     string   `gorm:"column:id_kota"`
	IsAdmin      	bool
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}