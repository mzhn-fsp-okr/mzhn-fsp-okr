package amqp

import (
	"context"
	"encoding/json"
	"log/slog"
	"mzhn/notification-cron/internal/config"
	"mzhn/notification-cron/pb/sspb"
	"mzhn/notification-cron/pkg/sl"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	cfg     *config.Config
	l       *slog.Logger
	ss      sspb.SubscriptionServiceClient
	channel *amqp091.Channel
}

func New(cfg *config.Config, channel *amqp091.Channel, ss sspb.SubscriptionServiceClient) *RabbitMQ {

	return &RabbitMQ{
		cfg:     cfg,
		channel: channel,
		ss:      ss,
		l:       slog.With(sl.Module("rabbitmq")),
	}
}

func (r *RabbitMQ) NotifyAboutUpcomingEvent(ctx context.Context, userId string, eventId string, daysLeft uint32, daysLeftEnum sspb.DaysLeft) error {
	r.l.Debug("publish new event to queue")
	payload := map[string]interface{}{
		"userId":   userId,
		"eventId":  eventId,
		"daysLeft": daysLeft,
	}

	payloadJson, err := json.Marshal(&payload)
	if err != nil {
		r.l.Error("Error while marshaling", sl.Err(err))
		return err
	}

	r.l.Info("publishing event to upcoming queue")
	if err := r.channel.PublishWithContext(ctx, r.cfg.Amqp.NotificationsExchange, r.cfg.Amqp.UpcomingQueue, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        payloadJson,
	}); err != nil {
		r.l.Error("Error while send event", sl.Err(err))
		return err
	}

	r.ss.NotifyUser(ctx, &sspb.NotifyUserRequest{UserId: userId, DaysLeft: daysLeftEnum, EventId: eventId})
	return nil
}
