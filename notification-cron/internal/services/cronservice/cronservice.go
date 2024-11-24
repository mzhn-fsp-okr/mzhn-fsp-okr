package cronservice

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"mzhn/notification-cron/internal/config"
	"mzhn/notification-cron/internal/domain"
	"mzhn/notification-cron/pb/espb"
	"mzhn/notification-cron/pb/sspb"
	"mzhn/notification-cron/pkg/sl"
	"sync"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	l   *slog.Logger
	cfg *config.Config

	sspb sspb.SubscriptionServiceClient
	espb espb.EventServiceClient
	pub  domain.UpcomingNotificationPublisher
}

func New(cfg *config.Config, sspb sspb.SubscriptionServiceClient, espb espb.EventServiceClient, pub domain.UpcomingNotificationPublisher) *Service {
	return &Service{
		l:   slog.With(sl.Module("CronService")),
		cfg: cfg,

		sspb: sspb,
		espb: espb,
		pub:  pub,
	}
}

func (s *Service) NotifyUsers(ctx context.Context) error {
	s.l.Debug("NotifyUsers")

	const workerCount = 5
	eventsChan := make(chan *espb.UpcomingEventResponse)
	errChan := make(chan error, 1)

	stream, err := s.espb.GetUpcomingEvents(ctx, &emptypb.Empty{})
	if err != nil {
		s.l.Error("cannot get upcoming events", sl.Err(err))
		return err
	}

	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for event := range eventsChan {
				if err := s.proccessEvent(ctx, event); err != nil {
					s.l.Error("failed to proccess event", sl.Err(err))
				}
			}
		}()
	}

	go func() {
		defer close(eventsChan)
		for {
			event, err := stream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				errChan <- err
				return
			}
			eventsChan <- event
		}
	}()

	select {
	case err := <-errChan:
		s.l.Error("Error while read notifyUser channel", sl.Err(err))
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Service) proccessEvent(ctx context.Context, event *espb.UpcomingEventResponse) error {
	daysLeft := sspb.DaysLeft_ThreeDays
	if event.DaysLeft <= 30 && event.DaysLeft > 7 {
		daysLeft = sspb.DaysLeft_Month
	} else if event.DaysLeft <= 7 && event.DaysLeft > 3 {
		daysLeft = sspb.DaysLeft_Week
	} else {
		daysLeft = sspb.DaysLeft_ThreeDays
	}

	userIds, err := s.getUsersToNotify(ctx, &sspb.UsersEventByDaysRequest{
		EventId:  event.Event.Id,
		DaysLeft: daysLeft,
	})
	if err != nil {
		s.l.Error("cannot get users to notify", sl.Err(err))
		return err
	}
	if len(userIds) != 0 {
		s.l.Debug("users to notify length", slog.Any("count", len(userIds)))
	}

	for _, userId := range userIds {
		s.l.Debug("trying to notify user")
		s.pub.NotifyAboutUpcomingEvent(ctx, userId, event.Event.Id, event.DaysLeft, daysLeft)
	}

	return nil
}

func (s *Service) getUsersToNotify(ctx context.Context, event *sspb.UsersEventByDaysRequest) ([]string, error) {
	// Канал для сбора результатов
	usersChan := make(chan *sspb.UsersEventByDaysResponse)
	// Канал для ошибок
	errChan := make(chan error, 1)

	stream, err := s.sspb.GetUsersFromEventByDaysLeft(ctx, event)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			user, err := stream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					close(usersChan)
					return
				}

				errChan <- err
				return
			}

			usersChan <- user
		}
	}()

	var userIds []string
	for {
		select {
		case err := <-errChan:
			return nil, err
		case user, ok := <-usersChan:
			if !ok {
				return userIds, nil
			}

			userIds = append(userIds, user.UserId)
		}
	}
}
