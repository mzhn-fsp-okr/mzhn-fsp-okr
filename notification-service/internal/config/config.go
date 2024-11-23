package config

import (
	"fmt"
	"log/slog"
	"os"

	"mzhn/notification-service/pkg/prettyslog"
	"mzhn/notification-service/pkg/sl"

	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	Name    string `env:"APP_NAME" env-default:"mzhn-notification-service"`
	Version string `env:"APP_VERSION" env-default:"@local"`
}

type Http struct {
	Host string `env:"HTTP_HOST" env-default:"0.0.0.0"`
	Port int    `env:"HTTP_PORT" env-default:"80"`
	Cors Cors
}

type Cors struct {
	AllowedOrigins string `env:"CORS_ALLOWED_ORIGINS" env-default:"localhost:3000"`
}

type Amqp struct {
	Host string `env:"AMQP_HOST" env-required:"true"`
	Port int    `env:"AMQP_PORT" env-required:"true"`
	User string `env:"AMQP_USER" env-required:"true"`
	Pass string `env:"AMQP_PASS" env-required:"true"`

	NotificationsExchange string `env:"AMQP_NOTIFICATIONS_EXCHANGE" env-default:"notifications"`
	NewEventsQueue        string `env:"AMQP_NEW_EVENTS_EVENTS_QUEUE" env-default:"new-events"`
	UpcomingEventsQueue   string `env:"AMQP_UPCOMING_EVENTS_QUEUE" env-default:"upcoming-events"`

	SubscriptionExchange string `env:"AMQP_SUBSCRIPTIONS_EXCHANGE" env-default:"subscriptions"`
	TelegramQueue        string `env:"AMQP_TELEGRAM_QUEUE" env-default:"telegram-queue"`
	MailQueue            string `env:"AMQP_MAIL_QUEUE" env-default:"mail-queue"`
}

func (am *Amqp) ConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d", am.User, am.Pass, am.Host, am.Port)
}

type Pg struct {
	Host string `env:"PG_HOST" env-required:"true"`
	Port int    `env:"PG_PORT" env-required:"true"`
	User string `env:"PG_USER" env-required:"true"`
	Pass string `env:"PG_PASS" env-required:"true"`
	Name string `env:"PG_NAME" env-required:"true"`
}

func (pg *Pg) ConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", pg.User, pg.Pass, pg.Host, pg.Port, pg.Name)
}

type AuthService struct {
	Protocol string `env:"AUTH_SERVICE_PROTOCOL" env-default:"grpc"`
	Host     string `env:"AUTH_SERVICE_HOST" env-required:"true"`
	Port     int    `env:"AUTH_SERVICE_PORT" env-required:"true"`
}

func (as *AuthService) ConnectionString() string {
	return fmt.Sprintf("%s://%s:%d", as.Protocol, as.Host, as.Port)
}

type SubscriptionService struct {
	Protocol string `env:"SUBSCRIPTION_SERVICE_PROTOCOL" env-default:"grpc"`
	Host     string `env:"SUBSCRIPTION_SERVICE_HOST" env-required:"true"`
	Port     int    `env:"SUBSCRIPTION_SERVICE_PORT" env-required:"true"`
}

func (s *SubscriptionService) ConnectionString() string {
	return fmt.Sprintf("%s://%s:%d", s.Protocol, s.Host, s.Port)
}

type Config struct {
	Env                 string `env:"ENV" env-default:"local"`
	App                 App
	Http                Http
	Amqp                Amqp
	Pg                  Pg
	AuthService         AuthService
	SubscriptionService SubscriptionService
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
