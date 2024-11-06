package entities

import (
	"gorm.io/gorm"
	"time"
)

type UserSearchHistory struct {
	gorm.Model
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	FilterSetID uint
	FilterSet   FilterSet `gorm:"foreignKey:FilterSetID"`
	Date        time.Time

	// // option
	// IsWatched bool
	// FilterName string
}
