package service

import (
	"context"
	"fmt"
	"creepy/internal/models"
	"creepy/internal/storage"
	"errors"

	"github.com/google/uuid"
)

type PropertyService struct {
    Repo storage.PropertyRepository
}

func NewPropertyService(repo storage.PropertyRepository) *PropertyService {
	// TODO: Error handling
    return &PropertyService{Repo: repo}
}

func (s *PropertyService) CreateProperty(ctx context.Context, property *models.Property) error {
	if property.ID == uuid.Nil {
        return fmt.Errorf("cant save property without ID")
	}
	// TODO: Error handling
    return s.Repo.Save(ctx, property)
}

func (s *PropertyService) GetProperty(ctx context.Context, id uuid.UUID) (*models.Property, error) {
    // TODO: Error handling
    return s.Repo.GetByID(ctx, id)
}

func (s *PropertyService) UpdateProperty(ctx context.Context, property *models.Property) error {
	// TODO: Error handling
	return s.Repo.Update(ctx, property)
}

func (s *PropertyService) DeleteProperty(ctx context.Context, id uuid.UUID) error {
	// TODO: Error handling
	return s.Repo.Delete(ctx, id)
}

func (s *PropertyService) ListProperties(ctx context.Context, filter *models.Filter) ([]*models.Property, error) {
	// TODO
	return nil, errors.New("not implemented")
}
