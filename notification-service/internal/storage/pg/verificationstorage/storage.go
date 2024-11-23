package verificationstorage

import (
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/pkg/sl"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	l    *slog.Logger
	cfg  *config.Config
	pool *pgxpool.Pool
}

func New(cfg *config.Config, pool *pgxpool.Pool) *Storage {
	return &Storage{
		l:    slog.With(sl.Module("verification-storage")),
		cfg:  cfg,
		pool: pool,
	}
}
