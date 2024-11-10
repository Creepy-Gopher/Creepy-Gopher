package mysql

import (
	"context"
	"creepy/internal/models"
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

func (ur *userSearchHistoryRepo) Insert(ctx context.Context, userSearchHistory *models.UserSearchHistory) error {
	if err := ur.db.WithContext(ctx).Save(userSearchHistory).Error; err != nil {
		return err
	}
	return nil
}

func (ur userSearchHistoryRepo) DeleteByID(ctx context.Context, id *uuid.UUID) error {
	userSearchHistory := models.User{Model: models.Model{ID: *id}}
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
    userSearchHistory := models.User{Model: models.Model{ID: *id}}
    result := ur.db.WithContext(ctx).Model(&userSearchHistory).Updates(updates)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}
