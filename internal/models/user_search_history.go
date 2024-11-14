package models

import (
	"time"

	"github.com/google/uuid"
)

type UserSearchHistory struct {
	Model
	UserID   uuid.UUID
	User     User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`  // ForeignKey and Cascade on Delete
	FilterID uuid.UUID `gorm:"index"`                                            // Index for faster lookups
	Filter   Filter    `gorm:"foreignKey:FilterID;constraint:OnDelete:SET NULL"` // ForeignKey and Set Null on Delete
	Date     time.Time `gorm:"not null"`
}
