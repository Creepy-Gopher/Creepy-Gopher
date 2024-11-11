package postgis

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type propertyRepo struct {
    DB *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) storage.PropertyRepository {
    return &propertyRepo{DB: db}
}

func (r *propertyRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Property, error) {
    var property models.Property
    result := r.DB.WithContext(ctx).First(&property, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &property, nil
}

func (r *propertyRepo) Save(ctx context.Context, entity *models.Property) error {
    if err := r.DB.WithContext(ctx).Save(entity).Error; err != nil {
        return err
    }
    return nil
}

// Update attributes with `struct`, will only update non-zero fields
func (r *propertyRepo) Update(ctx context.Context, entity *models.Property) error {
    result := r.DB.WithContext(ctx).Model(&models.Property{}).Updates(entity)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", entity.ID)
    }
    return nil
}

func (r *propertyRepo) Delete(ctx context.Context, id uuid.UUID) error {
    property := models.Property{Model: models.Model{ID: id}}
    result := r.DB.WithContext(ctx).Delete(&property, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}

func (r *propertyRepo) ListProperties(ctx context.Context, filter *models.Filter) ([]*models.Property, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *propertyRepo) GetPropertyByURL(ctx context.Context, url string) (*models.Property, error) {
    return nil, fmt.Errorf("not implemented")
}
