package repository

import (
	"context"
	"database/sql"
	"errors"

	"pbi/internal/pkg/entity"
	"pbi/internal/pkg/models"
)

var ErrCategoryNotFound = errors.New("category not found")

type CategoryRepository interface {
	Create(ctx context.Context, category *entity.Category) error
	GetAll(ctx context.Context) ([]*entity.Category, error)
	GetById(ctx context.Context, id int) (*entity.Category, error)
	Update(ctx context.Context, id int, req *models.UpdateRequest)  error
	Delete(ctx context.Context, id int)  error
}

type categoryRepositoryImpl struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepositoryImpl{db: db}
}

func (r *categoryRepositoryImpl) Create(ctx context.Context, category *entity.Category) error {
    query := `INSERT INTO category (nama_category, created_at, updated_at) VALUES (?, NOW(), NOW())`
    
    res, err := r.db.ExecContext(ctx, query, category.Nama)
    if err != nil {
        return err 
    }

    id, err := res.LastInsertId()
    if err != nil {
        return err
    }
    category.ID = int(id)

    return nil
}


func (r *categoryRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Category, error) {
	query := `SELECT id, nama_category FROM category`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Nama); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (r *categoryRepositoryImpl) GetById(ctx context.Context, id int) (*entity.Category, error) {
	query := `SELECT id, nama_category FROM category WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id) // <-- pakai QueryRowContext

	var category entity.Category
	if err := row.Scan(&category.ID, &category.Nama); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepositoryImpl) Update(ctx context.Context, id int, req *models.UpdateRequest) error {
	query := 
		`
		UPDATE category SET
			nama_category = COALESCE(?, nama_category)
		WHERE id = ?
	`
	result, err := r.db.ExecContext(ctx, query, req.Nama, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrCategoryNotFound
	}

	return nil
}

func (r *categoryRepositoryImpl) Delete(ctx context.Context, id int)  error {
	query := `DELETE FROM category WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrCategoryNotFound
	}

	return nil
}