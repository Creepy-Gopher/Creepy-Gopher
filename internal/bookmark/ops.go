package bookmark

import (
	"context"
)

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo}
}

func (o *Ops) CreateBookmark(ctx context.Context, b *Bookmark) error {
	return nil
}
