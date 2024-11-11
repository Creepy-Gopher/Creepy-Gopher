package postgis

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type watchListRepo struct {
	DB *gorm.DB
}

func NewWatchListRepo(db *gorm.DB) storage.WatchListRepository {
	return &watchListRepo{
		DB: db,
	}
}

func (r *watchListRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.WatchList, error) {
    var watchList models.WatchList
    result := r.DB.WithContext(ctx).First(&watchList, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &watchList, nil
}

func (r *watchListRepo) Save(ctx context.Context, entity *models.WatchList) error {
    	if err := r.DB.WithContext(ctx).Save(entity).Error; err != nil {
            return err
        }
        return nil
}

func (r *watchListRepo) Update(ctx context.Context, entity *models.WatchList) error {
    result := r.DB.WithContext(ctx).Model(&models.WatchList{}).Updates(entity)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", entity.ID)
    }
    return nil
}

func (r *watchListRepo) Delete(ctx context.Context, id uuid.UUID) error {
    watchList := models.WatchList{Model: models.Model{ID: id}}
    result := r.DB.WithContext(ctx).Delete(&watchList, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}
