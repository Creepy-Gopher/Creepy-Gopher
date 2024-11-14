package postgis

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type filterRepo struct {
	DB *gorm.DB
}

func NewFilterRepo(db *gorm.DB) storage.FilterRepository {
	return &filterRepo{
		DB: db,
	}
}

func (r *filterRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Filter, error) {
	var filter models.Filter
	result := r.DB.WithContext(ctx).First(&filter, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &filter, nil
}

func (r *filterRepo) Save(ctx context.Context, entity *models.Filter) error {
	if err := r.DB.WithContext(ctx).Save(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *filterRepo) Update(ctx context.Context, entity *models.Filter) error {
	result := r.DB.WithContext(ctx).Model(&models.Filter{}).Updates(entity)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with ID %v", entity.ID)
	}
	return nil
}

func (r *filterRepo) Delete(ctx context.Context, id uuid.UUID) error {
	filter := models.Filter{Model: models.Model{ID: id}}
	result := r.DB.WithContext(ctx).Delete(&filter, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with ID %v", id)
	}
	return nil
}

func (r *filterRepo) GetByFilter(ctx context.Context, filter *models.Filter) (*models.Filter, error) {
	var resFilter models.Filter
	result := r.DB.WithContext(ctx).Where(filter).First(&resFilter)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no record found with filter %v", *filter)
	}
	return &resFilter, nil
}

func (r *filterRepo) GetById(ctx context.Context, id uuid.UUID) (*models.Filter, error) {
	var filter models.Filter
	result := r.DB.WithContext(ctx).Where("id = ?", id).First(&filter)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no record found with id %v", id)
	}
	return &filter, nil
}
