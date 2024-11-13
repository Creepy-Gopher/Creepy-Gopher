package models

import "github.com/google/uuid"

type Bookmark struct {
	Model
	UserName   string
	User       User     `gorm:"foreignKey:UserName"`
	PropertyID uuid.UUID
	Property   Property `gorm:"foreignKey:PropertyID"`
}
