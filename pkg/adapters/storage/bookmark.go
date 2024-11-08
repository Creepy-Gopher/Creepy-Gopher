package storage

import (
	"context"
	"creepy/internal/bookmark"
	"creepy/pkg/adapters/storage/entities"
	"creepy/pkg/adapters/storage/mappers"
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

func (br *bookmarkRepo) Insert(ctx context.Context, b *bookmark.Bookmark) error {
	bookmarkEntity := mappers.BookmarkDomainToEntity(b)
	if err := br.db.WithContext(ctx).Save(&bookmarkEntity).Error; err != nil {
		return err
	}
	b.ID = bookmarkEntity.ID
	return nil
}

func (br bookmarkRepo) DeleteByID(ctx context.Context, id *uuid.UUID) error {
	bookmark := entities.Bookmark{Model: entities.Model{ID: *id}}
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
    bookmark := entities.Bookmark{Model: entities.Model{ID: *id}}
    result := br.db.WithContext(ctx).Model(&bookmark).Updates(updates)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}
