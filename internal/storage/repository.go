package storage

import (
	"context"
	"creepy/internal/models"

	"github.com/google/uuid"
)


// CRUD: Create, Read, Update, Delete
type Repository[T any] interface {
    GetByID(ctx context.Context, id uuid.UUID) (*T, error)
    Save(ctx context.Context, entity *T) error
    Update(ctx context.Context, entity *T) error
    Delete(ctx context.Context, id uuid.UUID) error
}

type PropertyRepository interface {
    Repository[models.Property]
    ListProperties(ctx context.Context, filter *models.Filter) ([]*models.Property, error)
    GetPropertyByURL(ctx context.Context, url string) (*models.Property, error)
}

type FilterRepository interface {
    Repository[models.Filter]
    GetByFilter(ctx context.Context, filter *models.Filter) (*models.Filter, error)
}

type UserRepository interface {
    Repository[models.User]
    GetByUserName(ctx context.Context, userName string) (*models.User, error)
}

type BookmarkRepository interface {
    Repository[models.Bookmark]
    // TODO: Add specific methods if needed
}

type UserSearchHistoryRepository interface {
    Repository[models.UserSearchHistory]
    ListSearchHistoryByUserName(ctx context.Context, userName string) ([]*models.UserSearchHistory, error)
}

type WatchListRepository interface {
    Repository[models.WatchList]
    // TODO: Add specific methods if needed
}
