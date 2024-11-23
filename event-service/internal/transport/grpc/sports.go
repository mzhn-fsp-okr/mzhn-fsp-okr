package grpc

import (
	"errors"
	"io"
	"log/slog"
	"mzhn/event-service/pb/espb"
	"mzhn/event-service/pkg/sl"

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
			SportSubType: &espb.SportSubtype{
				Id:   sp.Id,
				Name: sp.Name,
				Parent: &espb.SportType{
					Id:   sp.SportType.Id,
					Name: sp.SportType.Name,
				},
			},
		}

		log.Debug("sending sport response", slog.Any("response", response))
		if err := stream.Send(response); err != nil {
			log.Error("failed to send event response", sl.Err(err))
			return err
		}
	}
}
