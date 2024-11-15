package data

import "creepy/internal/models"

var Users []models.User

func init() {
	Users = []models.User{
		{
			UserName: "john_doe",
			IsPremium: false,
			IsAdmin: false,
		},
		{
			UserName: "jane_smith",
			IsPremium: true,
			IsAdmin: false,
		},
		{
			UserName: "admin_user",
			IsPremium: true,
			IsAdmin: true,
		},
		{
			UserName: "alex_jones",
			IsPremium: false,
			IsAdmin: true,
		},
	}
}
