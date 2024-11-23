package authservice

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/notification-cron/internal/config"
	"mzhn/notification-cron/internal/domain"
	"mzhn/notification-cron/pb/authpb"
	"mzhn/notification-cron/pkg/sl"

	"github.com/samber/lo"
)

type Service struct {
	l    *slog.Logger
	cfg  *config.Config
	auth authpb.AuthClient
}

func New(cfg *config.Config, auth authpb.AuthClient) *Service {
	return &Service{
		l:    slog.With(sl.Module("AuthService")),
		cfg:  cfg,
		auth: auth,
	}
}

func (s *Service) Authenticate(ctx context.Context, token string, roles ...domain.Role) (bool, error) {
	fn := "authservice.Authenticate"
	log := s.l.With(sl.Method(fn))

	auth, err := s.auth.Authenticate(ctx, &authpb.AuthenticateRequest{
		AccessToken: token,
		Roles: lo.Map(roles, func(r domain.Role, _ int) authpb.Role {
			return authpb.Role(r)
		}),
	})
	if err != nil {
		log.Error("failed to authenticate", sl.Err(err))
		return false, err
	}

	if !auth.Approved {
		return false, fmt.Errorf("%s: %w", fn, domain.ErrUnathorized)
	}

	return auth.Approved, nil
}

func (s *Service) Profile(ctx context.Context, token string) (*authpb.UserInfo, error) {
	fn := "authservice.Profile"
	log := s.l.With(sl.Method(fn))

	profile, err := s.auth.Profile(ctx, &authpb.ProfileRequest{
		AccessToken: token,
	})
	if err != nil {
		log.Error("failed to get profile", sl.Err(err))
		return nil, err
	}

	return profile.User, nil
}
