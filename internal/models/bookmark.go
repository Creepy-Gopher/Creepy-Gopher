package models

import "github.com/google/uuid"

type Bookmark struct {
	Model
	UserID     uuid.UUID
	User       User     `gorm:"foreignKey:UserID"`
	PropertyID uuid.UUID
	Property   Property `gorm:"foreignKey:PropertyID"`
	SharedWithMe bool
}
