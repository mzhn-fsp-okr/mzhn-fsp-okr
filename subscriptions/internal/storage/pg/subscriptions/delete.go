package subscriptions_storage

import "mzhn/subscriptions-service/internal/domain"

func (s *Storage) DeleteEvent(eventSubscription *domain.EventSubscription) error {
	result := s.db.Unscoped().Where(eventSubscription).Delete(&domain.EventSubscription{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Storage) DeleteSport(sportSubscription *domain.SportSubscription) error {
	result := s.db.Unscoped().Where(sportSubscription).Delete(&domain.SportSubscription{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
