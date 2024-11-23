package config

import (
	"fmt"
	"log/slog"
	"os"

	"mzhn/notification-cron/pkg/prettyslog"
	"mzhn/notification-cron/pkg/sl"

	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	Name    string `env:"APP_NAME" env-required:"true"`
	Version string `env:"APP_VERSION" env-required:"true"`
}

type AuthService struct {
	Host string `env:"AUTH_SERVICE_HOST" env-required:"true"`
	Port int    `env:"AUTH_SERVICE_PORT" env-required:"true"`
}

func (as *AuthService) ConnectionString() string {
	return fmt.Sprintf("%s:%d", as.Host, as.Port)
}

type Cron struct {
	Minutes int `env:"INTERVAL" env-default:"20"` // in minutes
}

type Config struct {
	Env         string `env:"ENV" env-default:"local"`
	App         App
	AuthService AuthService
	Cron        Cron
}

func New() *Config {
	config := new(Config)

	if err := cleanenv.ReadEnv(config); err != nil {
		slog.Error("error when reading env", sl.Err(err))
		header := fmt.Sprintf("%s - %s", os.Getenv("APP_NAME"), os.Getenv("APP_VERSION"))

		usage := cleanenv.FUsage(os.Stdout, config, &header)
		usage()

		os.Exit(-1)
	}

	setupLogger(config)

	slog.Debug("config", slog.Any("c", config))
	return config
}

func setupLogger(cfg *Config) {
	var log *slog.Logger

	switch cfg.Env {
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	default:
		log = slog.New(prettyslog.NewPrettyHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	slog.SetDefault(log)
}
