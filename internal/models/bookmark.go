package models

import (
	"github.com/google/uuid"
)

type Bookmark struct {
	Model
	UserID     uuid.UUID `gorm:"type:uuid"`
	User       User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	PropertyID uuid.UUID
	Property   Property `gorm:"foreignKey:PropertyID;constraint:OnDelete:CASCADE"`
}
