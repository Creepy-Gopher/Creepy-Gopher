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

func NewBookmarkRepository(db *gorm.DB) storage.BookmarkRepository {
	return &bookmarkRepo{
		DB: db,
	}
}

func (r *bookmarkRepo) CreateBookmark(ctx context.Context, entity *models.Bookmark) error {
	if err := r.DB.WithContext(ctx).Save(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookmarkRepo) DeleteBookmard(ctx context.Context, id uuid.UUID) error {
	bookmark := models.Bookmark{Model: models.Model{ID: id}}
	result := r.DB.WithContext(ctx).Delete(&bookmark, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no bookmark found with ID: %v", id)
	}
	return nil
}

func (r *bookmarkRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Bookmark, error) {
	var bookmark models.Bookmark
	result := r.DB.WithContext(ctx).First(&bookmark, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bookmark, nil
}
