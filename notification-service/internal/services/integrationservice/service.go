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
	Create(ctx context.Context, userId string) error
	Save(context.Context, *domain.SetIntegrations) error
}
type IntegrationsProvider interface {
	Find(ctx context.Context, userId string) (*domain.Integrations, error)
}

type Service struct {
	l   *slog.Logger
	cfg *config.Config
	is  IntegrationsSaver
	ip  IntegrationsProvider
}

func New(cfg *config.Config, is IntegrationsSaver, ip IntegrationsProvider) *Service {
	return &Service{
		l:   slog.With(sl.Module("integration-service")),
		cfg: cfg,
		is:  is,
		ip:  ip,
	}
}

func (s *Service) Find(ctx context.Context, userId string) (*domain.Integrations, error) {
	fn := "integrationservice.Find"
	log := s.l.With(sl.Module(fn))

	i, err := s.ip.Find(ctx, userId)
	if err != nil {
		log.Error("failed to find integrations", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	if i == nil {
		if err := s.is.Create(ctx, userId); err != nil {
			return nil, err
		}

		newI, err := s.ip.Find(ctx, userId)
		if err != nil {
			return nil, err
		}

		i = newI
	}

	return i, nil
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
