package bookmark

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"creepy/internal/property"
	"creepy/internal/user"
)

var (
	ErrSomeThing = errors.New("some error")
)

type Repo interface {
	Insert(ctx context.Context, b *Bookmark) error
	DeleteByID(ctx context.Context, id *uuid.UUID) error
	UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error
	// ...
}

type Bookmark struct {
	ID uuid.UUID

	UserID     uuid.UUID
	User       user.User // ?
	PropertyID uuid.UUID
	Property   property.Property // ?
}
