package models

import "github.com/google/uuid"

type WatchList struct {
	Model
	UserName   string
	User       User      `gorm:"foreignKey:UserName"`
	FilterID   uuid.UUID
	Filter     Filter    `gorm:"foreignKey:FilterID"`
	FilterName string
}