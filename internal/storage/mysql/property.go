package mysql

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


// gpt suggestions

type propertyRepo struct {
    DB *gorm.DB
}

func NewMySQLPropertyRepository(db *gorm.DB) storage.PropertyRepository {
    return &propertyRepo{DB: db}
}

func (r *propertyRepo) SaveProperty(ctx context.Context, property *models.Property) error {
    if property.ID == uuid.Nil {
		property.ID = uuid.New()
	}
    if err := r.DB.WithContext(ctx).Save(property).Error; err != nil {
		return err
	}
	return nil
}

func (r *propertyRepo) GetPropertyByID(ctx context.Context, id uuid.UUID) (*models.Property, error) {
    var property models.Property
    result := r.DB.WithContext(ctx).First(&property, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &property, nil
}

// Update attributes with `struct`, will only update non-zero fields
func (r *propertyRepo) UpdateProperty(ctx context.Context, property *models.Property) error {
    result := r.DB.WithContext(ctx).Model(&models.Property{}).Updates(property)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", property.ID)
    }
    return nil
}

func (r *propertyRepo) DeleteProperty(ctx context.Context, id uuid.UUID) error {
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
	return nil, fmt.Errorf("not implemented")
}



// type propertyRepo struct{
// 	db *gorm.DB
// }

// func NewPropertyRepo(db *gorm.DB) *propertyRepo {
// 	return &propertyRepo{
// 		db: db,
// 	}
// }

// func (pr *propertyRepo) Insert(ctx context.Context, property *models.Property) error {
// 	if err := pr.db.WithContext(ctx).Save(property).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (pr *propertyRepo) DeleteByID(ctx context.Context, id *uuid.UUID) error {
// 	property := models.Property{Model: models.Model{ID: *id}}
//     result := pr.db.WithContext(ctx).Delete(&property, id)
//     if result.Error != nil {
//         return result.Error
//     }
//     if result.RowsAffected == 0 {
//         return fmt.Errorf("no record found with ID %v", id)
//     }
//     return nil
// }

// func (pr *propertyRepo) UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error {
//     property := models.Property{Model: models.Model{ID: *id}}
//     result := pr.db.WithContext(ctx).Model(&property).Updates(updates)
//     if result.Error != nil {
//         return result.Error
//     }
//     if result.RowsAffected == 0 {
//         return fmt.Errorf("no record found with ID %v", id)
//     }
//     return nil
// }