package servicetest

import (
	"testing"
	"context"
	"fmt"
	"creepy/internal/models"
	"creepy/internal/service"
	"creepy/internal/storage"
	
	"github.com/google/uuid"
)


type MockPropertyRepository struct {
    Properties map[uuid.UUID]*models.Property
}

func NewMockPropertyRepository() storage.PropertyRepository {
    return &MockPropertyRepository{
        Properties: make(map[uuid.UUID]*models.Property),
    }
}

func (m *MockPropertyRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Property, error) {
    if property, exists := m.Properties[id]; exists {
        return property, nil
    }
    return nil, fmt.Errorf("property with ID %d not found", id)
}

func (m *MockPropertyRepository) Save(ctx context.Context, property *models.Property) error {
    m.Properties[property.ID] = property
    return nil
}

func (m *MockPropertyRepository) Update(ctx context.Context, property *models.Property) error {
    if _, exists := m.Properties[property.ID]; !exists {
        return fmt.Errorf("property with ID %d does not exist", property.ID)
    }
    m.Properties[property.ID] = property
    return nil
}

func (m *MockPropertyRepository) Delete(ctx context.Context, id uuid.UUID) error {
    if _, exists := m.Properties[id]; !exists {
        return fmt.Errorf("property with ID %d does not exist", id)
    }
    delete(m.Properties, id)
    return nil
}

func (m *MockPropertyRepository) ListProperties(ctx context.Context, filter *models.Filter) ([]*models.Property, error) {
    var result []*models.Property
    for _, property := range m.Properties {
        if property.BuyPrice >= filter.BuyPriceMin && property.BuyPrice <= filter.BuyPriceMax { // TODO: more check
            result = append(result, property)
        }
    }
    return result, nil
}


func TestPropertyService_GetProperty(t *testing.T) {
    // Initialize mock repository
    mockRepo := NewMockPropertyRepository()

    // Initialize service with mock repository
    propertyService := service.NewPropertyService(mockRepo)

	ctx := context.Background()
	id := uuid.New()

    // Test CreateProperty with non-existing ID
    err := propertyService.CreateProperty(ctx, &models.Property{
        Model:          models.Model{/*ID: id*/},
        Title:       "Test Property",
        Description: "A property for testing",
    })
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    // Test GetProperty with existing ID
    property, err := propertyService.GetProperty(ctx, id)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if property.Title != "Test Property" {
        t.Errorf("expected title 'Test Property', got '%s'", property.Title)
    }

	// Test GetProperty with non-existing ID
    _, err = propertyService.GetProperty(context.Background(), uuid.New())
    if err == nil {
        t.Fatalf("expected error for non-existing property, got nil")
    }
}