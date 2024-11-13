.
├── .env                    # تنظیمات محیطی (API tokens, database credentials)
│
├── cmd/
│   └── main.go           # نقطه شروع برنامه - راه‌اندازی همزمان کرالر و بات
│
├── internal/
│   ├── bot/
│   │   ├── bot.go        # راه‌اندازی و پیکربندی بات تلگرام
│   │
│   ├── crawler/
│   │   ├── crawler.go    # واسط و ساختارهای مشترک کرالرها
│   │   ├── divar.go      # پیاده‌سازی کرالر دیوار
│   │   ├── sheypoor.go   # پیاده‌سازی کرالر شیپور
│   │
│   ├── models/
│   │   ├── property.go    # مدل آگهی‌های ملک
│   │   ├── user.go        # مدل کاربران
│   │   ├── filter.go      # مدل فیلترها
│   │   └── bookmark.go    # مدل علاقه‌مندی‌ها
│   │   └── user_search_history.go
│   │
│   ├── storage/
│   │   ├── mysql/
│   │   │   ├── property.go  # پیاده‌سازی MySQL برای آگهی‌ها
│   │   │   ├── user.go      # پیاده‌سازی MySQL برای کاربران
│   │   │   └── filter.go    # پیاده‌سازی MySQL برای فیلترها
│   │   └── repository.go   # تعریف interface‌های repository
│   │
│   ├── service/
│   │   ├── property.go    # سرویس مدیریت آگهی‌ها
│   │   ├── user.go        # سرویس مدیریت کاربران
│   │   ├── filter.go      # سرویس مدیریت فیلترها
│   │   └── analytics.go   # سرویس تحلیل و آمار
│
├── pkg/
│   ├── config/
│   │   └── config.go     # تنظیمات برنامه
│   │
│   └── utils/
│       ├── logger.go     # ابزار لاگ گیری
│       ├── uuid.go        # تولید شناسه‌های یکتا
│       ├── monitoring.go     # متریک‌های عملکردی کرالر (مانیتورینگ با پشتیبانی از Prometheus)
│       └── errors.go      # مدیریت خطاها
│
└── go.mod
