package bot

import (
	"context"
	"creepy/internal/models"
	"creepy/internal/service"
	"creepy/pkg/config"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

// Bot represents the Telegram bot instance
type Bot struct {
	bot             *tele.Bot
	propertyService *service.PropertyService
	userService     *service.UserService
	filterService   *service.FilterService
	logger          *zap.Logger // Zap logger for logging events and errors

	mainMenu       *tele.ReplyMarkup
	filtersMenu    *tele.ReplyMarkup
	propertyMenu   *tele.ReplyMarkup
	managementMenu *tele.ReplyMarkup
}

// NewBot creates a new instance of the Telegram bot
func NewBot(cfg *config.Config, app *service.AppContainer, logger *zap.Logger) (*Bot, error) {
	// Retrieve the bot token from the configuration
	token := cfg.Telegram.BotToken
	if token == "" {
		return nil, fmt.Errorf("Token is null")
	}

	// Setup bot settings with polling timeout
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	// Create a new bot instance
	b, err := tele.NewBot(pref)
	if err != nil {
		logger.Error("Failed to create bot", zap.Error(err))
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	logger.Info("Telegram bot initialized")
	return &Bot{
		bot:             b,
		propertyService: app.PropertyService(),
		userService:     app.UserService(),
		filterService:   app.FilterService(),
		logger:          logger,
	}, nil
}

// Start the bot and initialize handlers
func (b *Bot) Start() {
	b.logger.Info("Bot started")
	b.initializeHandlers() // Setup handlers before starting
	b.bot.Start()
}

// initializeHandlers sets up all bot command handlers and main menu
func (b *Bot) initializeHandlers() {
	// Create main menu
	b.mainMenu = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnSearch := b.mainMenu.Text("🔍 Search")
	btnFilters := b.mainMenu.Text("⚙️ Filters")
	btnBookmarks := b.mainMenu.Text("⭐️ Bookmarks")
	btnProfile := b.mainMenu.Text("👤 Profile")
	btnProperty := b.mainMenu.Text("🏘️ Property")
	btnManagement := b.mainMenu.Text("🔧 Management")

	b.mainMenu.Reply(
		b.mainMenu.Row(btnSearch, btnFilters),
		b.mainMenu.Row(btnBookmarks, btnProfile),
		b.mainMenu.Row(btnProperty),
		b.mainMenu.Row(btnManagement),
	)

	// Create filters menu
	b.filtersMenu = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnCreateFilter := b.filtersMenu.Text("➕ Create Filter")
	btnViewFilter := b.filtersMenu.Text("👁 View Filter")
	btnUpdateFilter := b.filtersMenu.Text("✏️ Update Filter")
	btnDeleteFilter := b.filtersMenu.Text("🗑 Delete Filter")
	btnFiltersBack := b.filtersMenu.Text("🔙 Back")

	b.filtersMenu.Reply(
		b.filtersMenu.Row(btnCreateFilter, btnViewFilter),
		b.filtersMenu.Row(btnUpdateFilter, btnDeleteFilter),
		b.filtersMenu.Row(btnFiltersBack),
	)

	// Create property menu
	b.propertyMenu = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnAddProperty := b.propertyMenu.Text("➕ Add Property")
	btnMyProperties := b.propertyMenu.Text("📄 My Properties")
	btnPropertyBack := b.propertyMenu.Text("🔙 Back")

	b.propertyMenu.Reply(
		b.propertyMenu.Row(btnAddProperty, btnMyProperties),
		b.propertyMenu.Row(btnPropertyBack),
	)

	// Create management menu
	b.managementMenu = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnAddUser := b.managementMenu.Text("➕ Add User")
	btnViewUsers := b.managementMenu.Text("👥 View Users")
	btnManagementBack := b.managementMenu.Text("🔙 Back")

	b.managementMenu.Reply(
		b.managementMenu.Row(btnAddUser, btnViewUsers),
		b.managementMenu.Row(btnManagementBack),
	)

	// Start command handler
	b.bot.Handle("/start", func(c tele.Context) error {
		b.logger.Info("User started the bot", zap.Int64("UserID", c.Sender().ID))
		return c.Send("👋 Welcome to the Creepy-Gopher Bot! Please select an option:", b.mainMenu)
	})

	// Main menu handlers
	b.bot.Handle(&btnSearch, b.handleSearch())
	b.bot.Handle(&btnFilters, b.handleFilters())
	// Bookmarks handler
	// b.bot.Handle(&btnBookmarks, b.handleBookmarks())
	// Profile handler
	b.bot.Handle(&btnProfile, b.handleProfile())
	b.bot.Handle(&btnProperty, b.handleProperty())
	b.bot.Handle(&btnManagement, b.handleManagement())

	// Filters menu handlers
	b.bot.Handle(&btnCreateFilter, b.handleCreateFilter())
	b.bot.Handle(&btnViewFilter, b.handleListFilters())
	b.bot.Handle(&btnUpdateFilter, b.handleUpdateFilter())
	b.bot.Handle(&btnDeleteFilter, b.handleDeleteFilter())
	b.bot.Handle(&btnFiltersBack, func(c tele.Context) error {
		return c.Send("Returning to the main menu:", b.mainMenu)
	})

	// Property menu handlers
	b.bot.Handle(&btnAddProperty, b.handleAddProperty())
	b.bot.Handle(&btnMyProperties, b.handleUserProperties())
	b.bot.Handle(&btnPropertyBack, func(c tele.Context) error {
		return c.Send("Returning to the main menu:", b.mainMenu)
	})

	// Management menu handlers
	b.bot.Handle(&btnAddUser, b.handleAddUser())
	// b.bot.Handle(&btnViewUsers, b.handleViewUsers())
	b.bot.Handle(&btnManagementBack, func(c tele.Context) error {
		return c.Send("Returning to the main menu:", b.mainMenu)
	})
}

var telegramNamespaceUUID = uuid.MustParse("12345678-1234-5678-1234-567812345678")

func UUIDFromTelegramID(telegramID int64) uuid.UUID {
	telegramIDStr := strconv.FormatInt(telegramID, 10)
	return uuid.NewSHA1(telegramNamespaceUUID, []byte(telegramIDStr))
}

// handleManagement manages the super admin functionalities
func (b *Bot) handleManagement() tele.HandlerFunc {
	return func(c tele.Context) error {
		telegramID := c.Sender().ID
		userUUID := UUIDFromTelegramID(telegramID)
		fmt.Print(userUUID)
		user, err := b.userService.GetUser(context.Background(), userUUID)
		if err != nil || !user.IsAdmin {
			return c.Send("You do not have access to this menu.")
		}

		return c.Send("Super Admin Management:", b.managementMenu)
	}
}

// handleAddUser adds a new user by Telegram ID
func (b *Bot) handleAddUser() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("Admin is adding a new user", zap.Int64("AdminID", c.Sender().ID))

		// Prompt for Telegram User ID
		if !b.getUserInput(c, "Enter the Telegram User ID of the new admin user:", func(input string) {

			_, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				c.Send("Invalid Telegram ID.")
				return
			}
			user := &models.User{
				UserName:  input,
				IsPremium: false,
				IsAdmin:   true,
			}
			err = b.userService.CreateUser(context.Background(), user)
			if err != nil {
				b.logger.Error("Failed to add new admin user", zap.Error(err))
				c.Send("Failed to create user.")
			} else {
				c.Send("User added successfully.")
			}
		}) {
			return c.Send("Error receiving user ID.")
		}
		return nil
	}
}

