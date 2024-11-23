package domain

import "context"

type CronService interface {
	NotifyUsers(ctx context.Context, daysLeft uint32) error
}
