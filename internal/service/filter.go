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
	// TODO: Error handling
    return &FilterService{Repo: repo}
}

func (s *FilterService) CreateFilter(ctx context.Context, filter *models.Filter) error {
	// TODO: Error handling
    return s.Repo.Save(ctx, filter)
}

func (s *FilterService) GetFilter(ctx context.Context, id uuid.UUID) (*models.Filter, error) {
    // TODO: Error handling
    return s.Repo.GetByID(ctx, id)
}

func (s *FilterService) UpdateFilter(ctx context.Context, filter *models.Filter) error {
	// TODO: Error handling
	return s.Repo.Update(ctx, filter)
}

func (s *FilterService) DeleteFilter(ctx context.Context, id uuid.UUID) error {
	// TODO: Error handling
	return s.Repo.Delete(ctx, id)
}
