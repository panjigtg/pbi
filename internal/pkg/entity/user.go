package entity

import "time"

type User struct {
	ID           int       `gorm:"primaryKey;column:id"`
	Nama         string    `gorm:"column:nama"`
	Email        string    `gorm:"column:email"`
	NoTelp       string    `gorm:"column:notelp"`
	KataSandi    string    `gorm:"column:kata_sandi"`
	IDProvinsi   string    `gorm:"column:id_provinsi"`
	IDKota       string    `gorm:"column:id_kota"`
	TanggalLahir string    `gorm:"column:tanggal_lahir"`
	JenisKelamin string    `gorm:"column:jenis_kelamin"`
	Tentang      string    `gorm:"column:tentang"`
	Pekerjaan    string    `gorm:"column:pekerjaan"`
	IsAdmin      bool      `gorm:"column:is_admin"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
