package subscriptions_storage

import (
	"mzhn/subscriptions-service/internal/domain"

	"github.com/samber/lo"
)

func (s *Storage) GetUserEventsId(userId string) ([]string, error) {
	users := []*domain.EventSubscription{}
	result := s.db.Select("user_id").Where("user_id = ?", userId).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return lo.Map(users, func(item *domain.EventSubscription, _ int) string {
		return item.EventId
	}), nil
}
