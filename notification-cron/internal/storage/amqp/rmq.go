package amqp

import (
	"context"
	"encoding/json"
	"log/slog"
	"mzhn/notification-cron/internal/config"
	"mzhn/notification-cron/pkg/sl"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	cfg     *config.Config
	l       *slog.Logger
	channel *amqp091.Channel
}

func New(cfg *config.Config, channel *amqp091.Channel) *RabbitMQ {

	return &RabbitMQ{
		cfg:     cfg,
		channel: channel,
		l:       slog.With(sl.Module("rabbitmq")),
	}
}

func (r *RabbitMQ) NotifyAboutUpcomingEvent(ctx context.Context, userId string, eventId string, daysLeft uint32) error {
	payload := map[string]interface{}{
		"userId":   userId,
		"eventId":  eventId,
		"daysLeft": daysLeft,
	}

	payloadJson, err := json.Marshal(&payload)
	if err != nil {
		return err
	}

	r.l.Info("publishing event to upcoming queue")
	if err := r.channel.PublishWithContext(ctx, r.cfg.Amqp.NotificationsExchange, r.cfg.Amqp.UpcomingQueue, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        payloadJson,
	}); err != nil {
		return err
	}

	return nil
}
