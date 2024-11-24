package sportstorage

import (
	"log/slog"
	"mzhn/event-service/internal/config"
	"mzhn/event-service/pkg/sl"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	l    *slog.Logger
	cfg  *config.Config
	pool *pgxpool.Pool
}

func New(cfg *config.Config, pool *pgxpool.Pool) *Storage {
	return &Storage{
		l:    slog.With(sl.Module("sport-storage")),
		cfg:  cfg,
		pool: pool,
	}
}
