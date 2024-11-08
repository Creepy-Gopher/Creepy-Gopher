package watchlist

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"creepy/internal/filter"
	"creepy/internal/user"
)

var (
	ErrSomeThing = errors.New("some error")
)

type Repo interface {
	Insert(ctx context.Context, w *WatchList) error
	DeleteByID(ctx context.Context, id *uuid.UUID) error
	UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error
	// ...
}

type WatchList struct {
	ID uuid.UUID

	UserID     uuid.UUID
	User       user.User
	FilterID   uuid.UUID
	Filter     filter.Filter
	FilterName string
}
