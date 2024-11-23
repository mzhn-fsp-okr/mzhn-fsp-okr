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
	Minutes int `env:"INTERVAL" env-default:"1"` // in minutes
}

type Amqp struct {
	Host string `env:"AMQP_HOST" env-required:"true"`
	Port int    `env:"AMQP_PORT" env-required:"true"`
	User string `env:"AMQP_USER" env-required:"true"`
	Pass string `env:"AMQP_PASS" env-required:"true"`

	NotificationsExchange string `env:"AMQP_NOTIFICATIONS_EXCHANGE" env-required:"true"`
	UpcomingQueue         string `env:"AMQP_UPCOMING_EVENTS_QUEUE" env-required:"true"`
}

func (am *Amqp) ConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d", am.User, am.Pass, am.Host, am.Port)
}

type EventService struct {
	Host string `env:"EVENT_SERVICE_HOST" env-required:"true"`
	Port int    `env:"EVENT_SERVICE_PORT" env-required:"true"`
}

func (as *EventService) ConnectionString() string {
	return fmt.Sprintf("%s:%d", as.Host, as.Port)
}

type SubscriptionsService struct {
	Host string `env:"SUBSCRIPTIONS_SERVICE_HOST" env-required:"true"`
	Port int    `env:"SUBSCRIPTIONS_SERVICE_PORT" env-required:"true"`
}

func (as *SubscriptionsService) ConnectionString() string {
	return fmt.Sprintf("%s:%d", as.Host, as.Port)
}

type Config struct {
	Env                  string `env:"ENV" env-default:"local"`
	App                  App
	AuthService          AuthService
	EventService         EventService
	SubscriptionsService SubscriptionsService
	Cron                 Cron
	Amqp                 Amqp
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
