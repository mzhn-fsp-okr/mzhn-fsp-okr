// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	redis2 "github.com/redis/go-redis/v9"
	"log/slog"
	"mzhn/auth/internal/config"
	"mzhn/auth/internal/services/authservice"
	"mzhn/auth/internal/services/userservice"
	"mzhn/auth/internal/storage/pg"
	"mzhn/auth/internal/storage/redis"
	"mzhn/auth/internal/transport/grpc"
	"mzhn/auth/internal/transport/http"
	"time"
)

import (
	_ "github.com/jackc/pgx/stdlib"
)

// Injectors from wire.go:

func New() (*App, func(), error) {
	configConfig := config.New()
	db, cleanup, err := _pg(configConfig)
	if err != nil {
		return nil, nil, err
	}
	usersStorage := pg.NewUserStorage(db)
	roleStorage := pg.NewRoleStorage(db)
	client, cleanup2, err := _redis(configConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	sessionsStorage := redis.NewSessionsStorage(client, configConfig)
	authService := authservice.New(usersStorage, usersStorage, roleStorage, sessionsStorage, configConfig)
	service := userservice.New(usersStorage, configConfig)
	v := _servers(configConfig, authService, service)
	app := newApp(configConfig, v)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

func _pg(cfg *config.Config) (*sqlx.DB, func(), error) {
	host := cfg.Pg.Host
	port := cfg.Pg.Port
	user := cfg.Pg.User
	pass := cfg.Pg.Pass
	name := cfg.Pg.Name

	cs := fmt.Sprintf(`postgres://%s:%s@%s:%d/%s?sslmode=disable`, user, pass, host, port, name)

	db, err := sqlx.Connect("pgx", cs)
	if err != nil {
		return nil, nil, err
	}
	slog.Debug("connecting to database", slog.String("conn", cs))
	t := time.Now()
	if err := db.Ping(); err != nil {
		slog.Error("failed to connect to database", slog.String("err", err.Error()), slog.String("conn", cs))
		return nil, func() { db.Close() }, err
	}
	slog.Info("connected to database", slog.String("ping", fmt.Sprintf("%2.fs", time.Since(t).Seconds())))

	return db, func() { db.Close() }, nil
}

func _redis(cfg *config.Config) (*redis2.Client, func(), error) {
	host := cfg.Redis.Host
	port := cfg.Redis.Port
	pass := cfg.Redis.Pass

	cs := fmt.Sprintf(`redis://%s:%s@%s:%d`, host, pass, host, port)

	log := slog.With()

	client := redis2.NewClient(&redis2.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: pass,
		DB:       0,
	})

	log.Debug("connecting to redis", slog.String("conn", cs))
	t := time.Now()
	if err := client.Ping(context.Background()).Err(); err != nil {
		slog.Error("failed to connect to redis", slog.String("err", err.Error()), slog.String("conn", cs))
		return nil, func() { client.Close() }, err
	}
	log.Info("connected to redis", slog.String("ping", fmt.Sprintf("%2.fs", time.Since(t).Seconds())))

	return client, func() { client.Close() }, nil
}

func _servers(cfg *config.Config, as *authservice.AuthService, us *userservice.Service) []Server {
	servers := make([]Server, 0, 2)

	if cfg.Http.Enabled {
		servers = append(servers, http.New(cfg, as))
	}

	if cfg.Grpc.Enabled {
		servers = append(servers, grpc.New(cfg, as, us))
	}

	return servers
}
