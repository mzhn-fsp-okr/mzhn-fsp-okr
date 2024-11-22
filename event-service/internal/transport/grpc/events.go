package grpc

import (
	"errors"
	"io"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/pb/espb"
	"mzhn/event-service/pkg/sl"

	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Events implements espb.EventServiceServer.
func (s *Server) Events(stream espb.EventService_EventsServer) error {
	fn := "grpc.Events"
	log := s.l.With(sl.Method(fn))

	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		event, err := s.es.Find(stream.Context(), req.Id)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			log.Error("failed to find event", sl.Err(err))
			return err
		}

		if event == nil {
			return status.Error(codes.NotFound, "event not found")
		}

		response := &espb.EventResponse{
			Info: &espb.EventInfo{
				Id:           event.Id,
				EkpId:        event.EkpId,
				SportType:    event.SportType,
				SportSubtype: event.SportSubtype,
				Name:         event.Name,
				Description:  event.Description,
				Dates: &espb.DateRange{
					DateFrom: event.Dates.From.Format("02.01.2006"),
					DateTo:   event.Dates.To.Format("02.01.2006"),
				},
				Location:     event.Location,
				Participants: int32(event.Participants),
				ParticipantRequirements: lo.Map(event.ParticipantRequirements, func(pr domain.ParticipantRequirements, _ int) *espb.ParticipantRequirements {
					return &espb.ParticipantRequirements{
						Gender: pr.Gender,
						MinAge: pr.MinAge,
						MaxAge: pr.MaxAge,
					}
				}),
			},
		}
		if err := stream.Send(response); err != nil {
			return err
		}
	}
}
