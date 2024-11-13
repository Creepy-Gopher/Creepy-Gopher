package models

import (
	"time"

	"github.com/google/uuid"
)

type UserSearchHistory struct {
	Model
	UserName string    `gorm:"index;not null"`                                   // Index for faster lookups
	User     User      `gorm:"foreignKey:UserName;constraint:OnDelete:CASCADE"`  // ForeignKey and Cascade on Delete
	FilterID uuid.UUID `gorm:"index"`                                            // Index for faster lookups
	Filter   Filter    `gorm:"foreignKey:FilterID;constraint:OnDelete:SET NULL"` // ForeignKey and Set Null on Delete
	Date     time.Time `gorm:"not null"`
}
