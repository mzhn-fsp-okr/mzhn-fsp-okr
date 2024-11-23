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
	// Канал для сбора результатов
	eventsChan := make(chan *espb.UpcomingEventResponse)
	// Канал для ошибок
	errChan := make(chan error, 1)

	stream, err := s.espb.GetUpcomingEvents(context.Background(), &emptypb.Empty{})
	if err != nil {
		s.l.Error("cannot get upcoming events", sl.Err(err))
		return err
	}

	go func() {
		for {
			event, err := stream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					close(eventsChan)
					return
				}

				s.l.Error("Error whiel Recv()", slog.Any("error", err))
				errChan <- err
				return
			}

			eventsChan <- event
		}
	}()

	for {
		select {
		case err := <-errChan:
			return err
		case event, ok := <-eventsChan:
			if !ok {
				return nil
			}

			go s.proccessEvent(event)
		}
	}

}

func (s *Service) proccessEvent(event *espb.UpcomingEventResponse) error {
	var daysLeft sspb.DaysLeft

	if event.DaysLeft <= 30 && event.DaysLeft > 7 {
		daysLeft = sspb.DaysLeft_Month
	} else if event.DaysLeft <= 7 && event.DaysLeft > 3 {
		daysLeft = sspb.DaysLeft_Week
	} else {
		daysLeft = sspb.DaysLeft_ThreeDays
	}

	userIds, err := s.getUsersToNotify(context.Background(), &sspb.UsersEventByDaysRequest{
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
		go s.pub.NotifyAboutUpcomingEvent(context.Background(), userId, event.Event.Id, event.DaysLeft, daysLeft)
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
