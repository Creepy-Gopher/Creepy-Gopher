package postgis

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type propertyRepo struct {
    DB *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) storage.PropertyRepository {
    return &propertyRepo{DB: db}
}

func (r *propertyRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Property, error) {
    var property models.Property
    result := r.DB.WithContext(ctx).First(&property, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &property, nil
}

func (r *propertyRepo) Save(ctx context.Context, entity *models.Property) error {
    if err := r.DB.WithContext(ctx).Save(entity).Error; err != nil {
        return err
    }
    return nil
}

// Update attributes with `struct`, will only update non-zero fields
func (r *propertyRepo) Update(ctx context.Context, entity *models.Property) error {
    result := r.DB.WithContext(ctx).Model(&models.Property{}).Updates(entity)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", entity.ID)
    }
    return nil
}

func (r *propertyRepo) Delete(ctx context.Context, id uuid.UUID) error {
    property := models.Property{Model: models.Model{ID: id}}
    result := r.DB.WithContext(ctx).Delete(&property, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}

func (r *propertyRepo) ListProperties(ctx context.Context, filter *models.Filter) ([]*models.Property, error) {
	var properties []*models.Property
	query := r.DB.WithContext(ctx).Model(&models.Property{})

	// Apply filters if they are provided

	// Filter by city
	if filter.City != "" {
		query = query.Where("city = ?", filter.City)
	}

	// Filter by buy price range
	if filter.BuyPriceMin > 0 {
		query = query.Where("buy_price >= ?", filter.BuyPriceMin)
	}
	if filter.BuyPriceMax > 0 {
		query = query.Where("buy_price <= ?", filter.BuyPriceMax)
	}

	// Filter by rent price range
	if filter.RentPriceMin > 0 {
		query = query.Where("rent_price >= ?", filter.RentPriceMin)
	}
	if filter.RentPriceMax > 0 {
		query = query.Where("rent_price <= ?", filter.RentPriceMax)
	}

	// Filter by area range
	if filter.AreaMin > 0 {
		query = query.Where("area >= ?", filter.AreaMin)
	}
	if filter.AreaMax > 0 {
		query = query.Where("area <= ?", filter.AreaMax)
	}

	// Filter by rooms range
	if filter.RoomMin > 0 {
		query = query.Where("rooms >= ?", filter.RoomMin)
	}
	if filter.RoomMax > 0 {
		query = query.Where("rooms <= ?", filter.RoomMax)
	}

	// Filter by floor range
	if filter.FloorMin > 0 {
		query = query.Where("floor >= ?", filter.FloorMin)
	}
	if filter.FloorMax > 0 {
		query = query.Where("floor <= ?", filter.FloorMax)
	}

	// Filter by build year range
	if filter.BuildYearMin > 0 {
		query = query.Where("build_year >= ?", filter.BuildYearMin)
	}
	if filter.BuildYearMax > 0 {
		query = query.Where("build_year <= ?", filter.BuildYearMax)
	}

	// Filter by property type (buy, rent, rahn)
	if filter.DealingType != "" {
		query = query.Where("dealing_type = ?", filter.DealingType)
	}

	if filter.HasElevator {
		query = query.Where("has_elevator = ?", true)
	}
	if filter.HasStorage {
		query = query.Where("has_storage = ?", true)
	}
	if filter.HasParking {
		query = query.Where("has_parking = ?", true)
	}

	// Filter by location
	if filter.LocationLatitude != 0 && filter.LocationLongitude != 0 && filter.LocationRadius > 0 {
		// query = query.Where(
		// 	"ST_DWithin(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326), ?)", 
		// 	filter.LocationLongitude, filter.LocationLatitude, filter.LocationRadius)
	}

	if filter.District != "" {
		query = query.Where("district = ?", filter.District)
	}

	if filter.Source != "" {
		query = query.Where("source = ?", filter.Source)
	}

	// Execute the query
	result := query.Find(&properties)
	if result.Error != nil {
		return nil, result.Error
	}

	return properties, nil
}

func (r *propertyRepo) GetPropertyByURL(ctx context.Context, url string) (*models.Property, error) {
    var property models.Property
    result := r.DB.WithContext(ctx).Where("ulr = ?", url).First(&property)
    if result.Error != nil {
        return nil, result.Error
    }
    if result.RowsAffected == 0 {
        return nil, fmt.Errorf("no record found with URL %v", url)
    }
    return &property, nil
}
