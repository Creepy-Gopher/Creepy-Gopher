package service

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"

	"github.com/google/uuid"
)

type UserService struct {
    Repo storage.UserRepository
}

func NewUserService(repo storage.UserRepository) *UserService {
	// TODO: Error handling
    return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	// TODO: Error handling
    return s.Repo.Save(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
    // TODO: Error handling
    return s.Repo.GetByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	// TODO: Error handling
	return s.Repo.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	// TODO: Error handling
	return s.Repo.Delete(ctx, id)
}