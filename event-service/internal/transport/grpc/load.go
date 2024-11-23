package grpc

import (
	"errors"
	"io"
	"log/slog"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/pb/espb"
	"mzhn/event-service/pkg/sl"
	"time"

	"github.com/samber/lo"
)

func (s *Server) Load(stream espb.EventService_LoadServer) error {
	fn := "grpc.Load"
	log := s.l.With(sl.Method(fn))
	loaded := 0

	if err := s.es.Stale(stream.Context()); err != nil {
		log.Error("failed to mark events as stale", sl.Err(err))
		return err
	}

	for {
		req, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				stream.SendAndClose(&espb.LoadResponse{
					Saved: int32(loaded),
				})
				return nil
			}
			log.Error("failed to recv", sl.Err(err))
			return err
		}

		startDate, err := time.Parse("02.01.2006", req.Dates.DateFrom)
		if err != nil {
			log.Error("failed to parse date", sl.Err(err))
			return err
		}

		endDate, err := time.Parse("02.01.2006", req.Dates.DateTo)
		if err != nil {
			log.Error("failed to parse date", sl.Err(err))
			return err
		}

		log.Debug("loading event", slog.Any("event", req))
		if _, err := s.es.Load(stream.Context(), &domain.EventLoadInfo{
			EkpId:        req.EkpId,
			SportType:    req.SportType,
			SportSubtype: req.SportSubtype,
			Name:         req.Name,
			Description:  req.Description,
			Dates: domain.DateRange{
				From: startDate,
				To:   endDate,
			},
			Location:     req.Location,
			Participants: int(req.Participants),
			ParticipantRequirements: lo.Map(req.ParticipantRequirements, func(r *espb.ParticipantRequirements, _ int) domain.ParticipantRequirements {
				pr := domain.ParticipantRequirements{
					Gender: r.Gender,
					MinAge: r.MinAge,
					MaxAge: r.MaxAge,
				}

				return pr
			}),
		}); err != nil {
			log.Error("failed to load event", sl.Err(err))
			return err
		}

		loaded++
	}
}
