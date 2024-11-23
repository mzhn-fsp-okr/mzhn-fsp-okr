package grpc

import (
	"errors"
	"io"
	"log/slog"
	"mzhn/event-service/internal/storage/model"
	"mzhn/event-service/pb/espb"
	"mzhn/event-service/pkg/sl"

	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Events implements espb.EventServiceServer.
func (s *Server) Sports(stream espb.EventService_SportsServer) error {
	fn := "grpc.Events"
	log := s.l.With(sl.Method(fn))

	for {
		req, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			log.Error("failed to receive event request", sl.Err(err))
			return err
		}

		log.Debug("received event request", slog.Any("request", req))

		sp, err := s.ss.Find(stream.Context(), req.Id)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			log.Error("failed to find event", sl.Err(err))
			return err
		}

		if sp == nil {
			return status.Error(codes.NotFound, "event not found")
		}

		response := &espb.SportResponse{
			SportType: &espb.SportTypeWithSubtypes{
				Id:   sp.SportType.Id,
				Name: sp.SportType.Name,
				Subtypes: lo.Map(sp.Subtypes, func(sst model.Subtype, _ int) *espb.SportSubtype2 {
					return &espb.SportSubtype2{
						Id:   sst.Id,
						Name: sst.Name,
					}
				}),
			},
		}

		log.Debug("sending sport response", slog.Any("response", response))
		if err := stream.Send(response); err != nil {
			log.Error("failed to send event response", sl.Err(err))
			return err
		}
	}
}
