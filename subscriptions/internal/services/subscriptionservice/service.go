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
func (s *Service) CreateSubscriptionToSport(dto *domain.SportSubscription) (*domain.SportSubscription, error) {
	log := s.l.With(sl.Method("SubscriptionsService.CreateSubscriptionToSport"))

	log.Debug("creating sport subscription")
	return s.storage.CreateSport(dto)
}

func (s *Service) CreateSubscriptionToEvent(dto *domain.EventSubscription) (*domain.EventSubscription, error) {
	log := s.l.With(sl.Method("SubscriptionsService.CreateSubscriptionToEvent"))

	log.Debug("creating event subscription")
	return s.storage.CreateEvent(dto)
}
