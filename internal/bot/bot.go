package bot

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	tele "gopkg.in/telebot.v3"
)

// Bot represents the Telegram bot instance
type Bot struct {
	bot      *tele.Bot
	handlers map[string]tele.HandlerFunc
	mu       sync.RWMutex
}

// Menu represents a menu structure for the bot
type Menu struct {
	Title   string
	Buttons []tele.Btn
	Handler tele.HandlerFunc
}

// Filter represents search filters
type Filter struct {
	PriceRange   [2]float64
	City         string
	Neighborhood string
	AreaRange    [2]float64
	BedroomRange [2]int
	PropertyType string // rent, buy, mortgage
	BuildingAge  [2]int
	BuildingType string // apartment, villa
	Floor        [2]int
	HasStorage   *bool
	HasElevator  *bool
	DateRange    [2]time.Time
	Location     struct {
		Lat    float64
		Long   float64
		Radius float64
	}
}

func main() {
	bot, err := NewBot("")
	if err != nil {
		log.Fatalf("خطا در ایجاد بات: %v", err)
	}
	bot.Start()
}

// NewBot creates a new instance of the Telegram bot
func NewBot(token string) (*Bot, error) {
	token = os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("توکن بات تلگرام تنظیم نشده است")
	}
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	bot := &Bot{
		bot:      b,
		handlers: make(map[string]tele.HandlerFunc),
	}

	// Initialize handlers
	bot.initializeHandlers()

	return bot, nil
}

// Start starts the bot
func (b *Bot) Start() {
	log.Println("Bot started...")
	b.bot.Start()
}

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

	mainMenu.Reply(
		mainMenu.Row(btnSearch, btnFilters),
		mainMenu.Row(btnBookmarks, btnWatchlist),
		mainMenu.Row(btnPremium, btnProfile),
	)

	// Start command handler
	b.bot.Handle("/start", func(c tele.Context) error {
		return c.Send("👋 به بات جستجوی املاک خوش آمدید!\n\n"+
			"از منوی زیر گزینه مورد نظر خود را انتخاب کنید:", mainMenu)
	})

	// Search handler
	b.bot.Handle("🔍 جستجو", func(c tele.Context) error {
		// TODO: Implement search service integration
		return c.Send("لطفاً فیلترهای جستجو را تنظیم کنید و سپس دکمه جستجو را بزنید.")
	})

	// Filters handler
	b.bot.Handle("⚙️ فیلترها", b.handleFilters())

	// Bookmarks handler
	b.bot.Handle("⭐️ علاقه‌مندی‌ها", func(c tele.Context) error {
		// TODO: Implement bookmarks service integration
		return c.Send("لیست علاقه‌مندی‌های شما خالی است.")
	})

	// Watchlist handler
	b.bot.Handle("👀 لیست پیگیری", func(c tele.Context) error {
		// TODO: Implement watchlist service integration
		return c.Send("لیست پیگیری شما خالی است.")
	})

	// Premium handler
	b.bot.Handle("💎 حساب ویژه", func(c tele.Context) error {
		return c.Send("برای ارتقا به حساب ویژه با ادمین در ارتباط باشید.")
	})

	// Profile handler
	b.bot.Handle("👤 پروفایل", func(c tele.Context) error {
		// TODO: Implement user service integration
		return c.Send("اطلاعات پروفایل شما:\n" +
			"نام کاربری: " + c.Sender().Username)
	})
}

func (b *Bot) handleFilters() tele.HandlerFunc {
	return func(c tele.Context) error {
		// Create filters menu
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

		return c.Send("لطفاً فیلتر مورد نظر را انتخاب کنید:", filtersMenu)
	}
}

// Example of a filter setting handler
func (b *Bot) handlePriceRange() tele.HandlerFunc {
	return func(c tele.Context) error {
		// Create price range menu
		menu := &tele.ReplyMarkup{ResizeKeyboard: true}

		btnLow := menu.Text("تا 1 میلیارد")
		btnMed := menu.Text("1 تا 3 میلیارد")
		btnHigh := menu.Text("بالای 3 میلیارد")
		btnCustom := menu.Text("محدوده دلخواه")
		btnBack := menu.Text("🔙 بازگشت")

		menu.Reply(
			menu.Row(btnLow, btnMed),
			menu.Row(btnHigh, btnCustom),
			menu.Row(btnBack),
		)

		return c.Send("لطفاً محدوده قیمت را انتخاب کنید:", menu)
	}
}

// handleCustomPriceRange handles custom price range input
func (b *Bot) handleCustomPriceRange() tele.HandlerFunc {
	return func(c tele.Context) error {
		// TODO: Implement price range service integration
		return c.Send("لطفاً حداقل و حداکثر قیمت را به صورت زیر وارد کنید:\n" +
			"مثال: 1000000000-3000000000")
	}
}

// Example of how to handle property listings
func (b *Bot) handlePropertyListing(property interface{}) tele.HandlerFunc {
	return func(c tele.Context) error {
		// Create property view menu
		menu := &tele.ReplyMarkup{ResizeKeyboard: true}

		btnBookmark := menu.Text("⭐️ افزودن به علاقه‌مندی‌ها")
		btnShare := menu.Text("📤 اشتراک‌گذاری")
		btnWatch := menu.Text("👀 پیگیری تغییرات")
		btnBack := menu.Text("🔙 بازگشت")

		menu.Reply(
			menu.Row(btnBookmark, btnShare),
			menu.Row(btnWatch),
			menu.Row(btnBack),
		)

		// TODO: Implement property service integration
		return c.Send("اطلاعات ملک:\n"+
			"[اینجا اطلاعات ملک نمایش داده می‌شود]", menu)
	}
}

// Middleware example for premium features
func (b *Bot) premiumOnly(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		// TODO: Implement user service integration to check premium status
		isPremium := false
		if !isPremium {
			return c.Send("این قابلیت فقط برای کاربران ویژه در دسترس است.")
		}
		return next(c)
	}
}

// Middleware for rate limiting
func (b *Bot) rateLimit(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		// TODO: Implement rate limiting service
		return next(c)
	}
}
