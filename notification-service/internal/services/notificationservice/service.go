package notificationservice

import (
	"context"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/pkg/sl"
)

type SubscribersProvider interface {
	EventSubscribers(ctx context.Context, eventId string) ([]string, error)
	// SportSubscribers(ctx context.Context, sportId string) ([]string, error)
}

type UserProvider interface {
	Find(ctx context.Context, id string) (*domain.User, error)
}

type IntegrationProvider interface {
	Find(ctx context.Context, userId string) (*domain.Integrations, error)
}

type Notificator interface {
	SendTelegram(ctx context.Context, username string, event *domain.EventInfo, eventType domain.EventType) error
	SendMail(ctx context.Context, mail string, event *domain.EventInfo, eventType domain.EventType) error
}

type Service struct {
	l           *slog.Logger
	cfg         *config.Config
	sp          SubscribersProvider
	ip          IntegrationProvider
	up          UserProvider
	notificator Notificator
}

func New(cfg *config.Config, sp SubscribersProvider, ip IntegrationProvider, notificator Notificator, up UserProvider) *Service {
	return &Service{
		l:           slog.With(sl.Module("notification-service")),
		cfg:         cfg,
		sp:          sp,
		ip:          ip,
		up:          up,
		notificator: notificator,
	}
}