// Filters menu with "Back" and "Create Filter" button handler
func (b *Bot) handleProperty() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("User accessed Property menu", zap.Int64("UserID", c.Sender().ID))
		return c.Send("Please select a property action:", b.propertyMenu)
	}
}

// Filters menu with "Back" and "Create Filter" button handler
func (b *Bot) handleFilters() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("User accessed Filters menu", zap.Int64("UserID", c.Sender().ID))
		return c.Send("Please select a filter action:", b.filtersMenu)
	}
}

// Handle creating a new filter
func (b *Bot) handleCreateFilter() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("User started creating a new filter", zap.Int64("UserID", c.Sender().ID))
		// Collect filter fields from the user
		filter := &models.Filter{}

		// Example: Get minimum price
		if !b.getUserInput(c, "Please enter the minimum price for the filter:", func(input string) {
			priceMin, _ := strconv.ParseUint(input, 10, 64)
			filter.BuyPriceMin = priceMin
		}) {
			return c.Send("Error receiving minimum price.")
		}
		// Example: Get maximum price
		if !b.getUserInput(c, "Please enter the maximum price for the filter:", func(input string) {
			priceMax, _ := strconv.ParseUint(input, 10, 64)
			filter.BuyPriceMax = priceMax
		}) {
			return c.Send("Error receiving maximum price.")
		}
		// Save the new filter
		if err := b.filterService.CreateFilter(context.Background(), filter); err != nil {
			b.logger.Error("Failed to save filter", zap.Error(err))
			return c.Send("Error saving the filter: " + err.Error())
		}

		b.logger.Info("Filter successfully created", zap.Int64("UserID", c.Sender().ID))
		return c.Send("Your filter has been successfully created!")
	}
}

