package models

type User struct {
	Model
	UserName  string `gorm:"uniqueIndex"`
	IsPremium bool
	IsAdmin   bool
}
