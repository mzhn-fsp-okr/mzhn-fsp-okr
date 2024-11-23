package subscriptions_storage

import (
	"mzhn/subscriptions-service/internal/domain"

	"github.com/samber/lo"
)

func (s *Storage) GetUserEventsId(userId string) ([]string, error) {
	users := []*domain.EventSubscription{}
	result := s.db.Select("event_id").Where("user_id = ?", userId).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return lo.Map(users, func(item *domain.EventSubscription, _ int) string {
		return item.EventId
	}), nil
}

func (s *Storage) GetUsersSubscribedToEvent(eventId string) ([]string, error) {
	events := []*domain.EventSubscription{}
	result := s.db.Select("user_id").Where("event_id = ?", eventId).Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}

	return lo.Map(events, func(item *domain.EventSubscription, _ int) string {
		return item.UserId
	}), nil
}

func (s *Storage) GetUsersSubscribedToSport(sportId string) ([]string, error) {
	events := []*domain.SportSubscription{}
	result := s.db.Select("user_id").Where("sport_id = ?", sportId).Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}

	return lo.Map(events, func(item *domain.SportSubscription, _ int) string {
		return item.UserId
	}), nil
}
