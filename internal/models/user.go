package models

type User struct {
	Model
	UserName   string `gorm:"uniqueIndex;not null"`
	TelegramID int64
	IsPremium  bool
	IsAdmin    bool
}
