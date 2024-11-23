package amqp

import (
	"context"
	"encoding/json"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/domain"
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

func (r *RabbitMQ) SendTelegram(ctx context.Context, username string, event *domain.EventInfo, eType domain.EventType) error {

	message, err := eventJson(eType, username, event)
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

func (r *RabbitMQ) SendMail(ctx context.Context, email string, event *domain.EventInfo, eType domain.EventType) error {

	message, err := eventJson(eType, email, event)
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

func eventJson(t domain.EventType, r string, e *domain.EventInfo) ([]byte, error) {
	eventJson, err := json.Marshal(map[string]any{
		"eventType": t.String(),
		"receiver":  r,
		"event":     e,
	})
	if err != nil {
		return nil, err
	}

	return eventJson, nil
}
