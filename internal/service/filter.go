package service

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"

	"github.com/google/uuid"
)

type FilterService struct {
    Repo storage.FilterRepository
}

func NewFilterService(repo storage.FilterRepository) *FilterService {
    return &FilterService{Repo: repo}
}

func (s *FilterService) CreateFilter(ctx context.Context, filter *models.Filter) error {
    return s.Repo.Save(ctx, filter)
}

func (s *FilterService) GetFilter(ctx context.Context, id uuid.UUID) (*models.Filter, error) {
    return s.Repo.GetByID(ctx, id)
}

func (s *FilterService) UpdateFilter(ctx context.Context, filter *models.Filter) error {
	return s.Repo.Update(ctx, filter)
}

func (s *FilterService) DeleteFilter(ctx context.Context, id uuid.UUID) error {
	return s.Repo.Delete(ctx, id)
}

func (s *FilterService) GetByFilter(ctx context.Context, filter *models.Filter) (*models.Filter, error) {
	return s.Repo.GetByFilter(ctx, filter)
}
