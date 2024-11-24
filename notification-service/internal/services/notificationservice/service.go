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
	SportSubs(ctx context.Context, sportId string) ([]string, error)
}

type UserProvider interface {
	Find(ctx context.Context, id string) (*domain.User, error)
}

type IntegrationProvider interface {
	Find(ctx context.Context, userId string) (*domain.Integrations, error)
}

type Notificator interface {
	SendTelegram(ctx context.Context, event map[string]any) error
	SendMail(ctx context.Context, event map[string]any) error
}

type EventProvider interface {
	Event(ctx context.Context, eventId string) (*domain.EventInfo, error)
}
type SportProvider interface {
	Sport(ctx context.Context, sportId string) (*domain.SportTypeWithSubtypes, error)
}

type Service struct {
	l             *slog.Logger
	cfg           *config.Config
	sp            SubscribersProvider
	ip            IntegrationProvider
	up            UserProvider
	ep            EventProvider
	sportProvider SportProvider
	notificator   Notificator
}

func New(
	cfg *config.Config,
	sp SubscribersProvider,
	ip IntegrationProvider,
	notificator Notificator,
	up UserProvider,
	ep EventProvider,
	sportProvider SportProvider,
) *Service {
	return &Service{
		l:             slog.With(sl.Module("notification-service")),
		cfg:           cfg,
		sp:            sp,
		ip:            ip,
		up:            up,
		ep:            ep,
		notificator:   notificator,
		sportProvider: sportProvider,
	}
}
