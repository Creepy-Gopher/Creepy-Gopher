package models

import (
	"time"

	"github.com/google/uuid"
)

type UserSearchHistory struct {
	Model
	UserID   uuid.UUID `gorm:"type:uuid"`
	User     User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	FilterID uuid.UUID
	Filter   Filter `gorm:"foreignKey:FilterID;constraint:OnDelete:SET NULL"`
	Date     time.Time
}
