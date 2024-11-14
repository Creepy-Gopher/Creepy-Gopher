package service

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	Repo storage.UserRepository
}

func NewUserService(repo storage.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	if user.UserName == "" {
		return fmt.Errorf("username can't be empty")
	}
	return s.Repo.Save(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	_, err := s.GetUser(ctx, user.ID)
	if err == nil {
		return s.Repo.Update(ctx, user)
	} else {
		return fmt.Errorf("this user id doesn't  exist")
	}
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := s.GetUser(ctx, id)
	if err == nil {
		return s.Repo.Delete(ctx, id)
	} else {
		return fmt.Errorf("invalid id")
	}
}

func (s *UserService) GetByUserName(ctx context.Context, userName string) (*models.User, error) {
	return s.Repo.GetByUserName(ctx, userName)
}
