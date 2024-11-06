package entities

import (
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	UserID     uint
	User       User `gorm:"foreignKey:UserID"`
	PropertyID uint
	Property   Property `gorm:"foreignKey:PropertyID"`
}
