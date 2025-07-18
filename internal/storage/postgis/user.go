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
	result := r.DB.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", entity.ID).
		Updates(entity)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with ID %v", entity.ID)
	}
	return nil
}

func (r *userRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.DB.WithContext(ctx).Where("ID = ?", id).Delete(&models.User{})
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

func (r *userRepo) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) DeleteAllSoftDeletedUsers(ctx context.Context) error {
	result := r.DB.WithContext(ctx).Unscoped().Where("deleted_at IS NOT NULL").Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no soft-deleted users found")
	}
	return nil
}