// Handle list filters for user
func (b *Bot) handleListFilters() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("User requested filter list", zap.Int64("UserID", c.Sender().ID))
		userIDStr := strconv.FormatInt(c.Sender().ID, 10)
		userUUID, err := uuid.Parse(userIDStr)
		if err != nil {
			b.logger.Error("Failed to parse user ID", zap.Error(err))
			return c.Send("Error retrieving filters.")
		}

		filter, err := b.filterService.GetFilter(context.Background(), userUUID)
		if err != nil {
			b.logger.Error("Failed to fetch user filters", zap.Error(err))
			return c.Send("No filters found. Please create a new filter.")
		}

		// Display filter details
		return c.Send(fmt.Sprintf("Your filter:\nCity: %s\nArea: %d-%d", filter.City, filter.AreaMin, filter.AreaMax))
	}
}

// Handle update filter
func (b *Bot) handleUpdateFilter() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("User requested filter update", zap.Int64("UserID", c.Sender().ID))
		userIDStr := strconv.FormatInt(c.Sender().ID, 10)
		userUUID, err := uuid.Parse(userIDStr)
		if err != nil {
			b.logger.Error("Failed to parse user ID", zap.Error(err))
			return c.Send("Error updating the filter.")
		}

		filter, err := b.filterService.GetFilter(context.Background(), userUUID)
		if err != nil {
			b.logger.Error("Failed to fetch user filter for update", zap.Error(err))
			return c.Send("No filter found to update.")
		}

		if !b.getUserInput(c, "Enter new city:", func(input string) { filter.City = input }) {
			return c.Send("Error receiving new city.")
		}

		if err := b.filterService.UpdateFilter(context.Background(), filter); err != nil {
			b.logger.Error("Failed to update filter", zap.Error(err))
			return c.Send("Error updating the filter.")
		}

		return c.Send("Filter updated successfully!")
	}
}

// Handle delete filter
func (b *Bot) handleDeleteFilter() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("User requested filter deletion", zap.Int64("UserID", c.Sender().ID))
		userIDStr := strconv.FormatInt(c.Sender().ID, 10)
		userUUID, err := uuid.Parse(userIDStr)
		if err != nil {
			b.logger.Error("Failed to parse user ID", zap.Error(err))
			return c.Send("Error deleting the filter.")
		}

		if err := b.filterService.DeleteFilter(context.Background(), userUUID); err != nil {
			b.logger.Error("Failed to delete filter", zap.Error(err))
			return c.Send("Error deleting the filter.")
		}

		return c.Send("Filter deleted successfully.")
	}
}

