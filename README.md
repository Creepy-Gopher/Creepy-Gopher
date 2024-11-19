# 🏠 Creepy-Gopher

**Creepy-Gopher** is a comprehensive real estate management application developed in Go (Golang). It integrates a Telegram bot and a web crawler to help users search, filter, and manage property listings efficiently.

## 🚀 Features

### 🤖 Telegram Bot

- **Property Management**
  - Add new property listings
  - View your property listings
- **Filters**
  - Create, update, view, and delete search filters
  - Use filters for personalized property searches
- **User Management** (Admin & Super Admin)
  - Super Admin can create Admin users by their Telegram ID
  - Admins can manage users and properties

### 🕷️ Integrated Web Crawler

- Automatically fetches property listings from external sources
- Updates the database every 30 minutes

## 🏗️ Project Structure

```plaintext
Creepy-Gopher/
├── cmd/                 # Main application entry points
│   └── main.go         # Starts the crawler and bot
├── internal/           # Application core logic
│   ├── bot/           # Telegram bot implementation
│   ├── service/       # Business logic and services
│   ├── models/        # Data models and structures
│   └── storage/       # Data repositories
├── pkg/               # Shared packages
├── test/             # Tests
├── docker-compose.yml # Docker configuration
├── go.mod            # Go module configuration
└── README.md         # Project documentation
```

## 📋 Prerequisites

- Go (version 1.17 or higher)
- Docker and Docker Compose

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-username/Creepy-Gopher.git
   cd Creepy-Gopher
   ```

2. **Start the database**
   ```bash
   docker-compose up -d
   ```
   This will start PostgreSQL with PostGIS extension enabled.

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

4. **Configure the application**
   Create a `.env` file in the project root:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=youruser
   DB_PASSWORD=yourpassword
   DB_NAME=creepy_gopher
   TELEGRAM_BOT_TOKEN=your-telegram-bot-token
   ```

5. **Run the application**
   ```bash
   go run cmd/main.go
   ```

## 🎯 Usage

### Telegram Bot Commands

- `/start` - Initialize the bot
- Use the menu options for:
  - Property Management
  - Filter Management
  - Admin Features (for authorized users)

### Admin Features

- Super Admins can:
  - Create new admin users
  - Manage all properties and users
- Regular Admins can:
  - Manage properties
  - Handle user requests

### Web Crawler

The crawler operates automatically:
- Runs every 30 minutes
- Fetches new property listings
- Updates existing listings
- Maintains data freshness

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
