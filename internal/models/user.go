package models

type User struct {
	Model
	UserName  string `gorm:"uniqueIndex;not null"`
	IsPremium bool
	IsAdmin   bool
}
