package storage

import (
	"context"
	"creepy/internal/user"
	"creepy/pkg/adapters/storage/entities"
	"creepy/pkg/adapters/storage/mappers"
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

func (ur *userRepo) Insert(ctx context.Context, u *user.User) error {
	userEntity := mappers.UserDomainToEntity(u)
	if err := ur.db.WithContext(ctx).Save(&userEntity).Error; err != nil {
		return err
	}
	u.ID = userEntity.ID
	return nil
}

func (ur *userRepo) DeleteByID(ctx context.Context, id *uuid.UUID) error {
	user := entities.User{Model: entities.Model{ID: *id}}
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
    user := entities.User{Model: entities.Model{ID: *id}}
    result := ur.db.WithContext(ctx).Model(&user).Updates(updates)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no record found with ID %v", id)
    }
    return nil
}
