package subscriptions_storage

import (
	"fmt"
	"mzhn/subscriptions-service/internal/domain"
	"mzhn/subscriptions-service/pb/sspb"
	"time"
)

func (s *Storage) NotifyUser(userId string, daysLeft sspb.DaysLeft) error {
	f := ""
	switch daysLeft {
	case sspb.DaysLeft_Month:
		f = "moth_notified_at"
	case sspb.DaysLeft_Week:
		f = "week_notified_at"
	case sspb.DaysLeft_ThreeDays:
		f = "three_days_notified_at"
	default:
		return fmt.Errorf("wrong days left enum")
	}

	if result := s.db.Model(&domain.EventSubscription{}).Where("user_id = ?", userId).Update(f, time.Now()); result.Error != nil {
		s.logger.Error("cannot notify user (DB)")
		return result.Error
	}

	return nil
}
