package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// if you are using a database like PostgreSQL,
// which has support for UUIDs and built-in functions.
// However, in SQLite or MySQL,
// you will generate the UUID in your Go code using uuid.New() (or other libraries) instead.
type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
