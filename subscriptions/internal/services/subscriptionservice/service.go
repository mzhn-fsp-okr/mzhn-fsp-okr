package subscriptionservice

import (
	"log/slog"
	"mzhn/subscriptions-service/internal/domain"
	"mzhn/subscriptions-service/pkg/sl"
)

type Service struct {
	l       *slog.Logger
	storage domain.SubscriptionsStorage
}

func New(storage domain.SubscriptionsStorage) *Service {
	return &Service{
		l:       slog.With(sl.Module(("SubscriptionsService"))),
		storage: storage,
	}
}
func (s *Service) SubscribeToSport(dto *domain.SportSubscription) (*domain.SportSubscription, error) {
	log := s.l.With(sl.Method("SubscriptionsService.SubscribeToSport"))

	log.Debug("creating sport subscription", slog.Any("userId", dto.UserId), slog.Any("sportId", dto.SportId))
	return s.storage.CreateSport(dto)
}

func (s *Service) SubscribeToEvent(dto *domain.EventSubscription) (*domain.EventSubscription, error) {
	log := s.l.With(sl.Method("SubscriptionsService.SubscribeToEvent"))

	log.Debug("creating event subscription", slog.Any("userId", dto.UserId), slog.Any("eventId", dto.EventId))
	return s.storage.CreateEvent(dto)
}

func (s *Service) UnsubscribeFromSport(dto *domain.SportSubscription) error {
	log := s.l.With(sl.Method("SubscriptionsService.UnsubscribeFromSport"))

	log.Debug("unsubscribe from sport", slog.Any("userId", dto.UserId), slog.Any("sportId", dto.SportId))
	return s.storage.DeleteSport(dto)
}

func (s *Service) UnsubscribeFromEvent(dto *domain.EventSubscription) error {
	log := s.l.With(sl.Method("SubscriptionsService.UnsubscribeFromEvent"))

	log.Debug("unsubscribe from event", slog.Any("userId", dto.UserId), slog.Any("eventId", dto.EventId))
	return s.storage.DeleteEvent(dto)

}
