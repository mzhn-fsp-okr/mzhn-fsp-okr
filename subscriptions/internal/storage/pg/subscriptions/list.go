package subscriptions_storage

import (
	"fmt"
	"mzhn/subscriptions-service/internal/domain"
	"mzhn/subscriptions-service/pb/sspb"
	"mzhn/subscriptions-service/pkg/sl"

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

func (s *Storage) GetUsersFromEventByDaysLeft(eventId string, daysLeft sspb.DaysLeft) ([]string, error) {
	f := ""
	switch daysLeft {
	case sspb.DaysLeft_Month:
		f = "moth_notified_at"
	case sspb.DaysLeft_Week:
		f = "week_notified_at"
	case sspb.DaysLeft_ThreeDays:
		f = "three_days_notified_at"
	default:
		return nil, fmt.Errorf("wrong days left enum")
	}

	var events []domain.EventSubscription

	result := s.db.Select("user_id").Where(f+" IS NULL AND event_id = ?", eventId).Find(&events)
	if result.Error != nil {
		s.logger.Error("ERROR", sl.Err(result.Error))
		return nil, result.Error
	}

	return lo.Map(events, func(item domain.EventSubscription, _ int) string {
		return item.UserId
	}), nil
}
