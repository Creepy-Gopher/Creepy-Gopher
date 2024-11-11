package postgis


import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type bookmarkRepo struct {
	DB *gorm.DB
}

func NewBookmarkRepo(db *gorm.DB) storage.BookmarkRepository {
	return &bookmarkRepo{
		DB: db,
	}
}

func (r *bookmarkRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Bookmark, error) {
    var bookmark models.Bookmark
    result := r.DB.WithContext(ctx).First(&bookmark, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &bookmark, nil
}

func (r *bookmarkRepo) Save(ctx context.Context, entity *models.Bookmark) error {
    	if err := r.DB.WithContext(ctx).Save(entity).Error; err != nil {
            return err
        }
        return nil
}

func (r *bookmarkRepo) Update(ctx context.Context, entity *models.Bookmark) error {
    result := r.DB.WithContext(ctx).Model(&models.Bookmark{}).Updates(entity)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", entity.ID)
    }
    return nil
}

func (r *bookmarkRepo) Delete(ctx context.Context, id uuid.UUID) error {
    bookmark := models.Bookmark{Model: models.Model{ID: id}}
    result := r.DB.WithContext(ctx).Delete(&bookmark, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}
