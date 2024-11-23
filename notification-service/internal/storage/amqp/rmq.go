package amqp

import (
	"context"
	"encoding/json"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/pkg/sl"

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

func (r *RabbitMQ) SendTelegram(ctx context.Context, data map[string]any) error {

	message, err := json.Marshal(data)
	if err != nil {
		return err
	}

	log.Info("publishing event to notification queue")
	if err := r.channel.PublishWithContext(ctx, r.cfg.Amqp.SubscriptionExchange, r.cfg.Amqp.TelegramQueue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        message,
	}); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) SendMail(ctx context.Context, data map[string]any) error {
	message, err := json.Marshal(data)
	if err != nil {
		return err
	}

	log.Info("publishing event to notification queue")
	if err := r.channel.PublishWithContext(ctx, r.cfg.Amqp.SubscriptionExchange, r.cfg.Amqp.MailQueue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        message,
	}); err != nil {
		return err
	}

	return nil
}
