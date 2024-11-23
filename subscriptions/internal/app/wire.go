//go:build wireinject
// +build wireinject

package app

import (
	"fmt"
	"log/slog"

	"mzhn/subscriptions-service/internal/config"
	"mzhn/subscriptions-service/internal/domain"
	"mzhn/subscriptions-service/internal/services/authservice"
	"mzhn/subscriptions-service/internal/services/subscriptionservice"
	"mzhn/subscriptions-service/internal/storage/amqp"
	subscriptions_storage "mzhn/subscriptions-service/internal/storage/pg/subscriptions"
	"mzhn/subscriptions-service/internal/transport/http"
	"mzhn/subscriptions-service/pb/authpb"
	"mzhn/subscriptions-service/pb/espb"
	"mzhn/subscriptions-service/pkg/sl"

	ssgrpc "mzhn/subscriptions-service/internal/transport/grpc"

	"github.com/google/wire"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*App, func(), error) {
	panic(wire.Build(
		newApp,
		_servers,
		wire.Bind(new(domain.SubscribeNotificationPublisher), new(*amqp.RabbitMQ)),
		wire.Bind(new(domain.SubscriptionsStorage), new(*subscriptions_storage.Storage)),
		authservice.New,
		subscriptionservice.New,
		subscriptions_storage.New,
		amqp.New,
		_amqp,
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

func _amqp(cfg *config.Config) (*amqp091.Channel, func(), error) {

	cs := cfg.Amqp.ConnectionString()

	conn, err := amqp091.Dial(cs)
	if err != nil {
		return nil, nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, func() {
			conn.Close()
		}, err
	}

	slog.Info("declaring notifications exchange", slog.String("exchange", cfg.Amqp.NotificationsExchange))
	if err := channel.ExchangeDeclare(cfg.Amqp.NotificationsExchange, "direct", true, false, false, false, nil); err != nil {
		slog.Error("failed to declare notifications queue", sl.Err(err))
		return nil, func() {
			channel.Close()
			conn.Close()
		}, err
	}

	slog.Info("declaring upcoming events queue", slog.String("queue", cfg.Amqp.NewSubscriptionQueue))
	q, err := channel.QueueDeclare(cfg.Amqp.NewSubscriptionQueue, true, false, false, false, nil)
	if err != nil {
		slog.Error("failed to declare upcoming events queue", sl.Err(err))
		return nil, func() {
			channel.Close()
			conn.Close()
		}, err
	}

	slog.Info(
		"binding upcoming events queue",
		slog.String("queue", cfg.Amqp.NewSubscriptionQueue),
		slog.String("exchange", cfg.Amqp.NotificationsExchange),
	)
	if err := channel.QueueBind(q.Name, cfg.Amqp.NewSubscriptionQueue, cfg.Amqp.NotificationsExchange, false, nil); err != nil {
		slog.Error("failed to bind new events queue", sl.Err(err))
		return nil, func() {
			channel.Close()
			conn.Close()
		}, err
	}

	return channel, func() {
		channel.Close()
		conn.Close()
	}, nil
}
