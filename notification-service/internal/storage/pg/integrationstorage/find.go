package integrationstorage

import (
	"context"
	"mzhn/notification-service/internal/domain"
)

func (s *Storage) Find(ctx context.Context, userId string) (*domain.Integrations, error) {
	return s.find(ctx, userId)
}
