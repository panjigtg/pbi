package repository

import (
	"context"
	"database/sql"

	"pbi/internal/pkg/entity"
)

type TokoRepository interface {
	CreateToko(ctx context.Context, toko *entity.Toko) (*entity.Toko, error)
}

type tokoRepositoryImpl struct {
	db *sql.DB
}

func NewTokoRepository(db *sql.DB) TokoRepository {
	return &tokoRepositoryImpl{db: db}
}

func (r *tokoRepositoryImpl) CreateToko(ctx context.Context, toko *entity.Toko) (*entity.Toko, error) {
	query := `INSERT INTO toko (id_user, nama_toko, url_foto) VALUES (?, ?, ?)`
	res, err := r.db.ExecContext(ctx, query, toko.IDUser, toko.NamaToko, toko.UrlFoto)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	toko.ID = int(id)
	return toko, nil
}