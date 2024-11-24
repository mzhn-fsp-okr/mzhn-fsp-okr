//go:build wireinject
// +build wireinject

package app

import (
	"log/slog"
	"mzhn/notification-cron/internal/config"
	"mzhn/notification-cron/internal/cron"
	"mzhn/notification-cron/internal/domain"
	"mzhn/notification-cron/internal/services/cronservice"
	"mzhn/notification-cron/internal/storage/amqp"
	"mzhn/notification-cron/pb/espb"
	"mzhn/notification-cron/pb/sspb"
	"mzhn/notification-cron/pkg/sl"

	"github.com/google/wire"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New() (*App, func(), error) {
	panic(wire.Build(
		newApp,
		_servers,
		wire.Bind(new(domain.UpcomingNotificationPublisher), new(*amqp.RabbitMQ)),
		wire.Bind(new(domain.CronService), new(*cronservice.Service)),

		cronservice.New,

		_eventspb,
		_subscriptionspb,

		amqp.New,
		_amqp,

		config.New,
	))
}

func _servers(cfg *config.Config, cs domain.CronService) []Server {
	servers := make([]Server, 0, 1)
	servers = append(servers, cron.New(cfg, cs))

	return servers
}

func _eventspb(cfg *config.Config) (espb.EventServiceClient, error) {
	addr := cfg.EventService.ConnectionString()
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return espb.NewEventServiceClient(conn), nil
}

func _subscriptionspb(cfg *config.Config) (sspb.SubscriptionServiceClient, error) {
	addr := cfg.SubscriptionsService.ConnectionString()
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return sspb.NewSubscriptionServiceClient(conn), nil
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

	slog.Info("declaring upcoming events queue", slog.String("queue", cfg.Amqp.UpcomingQueue))
	q, err := channel.QueueDeclare(cfg.Amqp.UpcomingQueue, true, false, false, false, nil)
	if err != nil {
		slog.Error("failed to declare upcoming events queue", sl.Err(err))
		return nil, func() {
			channel.Close()
			conn.Close()
		}, err
	}

	slog.Info(
		"binding upcoming events queue",
		slog.String("queue", cfg.Amqp.UpcomingQueue),
		slog.String("exchange", cfg.Amqp.NotificationsExchange),
	)
	if err := channel.QueueBind(q.Name, cfg.Amqp.UpcomingQueue, cfg.Amqp.NotificationsExchange, false, nil); err != nil {
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
