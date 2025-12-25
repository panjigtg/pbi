package db

import (
	"context"
	"database/sql"
	"errors"

	"pbi/internal/pkg/entity"
	"pbi/internal/utils"
)

func SeedAdmin(db *sql.DB) error {
	ctx := context.Background()

	provinceID := "31" 
	cityID := "3171"    

	ok, err := utils.ValidateProvince(provinceID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("invalid province id")
	}

	ok, err = utils.ValidateCity(provinceID, cityID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("invalid city id")
	}

	hashedPassword, err := utils.HashPassword("admin123")
	if err != nil {
		return err
	}

	admin := &entity.User{
		Nama:         "Super Admin",
		Email:        "admin@mail.com",
		NoTelp:       "08123456783",
		KataSandi:    hashedPassword,
		IDProvinsi:   provinceID,
		IDKota:       cityID,
		IsAdmin:      true,
		TanggalLahir: "2000-01-01",
		JenisKelamin: "Laki-laki",
		Tentang:      "System Administrator",
		Pekerjaan:    "Admin",
	}

	query := `
	INSERT INTO user (
		nama, email, notelp, kata_sandi,
		id_provinsi, id_kota,
		is_admin, tanggal_lahir,
		jenis_kelamin, tentang, pekerjaan
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE email=email
	`

	_, err = db.ExecContext(
		ctx,
		query,
		admin.Nama,
		admin.Email,
		admin.NoTelp,
		admin.KataSandi,
		admin.IDProvinsi,
		admin.IDKota,
		admin.IsAdmin,
		admin.TanggalLahir,
		admin.JenisKelamin,
		admin.Tentang,
		admin.Pekerjaan,
	)

	return err
}