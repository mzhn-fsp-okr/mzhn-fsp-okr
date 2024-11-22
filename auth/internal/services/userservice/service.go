package userservice

import (
	"context"
	"log/slog"

	"mzhn/auth/internal/config"
	"mzhn/auth/internal/domain/entity"
	"mzhn/auth/pkg/sl"
)

//go:generate go run github.com/vektra/mockery/v2@v2.46.0 --name=UserProvider
type UserProvider interface {
	Find(ctx context.Context, slug string) (*entity.User, error)
	Count(ctx context.Context) (int64, error)
}

type Service struct {
	userProvider UserProvider
	cfg          *config.Config
	logger       *slog.Logger
}

func New(uprovider UserProvider, cfg *config.Config) *Service {
	svc := &Service{
		cfg:          cfg,
		userProvider: uprovider,
		logger:       slog.With(sl.Module("UserService")),
	}

	return svc
}
