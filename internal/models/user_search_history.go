package models

import (
	"time"

	"github.com/google/uuid"
)

type UserSearchHistory struct {
	Model
	UserName string
	User     User `gorm:"foreignKey:UserName;constraint:OnDelete:CASCADE"`
	FilterID uuid.UUID
	Filter   Filter `gorm:"foreignKey:FilterID;constraint:OnDelete:SET NULL"`
	Date     time.Time
}