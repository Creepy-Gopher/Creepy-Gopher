package usersearchhistory

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"creepy/internal/filter"
	"creepy/internal/user"
)

var (
	ErrSomeThing = errors.New("some error")
)

type Repo interface {
	Insert(ctx context.Context, h *UserSearchHistory) error
	DeleteByID(ctx context.Context, id *uuid.UUID) error
	UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error
	// ...
}

type UserSearchHistory struct {
	ID       uuid.UUID
	UserName string
	User     user.User
	FilterID uuid.UUID
	Filter   filter.Filter
	Date     time.Time
}
