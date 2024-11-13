package models

type User struct {
	Model
	UserName  string `gorm:"uniqueIndex"`
	Role      string
	IsPremium bool
	IsAdmin   bool
}
