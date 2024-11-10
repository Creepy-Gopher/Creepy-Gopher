package postgis


import (
	"context"
	"creepy/internal/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type bookmarkRepo struct {
	db *gorm.DB
}

func NewBookmarkRepo(db *gorm.DB) *bookmarkRepo {
	return &bookmarkRepo{
		db: db,
	}
}

func (br *bookmarkRepo) Insert(ctx context.Context, bookmark *models.Bookmark) error {
	if err := br.db.WithContext(ctx).Save(bookmark).Error; err != nil {
		return err
	}
	return nil
}

func (br bookmarkRepo) DeleteByID(ctx context.Context, id *uuid.UUID) error {
	bookmark := models.Bookmark{Model: models.Model{ID: *id}}
    result := br.db.WithContext(ctx).Delete(&bookmark, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}

func (br bookmarkRepo) UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error {
    bookmark := models.Bookmark{Model: models.Model{ID: *id}}
    result := br.db.WithContext(ctx).Model(&bookmark).Updates(updates)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}
