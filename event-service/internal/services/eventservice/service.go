package eventservice

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/event-service/internal/config"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/internal/storage/model"
	"mzhn/event-service/pkg/sl"
)

type EventProvider interface {
	Find(ctx context.Context, id string) (*domain.EventInfo, error)
	List(ctx context.Context, chEvents chan<- domain.EventInfo, filters ...model.EventsFilters) error
	Count(ctx context.Context, filters model.EventsFilters) (int64, error)
}

type EventManager interface {
	StaleAll(ctx context.Context) error
}

type EventLoader interface {
	Load(ctx context.Context, in *domain.EventLoadInfo) (*domain.EventInfo, error)
}

type NotificationPublisher interface {
	Notification(ctx context.Context, in *domain.EventInfo) error
}

type Service struct {
	l                     *slog.Logger
	cfg                   *config.Config
	ep                    EventProvider
	el                    EventLoader
	em                    EventManager
	notificationPublisher NotificationPublisher
}

func New(cfg *config.Config, ep EventProvider, el EventLoader, np NotificationPublisher, em EventManager) *Service {
	return &Service{
		l:                     slog.With(sl.Module("EventService")),
		cfg:                   cfg,
		ep:                    ep,
		el:                    el,
		em:                    em,
		notificationPublisher: np,
	}
}

func (s *Service) Load(ctx context.Context, in *domain.EventLoadInfo) (string, error) {
	fn := "EventService.Load"
	log := s.l.With(sl.Method(fn))

	log.Info("loading an event", slog.String("ekp", in.EkpId))
	event, err := s.el.Load(ctx, in)
	if err != nil {
		log.Error("failed to load event", sl.Err(err))
		return "", err
	}

	log.Info("notificating about a new event", slog.String("eventId", event.Id))
	if err := s.notificationPublisher.Notification(ctx, event); err != nil {
		log.Error("failed to publish notification", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	return event.Id, nil
}

func (s *Service) List(ctx context.Context, chEvents chan<- domain.EventInfo, filters ...model.EventsFilters) error {

	fn := "EventService.List"
	log := s.l.With(sl.Method(fn))

	err := s.ep.List(ctx, chEvents, filters...)
	if err != nil {
		log.Error("failed to list events", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Service) Count(ctx context.Context, filters model.EventsFilters) (int64, error) {
	fn := "EventService.Count"
	log := s.l.With(sl.Method(fn))

	count, err := s.ep.Count(ctx, filters)
	if err != nil {
		log.Error("failed to count events", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", fn, err)
	}

	return count, nil
}

func (s *Service) Stale(ctx context.Context) error {
	fn := "EventService.Stale"
	log := s.l.With(sl.Method(fn))

	if err := s.em.StaleAll(ctx); err != nil {
		log.Error("failed to stale events", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Service) Find(ctx context.Context, eid string) (*domain.EventInfo, error) {
	fn := "EventService.Find"
	log := s.l.With(sl.Method(fn))

	event, err := s.ep.Find(ctx, eid)
	if err != nil {
		log.Error("failed to find event", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return event, nil
}
