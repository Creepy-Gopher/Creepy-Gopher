package postgis

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userSearchHistoryRepo struct {
	DB *gorm.DB
}

func NewUserSearchHistoryRepo(db *gorm.DB) storage.UserSearchHistoryRepository {
	return &userSearchHistoryRepo{
		DB: db,
	}
}

func (r *userSearchHistoryRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.UserSearchHistory, error) {
    var userSearchHistory models.UserSearchHistory
    result := r.DB.WithContext(ctx).First(&userSearchHistory, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &userSearchHistory, nil
}

func (r *userSearchHistoryRepo) Save(ctx context.Context, entity *models.UserSearchHistory) error {
    	if err := r.DB.WithContext(ctx).Save(entity).Error; err != nil {
            return err
        }
        return nil
}

func (r *userSearchHistoryRepo) Update(ctx context.Context, entity *models.UserSearchHistory) error {
    result := r.DB.WithContext(ctx).Model(&models.UserSearchHistory{}).Updates(entity)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", entity.ID)
    }
    return nil
}

func (r *userSearchHistoryRepo) Delete(ctx context.Context, id uuid.UUID) error {
    userSearchHistory := models.UserSearchHistory{Model: models.Model{ID: id}}
    result := r.DB.WithContext(ctx).Delete(&userSearchHistory, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}


func (r *userSearchHistoryRepo) ListSearchHistoryByUserName(ctx context.Context, userName string) ([]*models.UserSearchHistory, error) {
    var searchHistory []*models.UserSearchHistory
    result := r.DB.WithContext(ctx).Where("user_name = ?", userName).Find(searchHistory)
    if result.Error != nil {
        return nil, result.Error
    }
    if result.RowsAffected == 0 {
        return nil, fmt.Errorf("no record found with user_name %v", userName)
    }
    return searchHistory, nil
}