package mappers

import (
	"creepy/internal/user"
	"creepy/pkg/adapters/storage/entities"
)

func UserEntityToDomain(u *entities.User) *user.User {
	return &user.User{
		ID:        u.ID,
		UserName:  u.UserName,
		Role:      u.Role,
		IsPremium: u.IsPremium,
	}
}

func UserDomainToEntity(u *user.User) *entities.User {
	return &entities.User{
		Model:     entities.Model{ID: u.ID},
		UserName:  u.UserName,
		Role:      u.Role,
		IsPremium: u.IsPremium,
	}
}
