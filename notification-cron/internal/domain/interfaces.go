package domain

import "context"

type CronService interface {
	NotifyUsers(ctx context.Context, daysLeft uint32) error
}

type UpcomingNotificationPublisher interface {
	NotifyAboutUpcomingEvent(ctx context.Context, userId string, eventId string, daysLeft uint32) error
}
