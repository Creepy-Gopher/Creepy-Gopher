package mysql

import (
	"context"
	"creepy/internal/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type FilterRepo struct {
	db *gorm.DB
}

func NewFilterRepo(db *gorm.DB) *FilterRepo {
	return &FilterRepo{db}
}

func (fr *FilterRepo) Insert(ctx context.Context, filter *models.Filter) error {
	if err := fr.db.WithContext(ctx).Save(filter).Error; err != nil {
		return err
	}
	return nil
}

func (fr FilterRepo) DeleteByID(ctx context.Context, id *uuid.UUID) error {
    result := fr.db.WithContext(ctx).Delete(&models.Filter{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}

func (fr FilterRepo) UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error {
    filter := models.Filter{Model: models.Model{ID: *id}}
    result := fr.db.WithContext(ctx).Model(&filter).Updates(updates)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}
