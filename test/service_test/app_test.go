package servicetest

import (
	"flag"
	"log"
	"os"
	"testing"

	"creepy/internal/service"
	"creepy/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey string

const WhoCallsKey contextKey = "WhoCalls"

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

func NewDevelopLogger() *zap.Logger {
	// Create a new logger with human-friendly, colorized output
	cfg := zap.NewDevelopmentConfig()

	// Adjust log level if necessary (change to zap.InfoLevel or zap.DebugLevel)
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	// Customize the console encoder settings for better readability
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	cfg.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder

	// Create the logger with the configuration
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

func TestAppService(t *testing.T) {
	cfg := readConfig()

	app, err := service.NewAppContainer(cfg)
	if err != nil {
		app.Cfg.Logger.Error(err.Error())
	}
	app.Cfg.Logger = NewDevelopLogger()
	app.Cfg.Logger.Info("Application starting", zap.String("version", "1.0.0"))

	UserServiceTest(app)
	FilterServiceTest(app)

	app.Cfg.Logger.Info("Application shutting down")
}