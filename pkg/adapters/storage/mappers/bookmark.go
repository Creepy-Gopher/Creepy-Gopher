package mappers

import (
	"creepy/internal/bookmark"
	"creepy/pkg/adapters/storage/entities"
)

func BookmarkEntityToDomain(b *entities.Bookmark) *bookmark.Bookmark {
	user := UserEntityToDomain(&b.User)
	property := PropertyEntityToDomain(&b.Property)
	return &bookmark.Bookmark{
		ID:         b.ID,
		UserID:     b.UserID,
		User:       *user,
		PropertyID: b.PropertyID,
		Property:   *property,
	}
}

func BookmarkDomainToEntity(b *bookmark.Bookmark) *entities.Bookmark {
	user := UserDomainToEntity(&b.User)
	property := PropertyDomainToEntity(&b.Property)
	return &entities.Bookmark{
		Model:      entities.Model{ID: b.ID},
		UserID:     b.UserID,
		User:       *user,
		PropertyID: b.PropertyID,
		Property:   *property,
	}
}
