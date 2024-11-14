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

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

// Bot represents the Telegram bot instance
type Bot struct {
	bot             *tele.Bot
	propertyService *service.PropertyService
	logger          *zap.Logger // Zap logger for logging events and errors
}

// NewBot creates a new instance of the Telegram bot
func NewBot(cfg *config.Config, app *service.AppContainer, logger *zap.Logger) (*Bot, error) {
	// Retrieve the bot token from the configuration
	token := cfg.Telegram.BotToken
	if token == "" {
		return nil, fmt.Errorf("توکن بات تلگرام تنظیم نشده است")
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
		//userService:     app.UserService(),
		//filterService:   app.FilterService(),
		logger: logger,
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
	mainMenu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnSearch := mainMenu.Text("🔍 Search")
	btnFilters := mainMenu.Text("⚙️ Filters")
	btnBookmarks := mainMenu.Text("⭐️ Bookmarks")
	btnProfile := mainMenu.Text("👤 Profile")
	btnAddProperty := mainMenu.Text("➕ Add Property")

	mainMenu.Reply(
		mainMenu.Row(btnSearch, btnFilters),
		mainMenu.Row(btnBookmarks, btnProfile),
		mainMenu.Row(btnAddProperty),
	)

	// Start command handler
	b.bot.Handle("/start", func(c tele.Context) error {
		b.logger.Info("User started the bot", zap.Int64("UserID", c.Sender().ID))
		return c.Send("👋 Welcome to the Real Estate Bot! Please select an option:", mainMenu)
	})

	// Search handler
	b.bot.Handle("🔍 Search", b.handleSearch())
	// Filters handler
	b.bot.Handle("⚙️ Filters", b.handleFilters())
	// Bookmarks handler
	//b.bot.Handle("⭐️ Bookmarks", b.handleBookmarks())
	// Profile handler
	//b.bot.Handle("👤 Profile", b.handleProfile())
	// Add property handler
	b.bot.Handle("➕ Add Property", b.handleAddProperty())
}

// handleSearch allows users to search properties
func (b *Bot) handleSearch() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("User accessed search", zap.Int64("UserID", c.Sender().ID))
		return c.Send("Please configure your search filters and press the search button.")
	}
}

// handleFilters allows users to apply filters to search properties
func (b *Bot) handleFilters() tele.HandlerFunc {
	return func(c tele.Context) error {
		filtersMenu := &tele.ReplyMarkup{ResizeKeyboard: true}
		btnPriceRange := filtersMenu.Text("💰 Price Range")
		btnLocation := filtersMenu.Text("📍 Location")
		btnPropertyType := filtersMenu.Text("🏠 Property Type")
		btnFeatures := filtersMenu.Text("✨ Features")
		btnBack := filtersMenu.Text("🔙 Back")

		filtersMenu.Reply(
			filtersMenu.Row(btnPriceRange, btnLocation),
			filtersMenu.Row(btnPropertyType, btnFeatures),
			filtersMenu.Row(btnBack),
		)

		b.logger.Info("User opened filters menu", zap.Int64("UserID", c.Sender().ID))
		return c.Send("Please select a filter option:", filtersMenu)
	}
}

// // handleBookmarks shows the list of bookmarks for the user
// func (b *Bot) handleBookmarks() tele.HandlerFunc {
// 	return func(c tele.Context) error {
// 		b.logger.Info("User viewed bookmarks", zap.Int64("UserID", c.Sender().ID))
// 		bookmarks, err := b.userService.GetBookmarks(context.Background(), c.Sender().ID)
// 		if err != nil {
// 			b.logger.Error("Failed to fetch bookmarks", zap.Error(err))
// 			return c.Send("Error fetching bookmarks.")
// 		}
// 		for _, bookmark := range bookmarks {
// 			c.Send(fmt.Sprintf("Property: %s\nDescription: %s", bookmark.Title, bookmark.Description))
// 		}
// 		return nil
// 	}
//}

// handleProfile shows the user's profile information
// func (b *Bot) handleProfile() tele.HandlerFunc {
// 	return func(c tele.Context) error {
// 		b.logger.Info("User viewed profile", zap.Int64("UserID", c.Sender().ID))
// 		user, err := b.userService.GetUser(context.Background(), c.Sender().ID)
// 		if err != nil {
// 			b.logger.Error("Failed to fetch user profile", zap.Error(err))
// 			return c.Send("Error fetching profile information.")
// 		}
// 		return c.Send(fmt.Sprintf("Profile Information:\nUsername: %s\nRole: %s", user.UserName, user.UserRole))
// 	}
// }

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
