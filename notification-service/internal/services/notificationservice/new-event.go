package notificationservice

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/pkg/sl"
)

func (s *Service) ProcessNewEvent(ctx context.Context, event *domain.EventInfo) error {

	fn := "notification-service.ProcessNewEvent"
	log := s.l.With(sl.Method(fn))

	log.Debug("getting subscribers for new event", slog.String("eventId", event.Id))
	// TODO UPDATE TO SportSubscribers
	subscribers, err := s.sp.EventSubscribers(ctx, event.Id)
	if err != nil {
		log.Error("failed getting subscribers for new event", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	if len(subscribers) == 0 {
		log.Debug("no subscribers for new event", slog.String("eventId", event.Id))
		return nil
	}

	for _, subscriberId := range subscribers {
		log.Debug("sending notification to subscriber", slog.Any("subscriber", subscriberId))

		user, err := s.up.Find(ctx, subscriberId)
		if err != nil {
			log.Error("failed getting user", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}

		integrations, err := s.ip.Find(ctx, user.Id)
		if err != nil {
			log.Error("failed getting integrations for subscriber", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}

		if integrations.TelegramUsername != nil {
			if err := s.notificator.SendTelegram(ctx, *integrations.TelegramUsername, event, domain.NewEvent); err != nil {
				log.Error("failed sending notification to subscriber (telegram)", sl.Err(err))
				return fmt.Errorf("%s: %w", fn, err)
			}
		}

		if integrations.WannaMail {
			if err := s.notificator.SendMail(ctx, user.Email, event, domain.NewEvent); err != nil {
				log.Error("failed sending notification to subscriber (mail)", sl.Err(err))
				return fmt.Errorf("%s: %w", fn, err)
			}
		}
	}

	return nil
}
