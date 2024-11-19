package models

type User struct {
	Model
	TelegramUserID string `gorm:"uniqueIndex;not null"`
	UserName       string
	IsPremium      bool
	IsAdmin        bool
	BookmarkList   []*Bookmark `gorm:"foreignKey:UserID"`
}
