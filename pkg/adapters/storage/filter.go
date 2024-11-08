package storage

import (
	"context"
	"creepy/internal/filter"
	"creepy/pkg/adapters/storage/entities"
	"creepy/pkg/adapters/storage/mappers"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilterRepo struct {
	db *gorm.DB
}

func NewFilterRepo(db *gorm.DB) filter.Repo {
	return &FilterRepo{db}
}

func (fr *FilterRepo) Insert(ctx context.Context, f *filter.Filter) error {
	filterEntity := mappers.FilterDomainToEntity(f)
	if err := fr.db.WithContext(ctx).Save(&filterEntity).Error; err != nil {
		return err
	}
	f.ID = filterEntity.ID
	return nil
}

func (fr FilterRepo) DeleteByID(ctx context.Context, id *uuid.UUID) error {
    result := fr.db.WithContext(ctx).Delete(&entities.Filter{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}

func (fr FilterRepo) UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error {
    filter := entities.Filter{Model: entities.Model{ID: *id}}
    result := fr.db.WithContext(ctx).Model(&filter).Updates(updates)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}
