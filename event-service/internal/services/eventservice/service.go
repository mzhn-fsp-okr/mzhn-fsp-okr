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
	List(ctx context.Context, chEvents chan<- domain.EventInfo, filters model.EventsFilters) error
}

type EventLoader interface {
	Load(ctx context.Context, in *domain.EventLoadInfo) (string, error)
}

type NotificationPublisher interface {
	Notification(ctx context.Context, in *domain.EventInfo) error
}

type Service struct {
	l                     *slog.Logger
	cfg                   *config.Config
	ep                    EventProvider
	el                    EventLoader
	notificationPublisher NotificationPublisher
}

func New(cfg *config.Config, ep EventProvider, el EventLoader, np NotificationPublisher) *Service {
	return &Service{
		l:                     slog.With(sl.Module("EventService")),
		cfg:                   cfg,
		ep:                    ep,
		el:                    el,
		notificationPublisher: np,
	}
}

func (s *Service) Load(ctx context.Context, in *domain.EventLoadInfo) (string, error) {
	fn := "EventService.Load"
	log := s.l.With(sl.Method(fn))

	log.Info("loading an event", slog.String("ekp", in.EkpId))
	eid, err := s.el.Load(ctx, in)
	if err != nil {
		log.Error("failed to load event", sl.Err(err))
		return "", err
	}

	event := domain.EventInfo{
		Id:                      eid,
		EkpId:                   in.EkpId,
		SportType:               in.SportType,
		SportSubtype:            in.SportSubtype,
		Name:                    in.Name,
		Description:             in.Description,
		Dates:                   in.Dates,
		Location:                in.Location,
		Participants:            in.Participants,
		ParticipantRequirements: in.ParticipantRequirements,
	}

	log.Info("notificating about a new event", slog.String("eventId", eid))
	if err := s.notificationPublisher.Notification(ctx, &event); err != nil {
		log.Error("failed to publish notification", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	return eid, nil
}

func (s *Service) List(ctx context.Context, chEvents chan<- domain.EventInfo, filters model.EventsFilters) error {
	fn := "EventService.List"
	log := s.l.With(sl.Method(fn))

	err := s.ep.List(ctx, chEvents, filters)
	if err != nil {
		log.Error("failed to list events", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}