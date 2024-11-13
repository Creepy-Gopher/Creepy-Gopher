package main

import (
	"creepy/internal/service"
	"creepy/pkg/config"
	"creepy/internal/models"
	"flag"
	"log"

	//"fmt"
	"os"
	"sync"
	"context"
	"go.uber.org/zap"
)

var configPath = flag.String("config", ".env", "path to the configuration file")

func main() {
	// Initialize the database connection
	cfg := readConfig()
	//config.Set(cfg)
	log.Println(cfg)
	//log.Fatal("out")
	// db, err := postgis.NewPostgresGormConnection(cfg.DB)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// if err := postgis.AddExtension(db); err != nil {
	//     log.Fatal(err)
	// }
	// if err := postgis.Migrate(db); err != nil {
	//     log.Fatal(err)
	// }

	// Initialize repositories
	//propertyRepo := postgis.NewPropertyRepository(db)
	// Initialize other repositories...

	// Initialize services with repository interfaces
	// propertyService := service.NewPropertyService(propertyRepo)
	//_ = propertyService
	// Initialize other services...

	// Pass services to your handlers, bots, etc.
	// ...

	// Use the logger
	app, err := service.NewAppContainer(cfg)
	if err != nil {
		app.Cfg.Logger.Error(err.Error())
	}
	cfg.Logger.Info("Application starting", zap.String("version", "1.0.0"))

	//sample create property usage:
	sampleProperty := &models.Property{
		Title:        "new prop",
		Description:  "a good view",
		BuyPrice:     450000,
		RentPrice:    45666666,
		RentPriceMin: 898888,
		RentPriceMax: 7878787,
		RahnPriceMin: 787878,
		RahnPriceMax: 7878787,
		Area:         455656,
		Rooms:        3,
		DealingType:  "ddd",
		Type:         "ddd",
		City:         "ddd",
		District:     "ffff",
		Address:      "ddd",
		BuildYear:    1398,
		Floor:        2,
		HasElevator:  false,
		HasStorage:   false,
		HasParking:   false,
		Latitude:     100,
		Longitude:    1000,
		Source:       "divar",
		URL:          "adjjajf;",
		Image:        "kjkjskskhs",
	}
	ctx := context.Background()
	err = app.PropertyService().CreateProperty(ctx, sampleProperty)

	if err!=nil{
		log.Fatal(err)
	}
	log.Fatal("out")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		runCrawler(cfg, app)
	}()

	go func() {
		defer wg.Done()
		runTelegramBot(cfg, app)
	}()

	wg.Wait()

	cfg.Logger.Info("Application shutting down")
}

func readConfig() config.Config {
	flag.Parse()

	if cfgPathEnv := os.Getenv("APP_CONFIG_PATH"); len(cfgPathEnv) > 0 {
		*configPath = cfgPathEnv
	}

	if len(*configPath) == 0 {
		log.Fatal("configuration file not found")
	}

	return config.NewConfig()
}

func runCrawler(cfg config.Config, app *service.AppContainer) {
	cfg.Logger.Info("Starting property crawler")
	// Implement your crawling logic here
}

func runTelegramBot(cfg config.Config, app *service.AppContainer) {
	cfg.Logger.Info("Starting Telegram bot")
	// Implement your Telegram bot logic here
}

/*
cfg.Logger.Info("User logged in",
    zap.String("username", "john_doe"),
    zap.Int("user_id", 12345))
*/
