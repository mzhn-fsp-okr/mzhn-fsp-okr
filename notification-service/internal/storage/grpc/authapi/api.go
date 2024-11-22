package authapi

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/pb/authpb"
	"mzhn/notification-service/pkg/sl"
)

type Api struct {
	l      *slog.Logger
	cfg    *config.Config
	client authpb.AuthClient
}

func New(cfg *config.Config, client authpb.AuthClient) *Api {
	return &Api{
		l:      slog.With(sl.Module("auth api")),
		cfg:    cfg,
		client: client,
	}
}

func (a *Api) Profile(ctx context.Context, token string) (*domain.User, error) {

	fn := "authapi.Profile"
	log := a.l.With(slog.String("fn", fn))

	response, err := a.client.Profile(ctx, &authpb.ProfileRequest{AccessToken: token})
	if err != nil {
		log.Error("failed to get profile", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	log.Debug("got profile", slog.Any("response", response))

	return &domain.User{
		Id:    response.User.Id,
		Email: response.User.Email,
	}, nil
}
