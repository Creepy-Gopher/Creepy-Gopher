package models

type User struct {
	Model
<<<<<<< HEAD
	TelegramUserID string `gorm:"uniqueIndex;not null"`
	UserName       string
	IsPremium      bool
	IsAdmin        bool
	BookmarkList   []*Bookmark `gorm:"foreignKey:UserID"`
=======
	UserName   string `gorm:"uniqueIndex;not null"`
	TelegramID int64
	IsPremium  bool
	IsAdmin    bool
>>>>>>> fad133822ee34d4cdac68b8bce73da7c30131b60
}
