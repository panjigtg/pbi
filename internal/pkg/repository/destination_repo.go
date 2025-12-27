package repository

import (
	"context"
	"pbi/internal/pkg/entity"

	"gorm.io/gorm"
)

type DestinationRepository interface {
	Create(ctx context.Context,req *entity.Destination) error
	Update(ctx context.Context, req *entity.Destination) error
	Delete(ctx context.Context, id int, userID int) error
	FindByID(ctx context.Context, id int, userID int) (*entity.Destination, error)
	FindAllByUserID(ctx context.Context, userID int) ([]entity.Destination, error)
}

type destinationImpl struct {
	orm *gorm.DB
}

func NewDestinationRepo(orm *gorm.DB) DestinationRepository {
	return &destinationImpl{
		orm: orm,
	}
}

func (d *destinationImpl) Create(ctx context.Context,req *entity.Destination) error {
    return d.orm.WithContext(ctx).Create(req).Error
}

func (d *destinationImpl) Update(ctx context.Context, req *entity.Destination) error {
	return d.orm.WithContext(ctx).
		Model(&entity.Destination{}).
		Where("id = ? AND id_user = ?", req.ID, req.UserID).
		Updates(req).Error
}

func (d *destinationImpl) Delete(ctx context.Context, id int, userID int) error {
	return d.orm.WithContext(ctx).
		Where("id = ? AND id_user = ?", id, userID).
		Delete(&entity.Destination{}).Error
}

func (d *destinationImpl) FindByID(ctx context.Context, id int, userID int) (*entity.Destination, error) {
	var dest entity.Destination

	err := d.orm.WithContext(ctx).
		Where("id = ? AND id_user = ?", id, userID).
		First(&dest).Error

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (d *destinationImpl) FindAllByUserID(ctx context.Context,userID int,) ([]entity.Destination, error) {

	var dests []entity.Destination

	err := d.orm.WithContext(ctx).
		Where("id_user = ?", userID).
		Order("id DESC").
		Find(&dests).
		Error

	if err != nil {
		return nil, err
	}

	return dests, nil
}
