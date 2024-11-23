package amqp

import (
	"context"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/services/notificationservice"
	"mzhn/notification-service/pkg/sl"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMqConsumer struct {
	l       *slog.Logger
	cfg     *config.Config
	channel *amqp091.Channel

	ns *notificationservice.Service

	upcomingEventsQueue string
	newEventsQueue      string
}

func New(cfg *config.Config, channel *amqp091.Channel, ns *notificationservice.Service) *RabbitMqConsumer {
	return &RabbitMqConsumer{
		l:       slog.With(sl.Module("rabbitmq-consumer")),
		cfg:     cfg,
		channel: channel,
		ns:      ns,

		upcomingEventsQueue: cfg.Amqp.UpcomingEventsQueue,
		newEventsQueue:      cfg.Amqp.NewEventsQueue,
	}

}

func (r *RabbitMqConsumer) Run(ctx context.Context) error {

	go func(ctx context.Context) {
		if err := r.consumeUpcomingEvents(ctx); err != nil {
			r.l.Error("consumeUpcomingEvents", sl.Err(err))
			return
		}
	}(ctx)

	// go func(ctx context.Context) {
	// 	if err := r.consumeNewEvents(ctx); err != nil {
	// 		r.l.Error("consumeUpcomingEvents", sl.Err(err))
	// 		return
	// 	}
	// }(ctx)

	<-ctx.Done()
	return nil
}
