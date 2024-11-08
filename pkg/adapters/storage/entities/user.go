package entities

type User struct {
	Model
	UserName  string `gorm:"uniqueIndex"`
	Role      string
	IsPremium bool
}
