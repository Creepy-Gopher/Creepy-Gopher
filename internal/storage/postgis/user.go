package postgis

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) storage.UserRepository {
	return &userRepo{
		DB: db,
	}
}

func (r *userRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
    var user models.User
    result := r.DB.WithContext(ctx).First(&user, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (r *userRepo) Save(ctx context.Context, entity *models.User) error {
    	if err := r.DB.WithContext(ctx).Save(entity).Error; err != nil {
            return err
        }
        return nil
}

func (r *userRepo) Update(ctx context.Context, entity *models.User) error {
    result := r.DB.WithContext(ctx).Model(&models.User{}).Updates(entity)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", entity.ID)
    }
    return nil
}

func (r *userRepo) Delete(ctx context.Context, id uuid.UUID) error {
    user := models.User{Model: models.Model{ID: id}}
    result := r.DB.WithContext(ctx).Delete(&user, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}

func (r *userRepo) GetByUserName(ctx context.Context, userName string) (*models.User, error) {
    var user models.User
    result := r.DB.WithContext(ctx).Where("user_name = ?", userName).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    if result.RowsAffected == 0 {
        return nil, fmt.Errorf("no record found with user_name %v", userName)
    }
    return &user, nil
}