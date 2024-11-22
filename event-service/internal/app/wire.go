package app

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"mzhn/event-service/internal/config"
	"mzhn/event-service/internal/services/eventservice"
	"mzhn/event-service/internal/storage/pg"
	"mzhn/event-service/internal/transport/grpc"
	"mzhn/event-service/internal/transport/http"

	"github.com/google/wire"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New() (*App, func(), error) {
	panic(wire.Build(
		newApp,
		_servers,

		eventservice.New,
		wire.Bind(new(eventservice.EventProvider), new(*pg.EventStorage)),

		pg.NewEventStorage,

		_pg,
		config.New,
	))
}

func _pg(cfg *config.Config) (*pgxpool.Pool, func(), error) {
	ctx := context.Background()
	cs := cfg.Pg.ConnectionString()
	pool, err := pgxpool.New(ctx, cs)
	if err != nil {
		return nil, nil, err
	}

	slog.Debug("connecting to database", slog.String("cs", cs))
	t := time.Now()
	if err := pool.Ping(ctx); err != nil {
		slog.Error("failed to connect to database", slog.String("err", err.Error()), slog.String("conn", cs))
		return nil, func() { pool.Close() }, err
	}
	slog.Info("connected to database", slog.String("ping", fmt.Sprintf("%2.fs", time.Since(t).Seconds())))

	return pool, func() { pool.Close() }, nil
}

func _servers(cfg *config.Config, es *eventservice.Service) []Server {
	servers := make([]Server, 0, 2)

	if cfg.Http.Enabled {
		servers = append(servers, http.New(cfg))
	}

	if cfg.Grpc.Enabled {
		servers = append(servers, grpc.New(cfg, es))
	}

	return servers
}
