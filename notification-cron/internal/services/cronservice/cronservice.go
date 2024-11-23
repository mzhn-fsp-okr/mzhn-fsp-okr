package cronservice

import (
	"context"
	"log/slog"
	"mzhn/notification-cron/internal/config"
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
}

func New(cfg *config.Config, sspb sspb.SubscriptionServiceClient, espb espb.EventServiceClient) *Service {
	return &Service{
		l:   slog.With(sl.Module("CronService")),
		cfg: cfg,

		sspb: sspb,
		espb: espb,
	}
}

func (s *Service) NotifyUsers(ctx context.Context, daysLeft uint32) error {
	if daysLeft > 30 {
		s.l.Debug("days left more than 30")
		return nil
	}

	stream, err := s.espb.GetUpcomingEvents(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	go func ()  {
		for {
			event, err := stream.Recv()
		}
	}


	if daysLeft <= 30 && daysLeft > 7 {

	} else if daysLeft <= 7 && daysLeft > 3 {
	} else {

	}

}
