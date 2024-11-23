package subscriptions_storage

import (
	"log/slog"
	"mzhn/subscriptions-service/pkg/sl"

	"gorm.io/gorm"
)

type Storage struct {
	logger *slog.Logger
	db     *gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{
		logger: slog.With(sl.Module("SubscriptionsStorage")),
		db:     db,
	}
}
