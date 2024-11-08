package storage

import (
	"context"
	"creepy/internal/watchlist"
	"creepy/pkg/adapters/storage/entities"
	"creepy/pkg/adapters/storage/mappers"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type watchlistRepo struct {
	db *gorm.DB
}

func NewWatchListRepo(db *gorm.DB) *watchlistRepo {
	return &watchlistRepo{
		db: db,
	}
}

func (wr *watchlistRepo) Insert(ctx context.Context, w *watchlist.WatchList) error {
	watchlistEntity := mappers.WatchListDomainToEntity(w)
	if err := wr.db.WithContext(ctx).Save(&watchlistEntity).Error; err != nil {
		return err
	}
	w.ID = watchlistEntity.ID
	return nil
}

func (wr *watchlistRepo) DeleteByID(ctx context.Context, id *uuid.UUID) error {
	watchlist := entities.WatchList{Model: entities.Model{ID: *id}}
    result := wr.db.WithContext(ctx).Delete(&watchlist, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}

func (wr *watchlistRepo) UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error {
    watchlist := entities.WatchList{Model: entities.Model{ID: *id}}
    result := wr.db.WithContext(ctx).Model(&watchlist).Updates(updates)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}
