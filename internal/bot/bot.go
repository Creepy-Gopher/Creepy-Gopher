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

	"github.com/google/uuid" // افزودن پکیج UUID
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
func NewBot(cfg *config.Config, propertyService *service.PropertyService, logger *zap.Logger) (*Bot, error) {
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
		propertyService: propertyService,
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
	mainMenu := &tele.ReplyMarkup{ResizeKeyboard: true}

	// Main menu buttons
	btnSearch := mainMenu.Text("🔍 جستجو")
	btnFilters := mainMenu.Text("⚙️ فیلترها")
	btnBookmarks := mainMenu.Text("⭐️ علاقه‌مندی‌ها")
	btnWatchlist := mainMenu.Text("👀 لیست پیگیری")
	btnPremium := mainMenu.Text("💎 حساب ویژه")
	btnProfile := mainMenu.Text("👤 پروفایل")
	btnAddProperty := mainMenu.Text("➕ ثبت آگهی")

	// Arrange buttons in rows
	mainMenu.Reply(
		mainMenu.Row(btnSearch, btnFilters),
		mainMenu.Row(btnBookmarks, btnWatchlist),
		mainMenu.Row(btnPremium, btnProfile),
		mainMenu.Row(btnAddProperty),
	)

	// Handler for the /start command
	b.bot.Handle("/start", func(c tele.Context) error {
		b.logger.Info("User started the bot", zap.Int64("UserID", c.Sender().ID))
		return c.Send("👋 به بات جستجوی املاک خوش آمدید!\n\n"+
			"از منوی زیر گزینه مورد نظر خود را انتخاب کنید:", mainMenu)
	})

	// Handler for each main menu button
	b.bot.Handle("🔍 جستجو", func(c tele.Context) error {
		b.logger.Info("User selected search", zap.Int64("UserID", c.Sender().ID))
		return c.Send("لطفاً فیلترهای جستجو را تنظیم کنید و سپس دکمه جستجو را بزنید.")
	})
	b.bot.Handle("⚙️ فیلترها", b.handleFilters())
	b.bot.Handle("⭐️ علاقه‌مندی‌ها", func(c tele.Context) error {
		b.logger.Info("User viewed bookmarks", zap.Int64("UserID", c.Sender().ID))
		return c.Send("لیست علاقه‌مندی‌های شما خالی است.")
	})
	b.bot.Handle("👀 لیست پیگیری", func(c tele.Context) error {
		b.logger.Info("User viewed watchlist", zap.Int64("UserID", c.Sender().ID))
		return c.Send("لیست پیگیری شما خالی است.")
	})
	b.bot.Handle("💎 حساب ویژه", func(c tele.Context) error {
		b.logger.Info("User selected premium", zap.Int64("UserID", c.Sender().ID))
		return c.Send("برای ارتقا به حساب ویژه با ادمین در ارتباط باشید.")
	})
	b.bot.Handle("👤 پروفایل", func(c tele.Context) error {
		b.logger.Info("User viewed profile", zap.Int64("UserID", c.Sender().ID))
		return c.Send("اطلاعات پروفایل شما:\n" + "نام کاربری: " + c.Sender().Username)
	})

	// Handler for adding a new property
	b.bot.Handle("➕ ثبت آگهی", b.handleAddProperty())
}

// handleAddProperty manages the process of registering a new property listing
func (b *Bot) handleAddProperty() tele.HandlerFunc {
	return func(c tele.Context) error {
		b.logger.Info("User started property listing registration", zap.Int64("UserID", c.Sender().ID))

		// Create a new property with a generated UUID for ID
		property := &models.Property{
			ID: uuid.New(), // Assign a new UUID
		}

		// Step 1: Title
		if !b.getUserInput(c, "لطفاً عنوان آگهی را وارد کنید:", func(input string) {
			property.Title = input
		}) {
			return c.Send("خطا در دریافت عنوان آگهی.")
		}

		// Step 2: Description
		if !b.getUserInput(c, "توضیحات آگهی را وارد کنید:", func(input string) {
			property.Description = input
		}) {
			return c.Send("خطا در دریافت توضیحات.")
		}

		// Step 3: Dealing Type
		if !b.getUserInput(c, "لطفاً نوع معامله (خرید: buy، اجاره: rent، رهن: rahn) را وارد کنید:", func(input string) {
			property.DealingType = input
		}) {
			return c.Send("خطا در دریافت نوع معامله.")
		}

		// Step 4: Price
		if property.DealingType == "buy" {
			if !b.getUserInput(c, "قیمت خرید را وارد کنید:", func(input string) {
				price, _ := strconv.ParseUint(input, 10, 64)
				property.BuyPrice = price
			}) {
				return c.Send("خطا در دریافت قیمت.")
			}
		} else {
			if !b.getUserInput(c, "بازه قیمت اجاره را به این شکل وارد کنید: حداقل - حداکثر", func(input string) {
				prices := strings.Split(input, "-")
				minPrice, _ := strconv.ParseUint(strings.TrimSpace(prices[0]), 10, 64)
				maxPrice, _ := strconv.ParseUint(strings.TrimSpace(prices[1]), 10, 64)
				property.RentPriceMin = minPrice
				property.RentPriceMax = maxPrice
			}) {
				return c.Send("خطا در دریافت بازه قیمت.")
			}
		}

		// Step 5: Area
		if !b.getUserInput(c, "متراژ ملک را وارد کنید:", func(input string) {
			area, _ := strconv.ParseUint(input, 10, 64)
			property.Area = area
		}) {
			return c.Send("خطا در دریافت متراژ.")
		}

		// Step 6: Rooms
		if !b.getUserInput(c, "تعداد اتاق‌ها را وارد کنید:", func(input string) {
			rooms, _ := strconv.Atoi(input)
			property.Rooms = uint(rooms)
		}) {
			return c.Send("خطا در دریافت تعداد اتاق‌ها.")
		}

		// Save the property using PropertyService
		if err := b.propertyService.CreateProperty(context.Background(), property); err != nil {
			b.logger.Error("Failed to save property", zap.Error(err))
			return c.Send("خطا در ثبت آگهی: " + err.Error())
		}

		b.logger.Info("Property successfully saved", zap.String("Title", property.Title), zap.Int64("UserID", c.Sender().ID))
		return c.Send("آگهی شما با موفقیت ثبت شد!")
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

// handleFilters provides a submenu for filters
func (b *Bot) handleFilters() tele.HandlerFunc {
	return func(c tele.Context) error {
		filtersMenu := &tele.ReplyMarkup{ResizeKeyboard: true}

		btnPriceRange := filtersMenu.Text("💰 محدوده قیمت")
		btnLocation := filtersMenu.Text("📍 موقعیت")
		btnPropertyType := filtersMenu.Text("🏠 نوع ملک")
		btnFeatures := filtersMenu.Text("✨ امکانات")
		btnBack := filtersMenu.Text("🔙 بازگشت")

		filtersMenu.Reply(
			filtersMenu.Row(btnPriceRange, btnLocation),
			filtersMenu.Row(btnPropertyType, btnFeatures),
			filtersMenu.Row(btnBack),
		)

		b.logger.Info("User opened filters menu", zap.Int64("UserID", c.Sender().ID))
		return c.Send("لطفاً فیلتر مورد نظر را انتخاب کنید:", filtersMenu)
	}
}
