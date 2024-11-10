package storage

import (
	"context"
	"creepy/internal/models"

	"github.com/google/uuid"
)

type PropertyRepository interface {
    SaveProperty(ctx context.Context, property *models.Property) error
	GetPropertyByID(ctx context.Context, id uuid.UUID) (*models.Property, error)
    UpdateProperty(ctx context.Context, property *models.Property) error
    DeleteProperty(ctx context.Context, id uuid.UUID) error
    ListProperties(ctx context.Context, filter *models.Filter) ([]*models.Property, error)
}

// Generic ?

// CRUD: Create, Read, Update, Delete
type Repository[T any] interface {
    GetByID(ctx context.Context, id uuid.UUID) (*T, error)
    Save(ctx context.Context, entity *T) error
    Update(ctx context.Context, entity *T) error
    Delete(ctx context.Context, id uuid.UUID) error
}

type FilterRepository interface {
    Repository[models.Property]
    // Add property-specific methods if needed
}
