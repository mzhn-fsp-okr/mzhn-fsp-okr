package subscriptionservice

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"mzhn/subscriptions-service/internal/domain"
	"mzhn/subscriptions-service/pb/espb"
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

	eventsInfo := make([]*espb.EventInfo, 0, len(eventIds))

	stream, err := s.es.Events(context.Background())
	if err != nil {
		return nil, err
	}

	done := make(chan error)

	go func() {
		for {
			eventInfo, err := stream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					done <- nil
				}

				if grpcErr, ok := status.FromError(err); ok {
					if grpcErr.Code() == codes.NotFound {
						// Todo delete
						continue
					}
				}
				done <- err
			}
			eventsInfo = append(eventsInfo, eventInfo.Info)
		}
	}()

	for _, eventId := range eventIds {
		if err := stream.Send(&espb.EventRequest{
			Id: eventId,
		}); err != nil {
			return nil, err
		}
	}
	if err := stream.CloseSend(); err != nil {
		return nil, err
	}

	if err := <-done; err != nil {
		return nil, err
	}

	return eventsInfo, nil
}
