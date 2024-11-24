package amqp

import (
	"context"
	"encoding/json"
	"log/slog"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/pkg/sl"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func (a *RabbitMqConsumer) consumeNewEvents(ctx context.Context) error {

	fn := "rmq-consumer.upcoming-events"
	log := a.l.With(sl.Module(fn))

	messages, err := a.channel.ConsumeWithContext(
		ctx,
		a.cfg.Amqp.NewEventsQueue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	start := time.Now()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-messages:
				if err := a.consumeNewEvent(ctx, message); err != nil {
					log.Error("failed to consume upcoming event", sl.Err(err), slog.String("body", string(message.Body)))
					continue
				}
			}
		}
	}()

	<-ctx.Done()
	log.Info("stop consuming new events", slog.String("uptime", time.Since(start).String()))
	return nil
}

func (a *RabbitMqConsumer) consumeNewEvent(ctx context.Context, message amqp091.Delivery) (err error) {
	defer func() {
		if err != nil {
			message.Nack(false, true)
		} else {
			err = message.Ack(false)
		}
	}()

	body := message.Body
	msg := domain.EventInfo{}
	if err := json.Unmarshal(body, &msg); err != nil {
		a.l.Error("failed to unmarshal event", sl.Err(err))
		return err
	}

	a.l.Info("received message", slog.Any("message", msg))

	if err := a.ns.ProcessNewEvent(ctx, &msg); err != nil {
		a.l.Error("failed to process upcoming event", sl.Err(err), slog.String("body", string(body)))
		return err
	}

	return nil
}