package service

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
    Repo storage.UserRepository
}

func NewUserService(repo storage.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
    return s.Repo.Save(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
    return s.Repo.GetByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	return s.Repo.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.Repo.Delete(ctx, id)
}

func (s *UserService) GetByUserName(ctx context.Context, userName string) (*models.User, error) {
    return s.Repo.GetByUserName(ctx, userName)
}

func (s *UserService) UserExist(ctx context.Context, userName string) bool {
	_, err := s.GetByUserName(ctx, userName)
	return err != gorm.ErrRecordNotFound
}

func (s *UserService) AllUsers(ctx context.Context) ([]models.User, error) {
	return s.Repo.GetAllUsers(ctx)
}

func (s *UserService) DeleteAllSoftDeletedUsers(ctx context.Context) error {
	return s.Repo.DeleteAllSoftDeletedUsers(ctx)
}

// super admin can set users as admin 
func (s *UserService) SetUserAsAdmin(ctx context.Context, userName string) error {
	// TODO: caller must be 'super admin'
	return fmt.Errorf("not implemented")
}

// super admin can set admins as user 
func (s *UserService) SetAdminAsUser() error {
	// TODO: caller must be 'super admin'
	return fmt.Errorf("not implemented")
}

// admins can upgrade user to premium account
func (s *UserService) SetUserAsPremiumUser() error {
	// TODO: caller must be 'admin'
	return fmt.Errorf("not implemented")
}

// admins can downgrade premium user to simple user
func (s *UserService) SetPremiumUserAsUser() error {
	// TODO: caller must be 'admin'
	return fmt.Errorf("not implemented")
}
