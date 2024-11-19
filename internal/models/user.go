package models

type User struct {
	Model
	TelegramID   int64 `gorm:"uniqueIndex;not null"`
	UserName     string
	IsPremium    bool
	IsAdmin      bool
	BookmarkList []*Bookmark `gorm:"foreignKey:UserID"`
}
