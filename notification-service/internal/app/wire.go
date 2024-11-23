//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/services/authservice"
	"mzhn/notification-service/internal/services/integrationservice"
	"mzhn/notification-service/internal/services/notificationservice"
	amqpclient "mzhn/notification-service/internal/storage/amqp"
	"mzhn/notification-service/internal/storage/grpc/authapi"
	"mzhn/notification-service/internal/storage/grpc/eventsapi"
	"mzhn/notification-service/internal/storage/grpc/subscribersapi"
	"mzhn/notification-service/internal/storage/pg/integrationstorage"
	"mzhn/notification-service/internal/storage/pg/verificationstorage"
	amqptransport "mzhn/notification-service/internal/transport/amqp"
	"mzhn/notification-service/internal/transport/http"
	"mzhn/notification-service/pb/authpb"
	"mzhn/notification-service/pb/espb"
	"mzhn/notification-service/pb/sspb"
	"mzhn/notification-service/pkg/sl"

	grpcserver "mzhn/notification-service/internal/transport/grpc"

	"github.com/google/wire"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New() (*App, func(), error) {
	panic(wire.Build(
		newApp,
		_servers,

		http.New,
		amqptransport.New,
		grpcserver.New,

		authservice.New,
		wire.Bind(new(authservice.ProfileProvider), new(*authapi.Api)),

		integrationservice.New,
		wire.Bind(new(integrationservice.IntegrationsSaver), new(*integrationstorage.Storage)),
		wire.Bind(new(integrationservice.IntegrationsProvider), new(*integrationstorage.Storage)),
		wire.Bind(new(integrationservice.VerificationSaver), new(*verificationstorage.Storage)),
		wire.Bind(new(integrationservice.VerificationProvider), new(*verificationstorage.Storage)),

		notificationservice.New,
		wire.Bind(new(notificationservice.UserProvider), new(*authapi.Api)),
		wire.Bind(new(notificationservice.EventProvider), new(*eventsapi.Api)),
		wire.Bind(new(notificationservice.SportProvider), new(*eventsapi.Api)),
		wire.Bind(new(notificationservice.Notificator), new(*amqpclient.RabbitMQ)),
		wire.Bind(new(notificationservice.SubscribersProvider), new(*subscribersapi.Api)),
		wire.Bind(new(notificationservice.IntegrationProvider), new(*integrationstorage.Storage)),

		authapi.New,
		subscribersapi.New,
		eventsapi.New,
		amqpclient.New,
		integrationstorage.New,
		verificationstorage.New,

		_authpb,
		_sspb,
		_espb,
		_amqp,
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

func _sspb(cfg *config.Config) (sspb.SubscriptionServiceClient, error) {
	addr := cfg.SubscriptionService.ConnectionString()
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return sspb.NewSubscriptionServiceClient(conn), nil
}

func _espb(cfg *config.Config) (espb.EventServiceClient, error) {
	addr := cfg.EventsService.ConnectionString()
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return espb.NewEventServiceClient(conn), nil
}

func _servers(
	http *http.Server,
	amqp *amqptransport.RabbitMqConsumer,
	grpc *grpcserver.Server,
) []Server {
	servers := make([]Server, 0, 2)
	servers = append(servers, amqp)
	servers = append(servers, http)
	servers = append(servers, grpc)
	return servers
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

	if err := amqp_setup_exchange(cfg, channel, cfg.Amqp.NotificationsExchange, cfg.Amqp.NewEventsQueue, cfg.Amqp.UpcomingEventsQueue); err != nil {
		slog.Error("failed to setup notifications exchange", sl.Err(err))
		return nil, func() {
			channel.Close()
			conn.Close()
		}, err
	}

	if err := amqp_setup_exchange(cfg, channel, cfg.Amqp.SubscriptionExchange, cfg.Amqp.TelegramQueue, cfg.Amqp.MailQueue); err != nil {
		slog.Error("failed to setup subscribtions exchange", sl.Err(err))
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
func amqp_setup_exchange(cfg *config.Config, channel *amqp091.Channel, exchange string, queues ...string) error {

	log := slog.With(slog.String("exchange", exchange))
	log.Info("declaring exchange")
	if err := channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
		slog.Error("failed to declare notifications queue", sl.Err(err))
		return err
	}

	for _, queueName := range queues {
		log.Info("declaring queue", slog.String("queue", queueName))
		queue, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
		if err != nil {
			log.Error("failed to declare queue", sl.Err(err), slog.String("queue", queueName))
			return err
		}

		log.Info("binding queue", slog.String("queue", queueName))
		if err := channel.QueueBind(queue.Name, queueName, exchange, false, nil); err != nil {
			log.Error("failed to bind queue", sl.Err(err), slog.String("queue", queueName))
			return err
		}
	}

	return nil
}
