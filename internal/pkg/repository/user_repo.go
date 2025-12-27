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
	UpdateProfile(ctx context.Context, userID int64, nama *string, notelp *string,email *string, idProvinsi *string, idKota *string, tentang *string, pekerjaan *string, kataSandi *string,tanggalLahir *string, jenisKelamin *string, ) error
	FindById(ctx context.Context, userID int64) (*entity.User, error)
	GetProfile(ctx context.Context, userID int64) (*entity.User, error)
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

func (r *userRepositoryImpl) UpdateProfile(ctx context.Context, userID int64, nama *string, notelp *string, email *string, idProvinsi *string, idKota *string, tentang *string, pekerjaan *string, kataSandi *string,tanggalLahir *string, jenisKelamin *string,) error {

	query := `
		UPDATE user SET
			nama = COALESCE(?, nama),
			email = COALESCE(?, email),
			kata_sandi = COALESCE(?, kata_sandi),
			notelp = COALESCE(?, notelp),
			tanggal_lahir = COALESCE(?, tanggal_lahir),
			jenis_kelamin = COALESCE(?, jenis_kelamin),
			id_provinsi = COALESCE(?, id_provinsi),
			id_kota = COALESCE(?, id_kota),
			tentang = COALESCE(?, tentang),
			pekerjaan = COALESCE(?, pekerjaan),
			updated_at = NOW()
		WHERE id = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		nama,           // 1
		email,          // 2
		kataSandi,      // 3
		notelp,         // 4
		tanggalLahir,   // 5
		jenisKelamin,   // 6
		idProvinsi,     // 7
		idKota,         // 8
		tentang,        // 9
		pekerjaan,      // 10
		userID,         // 11 -> WHERE id=?
	)

	return err
}

func (r *userRepositoryImpl) FindById(ctx context.Context, userID int64,) (*entity.User, error) {
	u := &entity.User{}

	query := `
		SELECT id, nama, email, notelp, tanggal_lahir, jenis_kelamin,
		       tentang, pekerjaan, id_provinsi, id_kota, is_admin,
		       created_at, updated_at
		FROM user
		WHERE id = ?
		LIMIT 1
	`

	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&u.ID,
		&u.Nama,
		&u.Email,
		&u.NoTelp,
		&u.TanggalLahir,
		&u.JenisKelamin,
		&u.Tentang,
		&u.Pekerjaan,
		&u.IDProvinsi,
		&u.IDKota,
		&u.IsAdmin,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepositoryImpl) GetProfile(ctx context.Context,userID int64,) (*entity.User, error) {
	return r.FindById(ctx, userID)
}
