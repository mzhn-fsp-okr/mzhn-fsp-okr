package amqp

import (
	"context"
	"encoding/json"
	"log/slog"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/pkg/sl"
	"time"
)

func (a *RabbitMqConsumer) consumeUpcomingEvents(ctx context.Context) error {

	fn := "rmq-consumer.upcoming-events"
	log := a.l.With(sl.Module(fn))

	messages, err := a.channel.ConsumeWithContext(
		ctx,
		a.cfg.Amqp.UpcomingEventsQueue,
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
				body := message.Body
				msg := domain.UpcomingEventMessage{}
				if err := json.Unmarshal(body, &msg); err != nil {
					log.Error("failed to unmarshal event", sl.Err(err))
					continue
				}

				log.Info("received message", slog.Any("message", msg))

				if err := a.ns.ProcessUpcomingEvent(ctx, &msg); err != nil {
					log.Error("failed to process upcoming event", sl.Err(err), slog.String("body", string(body)))
					continue
				}

				if err := message.Ack(false); err != nil {
					log.Error("failed to ack message", sl.Err(err), slog.String("body", string(body)))
					continue
				}
			}
		}
	}()

	<-ctx.Done()
	log.Info("stop consuming upcoming events", slog.String("uptime", time.Since(start).String()))
	return nil
}
