package servicetest

import (
	"context"
	"flag"
	"log"
	"os"
	"testing"

	"creepy/internal/service"
	"creepy/pkg/config"
	"creepy/test/data"

	"go.uber.org/zap"
)

type contextKey string

const isAdminKey contextKey = "isAdmin"

func StoreIsAdmin(ctx context.Context, value bool) context.Context {
    return context.WithValue(ctx, isAdminKey, value)
}

func GetIsAdmin(ctx context.Context) (bool, bool) {
    value, ok := ctx.Value(isAdminKey).(bool)
    return value, ok
}

var configPath = flag.String("config", ".env", "path to the configuration file")

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
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

func TestAppService(t *testing.T) {
	cfg := readConfig()

	app, err := service.NewAppContainer(cfg)
	if err != nil {
		app.Cfg.Logger.Error(err.Error())
	}
	cfg.Logger.Info("Application starting", zap.String("version", "1.0.0"))

	property := data.Properties[0]
	ctx := context.Background()
	ctx = StoreIsAdmin(ctx, true)

	err = app.PropertyService().CreatePropertyByAdmin(ctx, &property)
	if err != nil {
		app.Cfg.Logger.Fatal(err.Error())
	}

	app.Cfg.Logger.Info("Application shutting down")
}