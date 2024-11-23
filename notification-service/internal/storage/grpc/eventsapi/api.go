package eventsapi

import (
	"context"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/pb/espb"
	"mzhn/notification-service/pkg/sl"
	"time"

	"github.com/samber/lo"
)

type Api struct {
	l      *slog.Logger
	cfg    *config.Config
	client espb.EventServiceClient
}

func New(cfg *config.Config, client espb.EventServiceClient) *Api {
	return &Api{
		l:      slog.With(sl.Module("events-api")),
		cfg:    cfg,
		client: client,
	}
}

func (a *Api) Sport(ctx context.Context, id string) (*domain.SportTypeWithSubtypes, error) {
	fn := "events-api.Event"
	log := a.l.With(sl.Method(fn))

	log.Debug("try to find event", slog.String("id", id))
	res, err := a.client.Sport(ctx, &espb.SportRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	sp := res.SportType

	spsp := &domain.SportTypeWithSubtypes{
		Id:   sp.Id,
		Name: sp.Name,
		Subtypes: lo.Map(sp.Subtypes, func(st *espb.SportSubtype2, _ int) domain.SportSubtype2 {
			return domain.SportSubtype2{
				Id:   st.Id,
				Name: st.Name,
			}
		}),
	}

	return spsp, nil
}

func (a *Api) Event(ctx context.Context, id string) (*domain.EventInfo, error) {
	fn := "events-api.Event"
	log := a.l.With(sl.Method(fn))

	log.Debug("try to find event", slog.String("id", id))
	res, err := a.client.Event(ctx, &espb.EventRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	ev := res.Info

	from, err := time.Parse("02.01.2006", ev.Dates.DateFrom)
	if err != nil {
		log.Error("failed to parse date_from", sl.Err(err), slog.String("date_from", ev.Dates.DateFrom))
		return nil, err
	}
	to, err := time.Parse("02.01.2006", ev.Dates.DateTo)
	if err != nil {
		log.Error("failed to parse date_to", sl.Err(err), slog.String("date_to", ev.Dates.DateTo))
		return nil, err
	}

	e := &domain.EventInfo{
		Id:    ev.Id,
		EkpId: ev.EkpId,
		SportSubtype: domain.SportSubtype{
			Id:   ev.SportSubtype.Id,
			Name: ev.SportSubtype.Name,
			Parent: domain.SportType{
				Id:   ev.SportSubtype.Parent.Id,
				Name: ev.SportSubtype.Parent.Name,
			},
		},
		Name:        ev.Name,
		Description: ev.Description,
		Dates: domain.DateRange{
			From: from,
			To:   to,
		},
		Location:     ev.Location,
		Participants: int(ev.Participants),
		ParticipantRequirements: lo.Map(ev.ParticipantRequirements, func(pr *espb.ParticipantRequirements, _ int) domain.ParticipantRequirements {
			return domain.ParticipantRequirements{
				Gender: pr.Gender,
				MinAge: pr.MinAge,
				MaxAge: pr.MaxAge,
			}
		}),
	}

	return e, nil
}
