package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string
	Role      string
	IsPremium bool

	// Relationships
	// Bookmarks []Bookmark
}