// handleSearch handles the property search based on user filters
func (b *Bot) handleSearch() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("User requested property search", zap.Int64("UserID", c.Sender().ID))

		userIDStr := strconv.FormatInt(c.Sender().ID, 10)
		userUUID, err := uuid.Parse(userIDStr)
		if err != nil {
			b.logger.Error("Failed to parse user ID", zap.Error(err))
			return c.Send("Error generating user identifier.")
		}

		// Get the user's saved filter
		filter, err := b.filterService.GetFilter(context.Background(), userUUID)
		if err != nil {
			b.logger.Error("Failed to fetch user filter", zap.Error(err))
			return c.Send("Error fetching filters. Please set up your filters first.")
		}

		// Search properties using the user's filter
		properties, err := b.propertyService.ListProperties(context.Background(), filter)
		if err != nil {
			b.logger.Error("Failed to list properties", zap.Error(err))
			return c.Send("Error retrieving properties.")
		}

		if len(properties) == 0 {
			return c.Send("No properties found matching your criteria.")
		}

		for _, property := range properties {
			c.Send(fmt.Sprintf("Property: %s\nDescription: %s\nPrice: %d", property.Title, property.Description, property.BuyPrice))
		}

		return nil
	}
}

// handleProfile shows the user's profile information
func (b *Bot) handleProfile() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("User viewed profile", zap.Int64("UserID", c.Sender().ID))

		telegramID := c.Sender().ID
		userUUID := UUIDFromTelegramID(telegramID)

		user, err := b.userService.GetUser(context.Background(), userUUID)
		if err != nil {
			b.logger.Error("Failed to fetch user profile", zap.Error(err))
			return c.Send("Error fetching profile information.")
		}

		return c.Send(fmt.Sprintf("Profile Information:\nUsername: %v\nIsAdmin: %v\nIsPremium: %v\n", user.UserName, user.IsAdmin, user.IsPremium))
	}
}

// handleUserProperties displays properties created by the user
func (b *Bot) handleUserProperties() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("User requested properties", zap.Int64("UserID", c.Sender().ID))

		// Create an empty filter or pass nil for now
		filter := &models.Filter{}

		// Fetch properties created by the user using their ID
		properties, err := b.propertyService.ListProperties(context.Background(), filter)
		if err != nil {
			b.logger.Error("Failed to fetch properties", zap.Error(err))
			return c.Send("Error retrieving your properties.")
		}

		if len(properties) == 0 {
			return c.Send("No properties found.")
		}

		for _, property := range properties {
			msg := fmt.Sprintf("Title: %s\nDescription: %s\nPrice: %d", property.Title, property.Description, property.BuyPrice)
			c.Send(msg)
		}
		return nil
	}
}

