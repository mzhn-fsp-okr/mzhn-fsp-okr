package subscribersapi

import (
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/pkg/sl"
)

type Api struct {
	l   *slog.Logger
	cfg *config.Config
}

func New(cfg *config.Config) *Api {
	return &Api{
		l:   slog.With(sl.Module("subscribers-api")),
		cfg: cfg,
	}
}
