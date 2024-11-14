package service

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	//"creepy/internal/storage/postgis"
	"fmt"

	"github.com/google/uuid"
)

type PropertyService struct {
	Repo storage.PropertyRepository
}

func NewPropertyService(repo storage.PropertyRepository) *PropertyService {
	return &PropertyService{Repo: repo}
}

func (s *PropertyService) CreateProperty(ctx context.Context, property *models.Property) error {
	if property.URL == "" {
		return fmt.Errorf("URL field is empty")
	}
	return s.Repo.Save(ctx, property)
}

func (s *PropertyService) GetProperty(ctx context.Context, id uuid.UUID) (*models.Property, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *PropertyService) UpdateProperty(ctx context.Context, property *models.Property) error {
	_, err := s.GetProperty(ctx, property.ID)
	if err == nil {
		return s.Repo.Update(ctx, property)
	} else {
		return fmt.Errorf("this property id doesn't  exist")
	}
}

func (s *PropertyService) DeleteProperty(ctx context.Context, id uuid.UUID) error {
	_, err := s.GetProperty(ctx, id)
	if err == nil {
		return s.Repo.Delete(ctx, id)
	} else {
		return fmt.Errorf("invalid id")
	}

}

func (s *PropertyService) ListProperties(ctx context.Context, filter *models.Filter) ([]*models.Property, error) {
	if filter.AreaMin > filter.AreaMax {
		return nil, fmt.Errorf("invalid range: filter area ")
	}
	if filter.FloorMin > filter.FloorMax {
		return nil, fmt.Errorf("invalid range: filter floor ")
	}
	if filter.RoomMin > filter.RoomMax {
		return nil, fmt.Errorf("invalid range: filter room ")
	}
	if filter.BuildYearMin > filter.BuildYearMax {
		return nil, fmt.Errorf("invalid range: filter build year ")
	}
	if filter.BuyPriceMin > filter.BuyPriceMax {
		return nil, fmt.Errorf("invalid range: filter buy price ")
	}
	if filter.RentPriceMin > filter.RentPriceMax {
		return nil, fmt.Errorf("invalid range: filter rent price ")
	}

	return s.Repo.ListProperties(ctx, filter)
}

func (s *PropertyService) CreatePropertyByAdmin(ctx context.Context, property *models.Property) error {
	userContextKey := "user"
	user, ok := ctx.Value(userContextKey).(*models.User)
	if !ok {
		return fmt.Errorf("context has not this user admin")
	}
	if !user.IsAdmin {
		return fmt.Errorf("permission denied")
	}
	property.Source = "admin"
	return s.Repo.Save(ctx, property)
}

func (s *PropertyService) CreatePropertyByCrawler(ctx context.Context, property *models.Property) error {
	property.Source = "crawler"
	return s.Repo.Save(ctx, property)
}
