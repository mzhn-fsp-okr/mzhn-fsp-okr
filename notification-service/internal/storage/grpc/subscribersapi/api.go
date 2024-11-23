package subscribersapi

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/pb/sspb"
	"mzhn/notification-service/pkg/sl"
)

type Api struct {
	l      *slog.Logger
	cfg    *config.Config
	client sspb.SubscriptionServiceClient
}

func New(cfg *config.Config, client sspb.SubscriptionServiceClient) *Api {
	return &Api{
		l:      slog.With(sl.Module("subscribers-api")),
		cfg:    cfg,
		client: client,
	}
}

func (a *Api) SportSubs(ctx context.Context, sportId string) ([]string, error) {
	stream, err := a.client.GetUsersSubscribedToSport(ctx, &sspb.SubscriptionRequest{
		Id: sportId,
	})
	if err != nil {
		return nil, err
	}

	idid := make([]string, 0)
	for {
		sub, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		idid = append(idid, sub.UserId)
	}

	return idid, nil
}

func (a *Api) EventSubscribers(ctx context.Context, eventId string) ([]string, error) {
	stream, err := a.client.GetUsersSubscribedToEvent(ctx, &sspb.SubscriptionRequest{
		Id: eventId,
	})
	if err != nil {
		return nil, err
	}

	idid := make([]string, 0)
	for {
		sub, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		idid = append(idid, sub.UserId)
	}

	return idid, nil
}
