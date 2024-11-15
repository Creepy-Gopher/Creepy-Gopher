package servicetest

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/service"
	"fmt"
	"testing"
)

var user *models.User

func init() {
	user = &models.User{
		Model:     models.Model{},
		UserName:  "testuser",
		IsPremium: false,
		IsAdmin:   false,
	}
}

func UserServiceTest(app *service.AppContainer) {
	app.Cfg.Logger.Info("starting test user service")

	ctx := context.Background()

	// // Test Create User
	err := app.UserService().CreateUser(ctx, user)
	if err != nil {
		app.Cfg.Logger.Error(err.Error())
	}

	// Test User Exist (User should exist now)
	exist, _ := app.UserService().UserExist(ctx, user.UserName)
	if !exist {
		app.Cfg.Logger.Error(fmt.Sprintf("User '%v' should exist", user.UserName))
	}

	// Fetch the user by UserName to get the automatically generated ID
	fetchedUser, err := app.UserService().GetByUserName(ctx, user.UserName)
	if err != nil || fetchedUser == nil || fetchedUser.UserName != user.UserName {
		app.Cfg.Logger.Error(fmt.Sprintf("Failed to fetch user by UserName: %v", err))
	}

	// Now the user ID is populated from the fetched user
	userID := fetchedUser.ID

	// Test Get User by ID using the ID obtained from the previous step
	fetchedUserByID, err := app.UserService().GetUser(ctx, userID)
	if err != nil || fetchedUserByID == nil || fetchedUserByID.UserName != user.UserName {
		app.Cfg.Logger.Error(fmt.Sprintf("Failed to fetch user by ID: %v", err))
	}

	// Test Update User
	fetchedUserByID.UserName = "updateduser"
	err = app.UserService().UpdateUser(ctx, fetchedUserByID)
	if err != nil {
		app.Cfg.Logger.Error(fmt.Sprintf("Failed to update user: %v", err))
	}

	// Fetch the user again to verify update
	updatedUser, err := app.UserService().GetUser(ctx, userID)
	if err != nil || updatedUser.UserName != "updateduser" {
		app.Cfg.Logger.Error(fmt.Sprintf("User '%v' update failed: %v", userID, err))
	}

	// Test Delete User
	err = app.UserService().DeleteUser(ctx, userID)
	if err != nil {
		app.Cfg.Logger.Error(fmt.Sprintf("Failed to delete user: %v", err))
	}

	// Test User Exist (Deleted User should not exist)
	exist, err = app.UserService().UserExist(ctx, user.UserName)
	if err != nil || exist {
		app.Cfg.Logger.Error(fmt.Sprintf("User '%v' should not exist after deletion, but got error: %v", user.UserName, err))
	}

	// Test Get By UserName (Non-Existent User)
	_, err = app.UserService().GetByUserName(ctx, "nonexistentuser")
	if err == nil {
		app.Cfg.Logger.Error("Expected error when fetching a non-existent user")
	}

	// soft delete all users
	fetchedUsers, err := app.UserService().AllUsers(ctx)
	if err != nil {
		app.Cfg.Logger.Error(err.Error())
	}

	for _, user := range fetchedUsers {
		fmt.Printf("user_name: %v, id: %v", user.UserName, user.ID)
		err = app.UserService().DeleteUser(ctx, user.ID)
		if err != nil {
			app.Cfg.Logger.Error(err.Error())
		}
	}

	// hard delete all deleted users
	err = app.UserService().DeleteAllSoftDeletedUsers(ctx)
	if err != nil {
		app.Cfg.Logger.Error(err.Error())
	}
}

func TestUserService(t *testing.T) {
	cfg := readConfig()
	cfg.Logger = NewDevelopLogger()
	app, err := service.NewAppContainer(cfg)
	if err != nil {
		app.Cfg.Logger.Error(err.Error())
	}

	UserServiceTest(app)
}
