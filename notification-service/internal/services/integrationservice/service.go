package integrationservice

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/pkg/sl"
)

type IntegrationsSaver interface {
	Save(context.Context, *domain.SetIntegrations) error
}

type Service struct {
	l   *slog.Logger
	cfg *config.Config
	is  IntegrationsSaver
}

func New(cfg *config.Config, is IntegrationsSaver) *Service {
	return &Service{
		l:   slog.With(sl.Module("integration-service")),
		cfg: cfg,
		is:  is,
	}
}

func (s *Service) Save(ctx context.Context, req *domain.SetIntegrations) error {
	fn := "integrationservice.Save"
	log := s.l.With(sl.Module(fn))

	if err := s.is.Save(ctx, req); err != nil {
		log.Error("failed to save integrations", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
