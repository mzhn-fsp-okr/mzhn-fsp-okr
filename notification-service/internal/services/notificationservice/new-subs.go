package notificationservice

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/pkg/sl"
)

func (s *Service) ProcessNewSub(ctx context.Context, msg *domain.NewSubscriptionsMessage) error {

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

	if integrations == nil {
		log.Error("integrations not found", slog.Any("for", user))
		return nil
	}

	if msg.IsEvent {
		return s.pushEvent(ctx, user, msg.EntityId, integrations, domain.EventNewSub)
	} else {
		return s.pushSport(ctx, user, msg.EntityId, integrations, domain.EventNewSub)
	}
}

func (s *Service) pushEvent(ctx context.Context, user *domain.User, eventId string, integrations *domain.Integrations, t domain.EventType) error {
	fn := "notification-service.pushEvent"
	log := s.l.With(sl.Method(fn))

	log.Info(
		"sending notification to subscriber",
		slog.Any("subscriber", user),
		slog.Any("integrations", integrations),
		slog.Any("event", eventId),
		slog.Any("type", t.String()),
	)

	event, err := s.ep.Event(ctx, eventId)
	if err != nil {
		log.Error("failed getting event", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}
	if event == nil {
		log.Error("event not found", slog.Any("for", eventId))
		return nil
	}

	log.Debug("found event", slog.Any("event", event))

	j := map[string]any{
		"event": event,
		"type":  t.String(),
	}

	if integrations.TelegramUsername != nil {
		log.Debug("sending notification to subscriber (telegram)")
		j["receiver"] = *integrations.TelegramUsername
		if err := s.notificator.SendTelegram(ctx, j); err != nil {
			log.Error("failed sending notification to subscriber (telegram)", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}
	}

	if integrations.WannaMail {
		j["receiver"] = user.Email
		if err := s.notificator.SendMail(ctx, j); err != nil {
			log.Error("failed sending notification to subscriber (mail)", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}
	}

	return nil
}

func (s *Service) pushSport(ctx context.Context, user *domain.User, sportId string, integrations *domain.Integrations, t domain.EventType) error {
	fn := "notification-service.pushSport"
	log := s.l.With(sl.Method(fn))

	sport, err := s.sportProvider.Sport(ctx, sportId)
	if err != nil {
		log.Error("failed getting event", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}
	if sport == nil {
		log.Error("sport not found", slog.Any("for", sportId))
		return nil
	}

	j := map[string]any{
		"sport": sport,
		"type":  t.String(),
	}

	if integrations.TelegramUsername != nil {
		j["receiver"] = *integrations.TelegramUsername
		if err := s.notificator.SendTelegram(ctx, j); err != nil {
			log.Error("failed sending notification to subscriber (telegram)", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}
	}

	if integrations.WannaMail {
		j["receiver"] = user.Email
		if err := s.notificator.SendMail(ctx, j); err != nil {
			log.Error("failed sending notification to subscriber (mail)", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}
	}

	return nil
}
