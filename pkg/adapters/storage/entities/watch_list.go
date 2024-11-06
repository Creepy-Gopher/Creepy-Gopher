package entities

import (
	"gorm.io/gorm"
)

type WatchList struct {
	gorm.Model
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	FilterSetID uint
	FilterSet   FilterSet `gorm:"foreignKey:FilterSetID"`
	FilterName  string
}
