package subscriptionservice

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mzhn/subscriptions-service/internal/domain"
	"mzhn/subscriptions-service/pb/espb"
	"mzhn/subscriptions-service/pb/sspb"
	"mzhn/subscriptions-service/pkg/sl"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	l       *slog.Logger
	storage domain.SubscriptionsStorage
	es      espb.EventServiceClient
}

func New(storage domain.SubscriptionsStorage, es espb.EventServiceClient) *Service {
	return &Service{
		l:       slog.With(sl.Module(("SubscriptionsService"))),
		storage: storage,
		es:      es,
	}
}
func (s *Service) SubscribeToSport(dto *domain.SportSubscription) (*domain.SportSubscription, error) {
	log := s.l.With(sl.Method("SubscriptionsService.SubscribeToSport"))

	log.Debug("creating sport subscription", slog.Any("userId", dto.UserId), slog.Any("sportId", dto.SportId))
	return s.storage.CreateSport(dto)
}

func (s *Service) SubscribeToEvent(dto *domain.EventSubscription) (*domain.EventSubscription, error) {
	log := s.l.With(sl.Method("SubscriptionsService.SubscribeToEvent"))

	log.Debug("creating event subscription", slog.Any("userId", dto.UserId), slog.Any("eventId", dto.EventId))
	return s.storage.CreateEvent(dto)
}

func (s *Service) UnsubscribeFromSport(dto *domain.SportSubscription) error {
	log := s.l.With(sl.Method("SubscriptionsService.UnsubscribeFromSport"))

	log.Debug("unsubscribe from sport", slog.Any("userId", dto.UserId), slog.Any("sportId", dto.SportId))
	return s.storage.DeleteSport(dto)
}

func (s *Service) UnsubscribeFromEvent(dto *domain.EventSubscription) error {
	log := s.l.With(sl.Method("SubscriptionsService.UnsubscribeFromEvent"))

	log.Debug("unsubscribe from event", slog.Any("userId", dto.UserId), slog.Any("eventId", dto.EventId))
	return s.storage.DeleteEvent(dto)

}

func (s *Service) GetUserEvents(userId string) ([]*espb.EventInfo, error) {
	log := s.l.With(sl.Method("SubscriptionsService.GetUserEvents"))
	log.Debug("get user events", slog.Any("userId", userId))

	eventIds, err := s.storage.GetUserEventsId(userId)
	if err != nil {
		return nil, err
	}

	if len(eventIds) == 0 {
		return []*espb.EventInfo{}, nil
	}

	stream, err := s.es.Events(context.Background())
	if err != nil {
		return nil, err
	}
	defer stream.CloseSend()

	// Канал для сбора результатов
	eventsChan := make(chan *espb.EventInfo)
	// Канал для ошибок
	errChan := make(chan error, 1)
	// Канал для сигнализации о завершении отправки
	sendDone := make(chan struct{})

	// Горутина для отправки запросов
	go func() {
		for _, eventId := range eventIds {
			if err := stream.Send(&espb.EventRequest{
				Id: eventId,
			}); err != nil {
				s.l.Error("cannot send while get user events", slog.Any("error", err))
				errChan <- err
				return
			}
		}
		if err := stream.CloseSend(); err != nil {
			s.l.Error("cannot close send while get user events", slog.Any("error", err))
			errChan <- err
			return
		}
		close(sendDone)
	}()

	// Горутина для получения ответов
	go func() {
		for {
			eventInfo, err := stream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					// Нормальное завершение стрима
					close(eventsChan)
					return
				}
				if grpcErr, ok := status.FromError(err); ok {
					if grpcErr.Code() == codes.NotFound {
						// Пропускаем не найденные события
						continue
					}
				}
				errChan <- err
				return
			}
			fmt.Println(eventInfo.Info.Id)
			eventsChan <- eventInfo.Info
		}
	}()

	// Собираем результаты
	var events []*espb.EventInfo

	// Ждем завершения отправки и получения всех событий
	for {
		select {
		case err := <-errChan:
			return nil, err
		case event, ok := <-eventsChan:
			if !ok {
				// Канал закрыт - все события получены
				return events, nil
			}
			events = append(events, event)
		}
	}
}

func (s *Service) GetUsersSubscribedToEvent(eventId string) ([]string, error) {
	log := s.l.With(sl.Method("SubscriptionsService.GetUsersSubscribedToEvent"))

	log.Debug("get users subscribed to event", slog.Any("eventId", eventId))
	return s.storage.GetUsersSubscribedToEvent(eventId)
}

func (s *Service) GetUsersSubscribedToSport(sportId string) ([]string, error) {
	log := s.l.With(sl.Method("SubscriptionsService.GetUsersSubscribedToSport"))

	log.Debug("get users subscribed to sport", slog.Any("sportId", sportId))
	return s.storage.GetUsersSubscribedToSport(sportId)
}

func (s *Service) GetUsersFromEventByDaysLeft(eventId string, daysLeft sspb.DaysLeft) ([]string, error) {
	log := s.l.With(sl.Method("SubscriptionsService.GetUsersFromEventByDaysLeft"))

	log.Debug("get users from event by days left", slog.Any("eventId", eventId))
	return s.storage.GetUsersFromEventByDaysLeft(eventId, daysLeft)
}

func (s *Service) NotifyUser(userId string, daysLeft sspb.DaysLeft) error {
	log := s.l.With(sl.Method("SubscriptionsService.NotifyUser"))

	log.Debug("notify user", slog.Any("userId", userId), slog.Any("daysLeft", daysLeft.String()))
	return s.storage.NotifyUser(userId, daysLeft)
}
