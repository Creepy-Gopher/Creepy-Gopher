package mysql

import (
	"context"
	"creepy/internal/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Insert(ctx context.Context, user *models.User) error {
	if err := ur.db.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) DeleteByID(ctx context.Context, id *uuid.UUID) error {
	user := models.User{Model: models.Model{ID: *id}}
    result := ur.db.WithContext(ctx).Delete(&user, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}

func (ur *userRepo) UpdateByID(ctx context.Context, id *uuid.UUID, updates map[string]interface{}) error {
    user := models.User{Model: models.Model{ID: *id}}
    result := ur.db.WithContext(ctx).Model(&user).Updates(updates)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}
