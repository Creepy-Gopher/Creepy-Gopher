package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrSomeThing = errors.New("some error")
)

type Repo interface {
	Insert(ctx context.Context, u *User) error
	DeleteByID(ctx context.Context, id *uuid.UUID) error
	UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error
	// ...
}

type User struct {
	ID uuid.UUID

	UserName  string
	Role      string
	IsPremium bool
}
