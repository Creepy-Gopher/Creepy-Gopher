package service

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"fmt"
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

func (s *FilterService) GetFilter(ctx context.Context, filter *models.Filter) (*models.Filter, error) {
	return s.Repo.GetByFilter(ctx, filter)
}

func (s *FilterService) UpdateFilter(ctx context.Context, filter *models.Filter) error {
	_, err := s.GetFilter(ctx, filter)
	if err == nil {
		return s.Repo.Update(ctx, filter)
	} else {
		return fmt.Errorf("this filter id doesn't  exist")
	}
}

func (s *FilterService) DeleteFilter(ctx context.Context, id uuid.UUID) error {
	_, err := s.GetById(ctx, id)
	if err == nil {
		return s.Repo.Delete(ctx, id)
	} else {
		return fmt.Errorf("invalid id")
	}
}

func (s *FilterService) GetById(ctx context.Context, id uuid.UUID) (*models.Filter, error) {
	return s.Repo.GetByID(ctx, id)
}
