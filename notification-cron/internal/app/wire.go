//go:build wireinject
// +build wireinject

package app

import (
	"fmt"

	"mzhn/notification-cron/internal/config"
	"mzhn/notification-cron/internal/domain"
	"mzhn/notification-cron/internal/services/authservice"
	"mzhn/notification-cron/internal/services/subscriptionservice"
	subscriptions_storage "mzhn/notification-cron/internal/storage/pg/subscriptions"
	"mzhn/notification-cron/internal/transport/http"
	"mzhn/notification-cron/pb/authpb"
	"mzhn/notification-cron/pb/espb"

	ssgrpc "mzhn/notification-cron/internal/transport/grpc"

	"github.com/google/wire"
	_ "github.com/jackc/pgx/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*App, func(), error) {
	panic(wire.Build(
		newApp,
		_servers,
		wire.Bind(new(domain.SubscriptionsStorage), new(*subscriptions_storage.Storage)),
		authservice.New,
		subscriptionservice.New,
		subscriptions_storage.New,
		_authpb,
		_eventspb,
		_db,
		config.New,
	))
}

func _servers(cfg *config.Config, ss *subscriptionservice.Service, as *authservice.Service) []Server {
	servers := make([]Server, 0, 2)

	if cfg.Http.Enabled {
		servers = append(servers, http.New(cfg, as, ss))
	}

	if cfg.Grpc.Enabled {
		servers = append(servers, ssgrpc.New(cfg, ss))
	}
	return servers
}

func _authpb(cfg *config.Config) (authpb.AuthClient, error) {
	addr := cfg.AuthService.ConnectionString()
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return authpb.NewAuthClient(conn), nil
}

func _eventspb(cfg *config.Config) (espb.EventServiceClient, error) {
	addr := cfg.EventService.ConnectionString()
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return espb.NewEventServiceClient(conn), nil
}

func _db(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Pg.Host, cfg.Pg.User, cfg.Pg.Pass, cfg.Pg.Name, cfg.Pg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&domain.SportSubscription{}, &domain.EventSubscription{}); err != nil {
		return nil, err
	}

	return db, nil
}
