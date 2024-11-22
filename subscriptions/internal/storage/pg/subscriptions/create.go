package subscriptions_storage

import (
	"mzhn/subscriptions-service/internal/domain"
)

func (s *Storage) CreateSport(sportSubscription *domain.SportSubscription) (*domain.SportSubscription, error) {
	result := s.db.Create(sportSubscription)
	if result.Error != nil {
		return nil, result.Error
	}

	return sportSubscription, nil
}

func (s *Storage) CreateEvent(eventSubscription *domain.EventSubscription) (*domain.EventSubscription, error) {
	result := s.db.Create(eventSubscription)
	if result.Error != nil {
		return nil, result.Error
	}

	return eventSubscription, nil
}
