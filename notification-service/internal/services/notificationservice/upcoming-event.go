package notificationservice

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/pkg/sl"
)

func (s *Service) ProcessUpcomingEvent(ctx context.Context, msg *domain.UpcomingEventMessage) error {

	fn := "notification-service.ProcessUpcomingEvent"
	log := s.l.With(sl.Method(fn))

	log.Debug("sending notification to subscriber", slog.Any("subscriber", msg.UserId))

	user, err := s.up.Find(ctx, msg.UserId)
	if err != nil {
		log.Error("failed getting user", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	integrations, err := s.ip.Find(ctx, user.Id)
	if err != nil {
		log.Error("failed getting integrations for subscriber", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	event, err := s.ep.Find(ctx, msg.EventId)
	if err != nil {
		log.Error("failed getting event", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	if integrations.TelegramUsername != nil {
		if err := s.notificator.SendTelegram(ctx, *integrations.TelegramUsername, event, domain.EventTypeUpcoming); err != nil {
			log.Error("failed sending notification to subscriber (telegram)", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}
	}

	if integrations.WannaMail {
		if err := s.notificator.SendMail(ctx, user.Email, event, domain.EventTypeUpcoming); err != nil {
			log.Error("failed sending notification to subscriber (mail)", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}
	}

	return nil
}
