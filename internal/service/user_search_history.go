package service

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"

	"github.com/google/uuid"
)

type UserSearchHistoryService struct {
	Repo storage.UserSearchHistoryRepository
}

func NewUserSearchHistoryService(repo storage.UserSearchHistoryRepository) *UserSearchHistoryService {
	return &UserSearchHistoryService{Repo: repo}
}

func (s *UserSearchHistoryService) CreateUserHistory(ctx context.Context, userHistory *models.UserSearchHistory) error {
    return s.Repo.Save(ctx, userHistory)
}

// func (s *UserSearchHistoryService) GetUserHistory(ctx context.Context, userName string) (*models.UserSearchHistory, error) {
//     return s.Repo.GetByUserName(ctx, userName)
// }

func (s *UserSearchHistoryService) UpdateUserHistory(ctx context.Context, userHistory *models.UserSearchHistory) error {
	return s.Repo.Update(ctx, userHistory)
}

func (s *UserSearchHistoryService) DeleteUserHistory(ctx context.Context, id uuid.UUID) error {
	return s.Repo.Delete(ctx, id)
}

func (s *UserSearchHistoryService) ListSearchHistoryByUserName(ctx context.Context, userName string) ([]*models.UserSearchHistory, error) {
	return s.Repo.ListSearchHistoryByUserName(ctx, userName)
}
