package amqp

import (
	"context"
	"encoding/json"
	"log/slog"
	"mzhn/subscriptions-service/internal/config"
	"mzhn/subscriptions-service/pkg/sl"

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

func (r *RabbitMQ) NotifyAboutSubscription(ctx context.Context, userId string, entityId string, isEvent bool) error {
	r.l.Debug("publish new event to queue")
	payload := map[string]interface{}{
		"userId":   userId,
		"entityId": entityId,
		"isEvent":  isEvent,
	}

	payloadJson, err := json.Marshal(&payload)
	if err != nil {
		r.l.Error("Error while marshaling", sl.Err(err))
		return err
	}

	r.l.Info("publishing event to upcoming queue")
	if err := r.channel.PublishWithContext(ctx, r.cfg.Amqp.NotificationsExchange, r.cfg.Amqp.NewSubscriptionQueue, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        payloadJson,
	}); err != nil {
		r.l.Error("Error while send event", sl.Err(err))
		return err
	}
	return nil
}
