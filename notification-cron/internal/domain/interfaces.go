package domain

import (
	"context"
	"mzhn/notification-cron/pb/sspb"
)

type CronService interface {
	NotifyUsers(ctx context.Context) error
}

type UpcomingNotificationPublisher interface {
	NotifyAboutUpcomingEvent(ctx context.Context, userId string, eventId string, daysLeft uint32, daysLeftEnum sspb.DaysLeft) error
}
