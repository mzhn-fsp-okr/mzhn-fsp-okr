package authservice

import (
	"context"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/pkg/sl"
)

type ProfileProvider interface {
	Profile(ctx context.Context, token string) (*domain.User, error)
}

type Service struct {
	l   *slog.Logger
	cfg *config.Config
	pp  ProfileProvider
}

func New(cfg *config.Config, pp ProfileProvider) *Service {
	return &Service{
		l:   slog.With(sl.Module("AuthService")),
		cfg: cfg,
		pp:  pp,
	}
}

func (s *Service) Profile(ctx context.Context, token string) (*domain.User, error) {
	fn := "service.Profile"
	log := s.l.With(sl.Method(fn))

	user, err := s.pp.Profile(ctx, token)
	if err != nil {
		log.Error("failed to get profile", sl.Err(err))
		return nil, err
	}

	return user, nil
}
