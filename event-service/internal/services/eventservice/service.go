package eventservice

import (
	"context"
	"log/slog"
	"mzhn/event-service/internal/config"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/pkg/sl"
)

type EventProvider interface {
	Find(ctx context.Context, id string) (*domain.EventInfo, error)
	Load(ctx context.Context, in *domain.EventLoadInfo) (string, error)
}

type Service struct {
	l   *slog.Logger
	cfg *config.Config
	ep  EventProvider
}

func New(cfg *config.Config, ep EventProvider) *Service {
	return &Service{
		l:   slog.With(sl.Module("EventService")),
		cfg: cfg,
		ep:  ep,
	}
}

func (s *Service) Load(ctx context.Context, in *domain.EventLoadInfo) (string, error) {
	fn := "EventService.Load"
	log := s.l.With(sl.Method(fn))

	eid, err := s.ep.Load(ctx, in)
	if err != nil {
		log.Error("failed to load event", sl.Err(err))
		return "", err
	}

	return eid, nil
}
