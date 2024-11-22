//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"mzhn/event-service/internal/config"
	"mzhn/event-service/internal/services/authservice"
	"mzhn/event-service/internal/services/eventservice"
	"mzhn/event-service/internal/storage/pg/eventstorage"
	esgrpc "mzhn/event-service/internal/transport/grpc"
	"mzhn/event-service/internal/transport/http"
	"mzhn/event-service/pb/authpb"

	"github.com/google/wire"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New() (*App, func(), error) {
	panic(wire.Build(
		newApp,
		_servers,

		eventservice.New,
		wire.Bind(new(eventservice.EventProvider), new(*eventstorage.Storage)),
		wire.Bind(new(eventservice.EventLoader), new(*eventstorage.Storage)),

		authservice.New,

		eventstorage.New,

		_authpb,
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

func _authpb(cfg *config.Config) (authpb.AuthClient, error) {
	addr := cfg.AuthService.ConnectionString()
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return authpb.NewAuthClient(conn), nil
}

func _servers(cfg *config.Config, es *eventservice.Service, as *authservice.Service) []Server {
	servers := make([]Server, 0, 2)

	if cfg.Http.Enabled {
		servers = append(servers, http.New(cfg, as, es))
	}

	if cfg.Grpc.Enabled {
		servers = append(servers, esgrpc.New(cfg, es))
	}

	return servers
}
