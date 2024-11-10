package service

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"errors"

	"github.com/google/uuid"
)

type PropertyService struct {
    Repo storage.PropertyRepository
}

func NewPropertyService(repo storage.PropertyRepository) *PropertyService {
    return &PropertyService{Repo: repo}
}

func (s *PropertyService) CreateProperty(ctx context.Context, property *models.Property) error {
    return s.Repo.SaveProperty(ctx, property)
}

func (s *PropertyService) GetProperty(ctx context.Context, id uuid.UUID) (*models.Property, error) {
    return s.Repo.GetPropertyByID(ctx, id)
}

func (s *PropertyService) UpdateProperty(ctx context.Context, property *models.Property) error {
	return s.Repo.UpdateProperty(ctx, property)
}

func (s *PropertyService) DeleteProperty(ctx context.Context, id uuid.UUID) error {
	return s.Repo.DeleteProperty(ctx, id)
}

// Additional business logic methods...

func (s *PropertyService) ListProperties(ctx context.Context, filter *models.Filter) ([]*models.Property, error) {
	return nil, errors.New("not implemented")
}
