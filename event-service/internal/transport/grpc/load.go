package grpc

import (
	"errors"
	"io"
	"log/slog"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/pb/espb"
	"mzhn/event-service/pkg/sl"
)

func (s *Server) Load(stream espb.EventService_LoadServer) error {
	fn := "grpc.Load"
	log := s.l.With(sl.Method(fn))
	loaded := 0

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

		log.Debug("loading event", slog.Any("event", req.Info))
		if _, err := s.es.Load(stream.Context(), &domain.EventLoadInfo{
			EkpId:        req.Info.EkpId,
			SportType:    req.Info.SportType,
			SportSubtype: req.Info.SportSubtype,
			Name:         req.Info.Name,
			Description:  req.Info.Description,
			Dates: domain.DateRange{
				From: req.Info.Dates.From,
				To:   req.Info.Dates.To,
			},
			Location:     req.Info.Location,
			Participants: int(req.Info.Participants),
		}); err != nil {
			log.Error("failed to load event", sl.Err(err))
			return err
		}

		loaded++
	}
}
