package repository

import (
	"context"
	"pbi/internal/pkg/entity"
	"database/sql"
)

type UserRepository interface {
	CheckEmailPhone(ctx context.Context, email string, notelp string) (*entity.User, error)
	CheckEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) CheckEmailPhone(ctx context.Context, email string, notelp string) (*entity.User, error) {
	u := &entity.User{}

	query := `SELECT id, nama, email, notelp, kata_sandi FROM user WHERE email = ? OR notelp = ? LIMIT 1`

	 err := r.db.QueryRowContext(ctx, query, email, notelp).Scan(
        &u.ID, &u.Nama, &u.Email, &u.NoTelp, &u.KataSandi,
    )

    if err == sql.ErrNoRows {
        return nil, nil
    } else if err != nil {
        return nil, err
    }
    return u, nil
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `
		INSERT INTO user 
			(nama, email, notelp, kata_sandi, tanggal_lahir, jenis_kelamin, tentang, pekerjaan, id_provinsi, id_kota, is_admin, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
	`
	res, err := r.db.ExecContext(ctx,
		query,
		user.Nama,
		user.Email,
		user.NoTelp,
		user.KataSandi,
		user.TanggalLahir,
		user.JenisKelamin,
		user.Tentang,
		user.Pekerjaan,
		user.IDProvinsi,
		user.IDKota,
		user.IsAdmin,
	)
	if err != nil {
		return nil, err
	}

	id, _ := res.LastInsertId()
	user.ID = int(id)
	return user, nil
}

func (r *userRepositoryImpl) CheckEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}

	query := `
	SELECT id, nama, email, notelp, kata_sandi, tanggal_lahir, jenis_kelamin, tentang, pekerjaan, id_provinsi, id_kota, is_admin
	FROM user
	WHERE email = ?
	LIMIT 1
	`

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&u.ID,
		&u.Nama,
		&u.Email,
		&u.NoTelp,
		&u.KataSandi,
		&u.TanggalLahir,
		&u.JenisKelamin,
		&u.Tentang,
		&u.Pekerjaan,
		&u.IDProvinsi,
		&u.IDKota,
		&u.IsAdmin,
	)

	if err == sql.ErrNoRows {
		return nil, nil 
	} else if err != nil {
		return nil, err
	}

	return u, nil
}