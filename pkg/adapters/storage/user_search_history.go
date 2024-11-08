package storage

import (
	"context"
	history "creepy/internal/user_search_history"
	"creepy/pkg/adapters/storage/entities"
	"creepy/pkg/adapters/storage/mappers"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userSearchHistoryRepo struct {
	db *gorm.DB
}

func NewUserSearchHistoryRepo(db *gorm.DB) *userSearchHistoryRepo {
	return &userSearchHistoryRepo{
		db: db,
	}
}

func (ur *userSearchHistoryRepo) Insert(ctx context.Context, h *history.UserSearchHistory) error {
	userSearchHistoryEntity := mappers.UserSearchHistoryDomainToEntity(h)
	if err := ur.db.WithContext(ctx).Save(&userSearchHistoryEntity).Error; err != nil {
		return err
	}
	h.ID = userSearchHistoryEntity.ID
	return nil
}

func (ur userSearchHistoryRepo) DeleteByID(ctx context.Context, id *uuid.UUID) error {
	userSearchHistory := entities.User{Model: entities.Model{ID: *id}}
    result := ur.db.WithContext(ctx).Delete(&userSearchHistory, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}

func (ur userSearchHistoryRepo) UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error {
    userSearchHistory := entities.User{Model: entities.Model{ID: *id}}
    result := ur.db.WithContext(ctx).Model(&userSearchHistory).Updates(updates)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}
