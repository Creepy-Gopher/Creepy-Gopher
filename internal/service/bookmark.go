package service

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/storage"
	"errors"

	"github.com/google/uuid"
)

type BookmarkService struct {
	Repo         storage.BookmarkRepository
	PropertyRepo storage.PropertyRepository
}

func NewBookmarkService(repo storage.BookmarkRepository, pRepo storage.PropertyRepository) *BookmarkService {
	return &BookmarkService{
		Repo:         repo,
		PropertyRepo: pRepo,
	}
}

func (s *BookmarkService) CreateBookmark(ctx context.Context, bookmark *models.Bookmark) error {

	_, err := s.PropertyRepo.GetByID(ctx, bookmark.PropertyID)
	if err != nil {
		return errors.New("Invalid property id")
	}

	return s.Repo.CreateBookmark(ctx, bookmark)
}

func (s *BookmarkService) GetBookmarkList(ctx context.Context, id uuid.UUID) (*models.User, error) {
	// TODO: Error handling
	// return s.Repo.GetByID(ctx, id)
}

// func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
// 	// TODO: Error handling
// 	return s.Repo.Update(ctx, user)
// }

// func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
// 	// TODO: Error handling
// 	return s.Repo.Delete(ctx, id)
// }