// handleAddProperty guides the user through adding a property listing with full details
func (b *Bot) handleAddProperty() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("User started property listing registration", zap.Int64("UserID", c.Sender().ID))
		property := &models.Property{}

		// Step-by-step input collection for each required field
		if !b.getUserInput(c, "Please enter the property title:", func(input string) { property.Title = input }) {
			return c.Send("Error receiving the title.")
		}
		if !b.getUserInput(c, "Please enter a description:", func(input string) { property.Description = input }) {
			return c.Send("Error receiving the description.")
		}
		if !b.getUserInput(c, "Enter deal type (buy, rent, mortgage):", func(input string) { property.DealingType = input }) {
			return c.Send("Error receiving the deal type.")
		}
		if property.DealingType == "buy" {
			if !b.getUserInput(c, "Enter the buying price:", func(input string) {
				price, _ := strconv.ParseUint(input, 10, 64)
				property.BuyPrice = price
			}) {
				return c.Send("Error receiving buying price.")
			}
		} else if property.DealingType == "rent" {
			if !b.getUserInput(c, "Enter rent range as min-max:", func(input string) {
				prices := strings.Split(input, "-")
				minPrice, _ := strconv.ParseUint(strings.TrimSpace(prices[0]), 10, 64)
				maxPrice, _ := strconv.ParseUint(strings.TrimSpace(prices[1]), 10, 64)
				property.RentPriceMin = minPrice
				property.RentPriceMax = maxPrice
			}) {
				return c.Send("Error receiving rent range.")
			}
		} else if property.DealingType == "mortgage" {
			if !b.getUserInput(c, "Enter mortgage range as min-max:", func(input string) {
				prices := strings.Split(input, "-")
				minPrice, _ := strconv.ParseUint(strings.TrimSpace(prices[0]), 10, 64)
				maxPrice, _ := strconv.ParseUint(strings.TrimSpace(prices[1]), 10, 64)
				property.RahnPriceMin = minPrice
				property.RahnPriceMax = maxPrice
			}) {
				return c.Send("Error receiving mortgage range.")
			}
		}
		if !b.getUserInput(c, "Enter area in square meters:", func(input string) {
			area, _ := strconv.ParseUint(input, 10, 64)
			property.Area = area
		}) {
			return c.Send("Error receiving area.")
		}
		if !b.getUserInput(c, "Enter number of rooms:", func(input string) {
			rooms, _ := strconv.Atoi(input)
			property.Rooms = uint(rooms)
		}) {
			return c.Send("Error receiving the number of rooms.")
		}
		if !b.getUserInput(c, "Enter building type (apartment or villa):", func(input string) { property.Type = input }) {
			return c.Send("Error receiving the building type.")
		}
		if !b.getUserInput(c, "Enter city:", func(input string) { property.City = input }) {
			return c.Send("Error receiving the city.")
		}
		if !b.getUserInput(c, "Enter district:", func(input string) { property.District = input }) {
			return c.Send("Error receiving the district.")
		}
		if !b.getUserInput(c, "Enter the build year:", func(input string) {
			buildYear, _ := strconv.Atoi(input)
			property.BuildYear = uint(buildYear)
		}) {
			return c.Send("Error receiving the build year.")
		}
		if !b.getUserInput(c, "Enter floor number:", func(input string) {
			floor, _ := strconv.Atoi(input)
			property.Floor = uint(floor)
		}) {
			return c.Send("Error receiving the floor number.")
		}
		if !b.getUserInput(c, "Does it have an elevator? (yes/no):", func(input string) {
			property.HasElevator = strings.ToLower(input) == "yes"
		}) {
			return c.Send("Error receiving elevator information.")
		}
		if !b.getUserInput(c, "Does it have storage? (yes/no):", func(input string) {
			property.HasStorage = strings.ToLower(input) == "yes"
		}) {
			return c.Send("Error receiving storage information.")
		}
		if !b.getUserInput(c, "Does it have parking? (yes/no):", func(input string) {
			property.HasParking = strings.ToLower(input) == "yes"
		}) {
			return c.Send("Error receiving parking information.")
		}
		if !b.getUserInput(c, "Enter latitude:", func(input string) {
			latitude, _ := strconv.ParseFloat(input, 64)
			property.Latitude = latitude
		}) {
			return c.Send("Error receiving latitude.")
		}
		if !b.getUserInput(c, "Enter longitude:", func(input string) {
			longitude, _ := strconv.ParseFloat(input, 64)
			property.Longitude = longitude
		}) {
			return c.Send("Error receiving longitude.")
		}
		if !b.getUserInput(c, "Enter property URL:", func(input string) { property.URL = input }) {
			return c.Send("Error receiving the property URL.")
		}

		// Save the property using PropertyService
		if err := b.propertyService.CreateProperty(context.Background(), property); err != nil {
			b.logger.Error("Failed to save property", zap.Error(err))
			return c.Send("Error saving the property: " + err.Error())
		}

		b.logger.Info("Property successfully saved", zap.String("Title", property.Title), zap.Int64("UserID", c.Sender().ID))
		return c.Send("Your property listing has been successfully saved!")
	}
}

// getUserInput is a helper function to simplify user input handling
func (b *Bot) getUserInput(c tele.Context, prompt string, handler func(input string)) bool {
	if err := c.Send(prompt); err != nil {
		b.logger.Error("Failed to send prompt", zap.Error(err))
		return false
	}

	// Wait for the user's response
	received := make(chan struct{})
	c.Bot().Handle(tele.OnText, func(ct tele.Context) error {
		handler(ct.Text())
		received <- struct{}{}
		return nil
	})

	select {
	case <-received:
		return true
	case <-time.After(30 * time.Second):
		b.logger.Warn("User input timed out", zap.Int64("UserID", c.Sender().ID))
		return false
	}
}
