package models

import "github.com/google/uuid"

type WatchList struct {
	Model
	UserID     uuid.UUID
	User       User      `gorm:"foreignKey:UserID"`
	FilterID   uuid.UUID
	Filter     Filter    `gorm:"foreignKey:FilterID"`
	FilterName string
}