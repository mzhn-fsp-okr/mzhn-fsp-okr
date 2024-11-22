package amqp

import (
	"context"
	"encoding/json"
	"log/slog"
	"mzhn/event-service/internal/config"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/pkg/sl"

	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	cfg     *config.Config
	l       *slog.Logger
	channel *amqp.Channel
}

func New(cfg *config.Config, channel *amqp.Channel) *RabbitMQ {

	return &RabbitMQ{
		cfg:     cfg,
		channel: channel,
		l:       slog.With(sl.Module("rabbitmq")),
	}
}

func (r *RabbitMQ) Notification(ctx context.Context, event *domain.EventInfo) error {

	eventJson, err := json.Marshal(event)
	if err != nil {
		return err
	}

	log.Info("publishing event to notification queue")
	if err := r.channel.PublishWithContext(ctx, r.cfg.Amqp.NotificationsExchange, r.cfg.Amqp.NewEventsQueue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        eventJson,
	}); err != nil {
		return err
	}

	return nil
}
