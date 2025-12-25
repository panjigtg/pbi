package entity

import "time"

type User struct {
	ID           int
	Nama         string
	Email        string
	NoTelp       string
	KataSandi    string
	IDProvinsi   string
	IDKota       string
	TanggalLahir string
	JenisKelamin string
	Tentang      string
	Pekerjaan    string
	IsAdmin      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
