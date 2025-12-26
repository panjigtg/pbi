package repository

import (
	"context"
	"database/sql"

	"pbi/internal/pkg/entity"
)

type TokoRepository interface {
	Create(ctx context.Context, toko *entity.Toko) (*entity.Toko, error)
	GetAll(ctx context.Context) ([]*entity.Toko, error)
	GetByUserID(ctx context.Context, userID int) ([]*entity.Toko, error)
	GetByID(ctx context.Context, tokoID int) (*entity.Toko, error)
	Update(ctx context.Context, tokoID int, data *entity.Toko) (*entity.Toko, error)
}

type tokoRepositoryImpl struct {
	db *sql.DB
}

func NewTokoRepository(db *sql.DB) TokoRepository {
	return &tokoRepositoryImpl{db: db}
}

func (r *tokoRepositoryImpl) Create(
	ctx context.Context,
	toko *entity.Toko,
) (*entity.Toko, error) {

	query := `
		INSERT INTO toko (id_user, nama_toko, url_foto)
		VALUES (?, ?, ?)
	`

	res, err := r.db.ExecContext(
		ctx,
		query,
		toko.IDUser,
		toko.NamaToko,
		toko.UrlFoto,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	toko.ID = int(id)
	return toko, nil
}

func (r *tokoRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Toko, error) {
	query := `
		SELECT id, id_user, nama_toko, url_foto
		FROM toko
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokos []*entity.Toko
	for rows.Next() {
		var t entity.Toko
		if err := rows.Scan(
			&t.ID,
			&t.IDUser,
			&t.NamaToko,
			&t.UrlFoto,
		); err != nil {
			return nil, err
		}
		tokos = append(tokos, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tokos, nil
}

func (r *tokoRepositoryImpl) GetByUserID(
	ctx context.Context,
	userID int,
) ([]*entity.Toko, error) {

	query := `
		SELECT id, id_user, nama_toko, url_foto
		FROM toko
		WHERE id_user = ?
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokos []*entity.Toko
	for rows.Next() {
		var t entity.Toko
		if err := rows.Scan(
			&t.ID,
			&t.IDUser,
			&t.NamaToko,
			&t.UrlFoto,
		); err != nil {
			return nil, err
		}
		tokos = append(tokos, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tokos, nil
}

func (r *tokoRepositoryImpl) GetByID(ctx context.Context, tokoID int) (*entity.Toko, error) {
    query := `SELECT id, id_user, nama_toko, url_foto, created_at, updated_at FROM toko WHERE id = ?`
    row := r.db.QueryRowContext(ctx, query, tokoID)

    var t entity.Toko
    if err := row.Scan(&t.ID, &t.IDUser, &t.NamaToko, &t.UrlFoto, &t.CreatedAt, &t.UpdatedAt); err != nil {
        return nil, err
    }

    return &t, nil
}


func (r *tokoRepositoryImpl) Update(ctx context.Context, tokoID int, data *entity.Toko) (*entity.Toko, error) {
	query := `UPDATE toko 
          SET nama_toko = COALESCE(?, nama_toko), 
              url_foto = COALESCE(?, url_foto) 
          WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, data.NamaToko, data.UrlFoto, tokoID)
	if err != nil {
		return nil, err
	}

	data.ID = tokoID
	return data, nil
}