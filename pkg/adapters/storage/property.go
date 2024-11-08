package storage

import (
	"context"
	"creepy/internal/property"
	"creepy/pkg/adapters/storage/entities"
	"creepy/pkg/adapters/storage/mappers"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type propertyRepo struct{
	db *gorm.DB
}

func NewPropertyRepo(db *gorm.DB) *propertyRepo {
	return &propertyRepo{
		db: db,
	}
}

func (pr *propertyRepo) Insert(ctx context.Context, p *property.Property) error {
	propertyEntity := mappers.PropertyDomainToEntity(p)
	if err := pr.db.WithContext(ctx).Save(&propertyEntity).Error; err != nil {
		return err
	}
	p.ID = propertyEntity.ID
	return nil
}

func (pr *propertyRepo) DeleteByID(ctx context.Context, id *uuid.UUID) error {
	property := entities.Property{Model: entities.Model{ID: *id}}
    result := pr.db.WithContext(ctx).Delete(&property, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}

func (pr *propertyRepo) UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error {
    property := entities.Property{Model: entities.Model{ID: *id}}
    result := pr.db.WithContext(ctx).Model(&property).Updates(updates)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}
