package grpc

import (
	"log/slog"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/internal/storage/model"
	"mzhn/event-service/pb/espb"
	"mzhn/event-service/pkg/sl"
	"time"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetUpcomingEvents implements espb.EventServiceServer.
func (s *Server) GetUpcomingEvents(req *emptypb.Empty, stream espb.EventService_GetUpcomingEventsServer) error {

	chEvents := make(chan domain.EventInfo, 10)
	done := make(chan error, 1)

	endDate := time.Now().AddDate(0, 1, 0)

	go func() {
		done <- s.es.List(stream.Context(), chEvents, model.EventsFilters{
			EndDate: &endDate,
		})

		close(chEvents)
	}()

	go func() {

		for {

			event, ok := <-chEvents
			if !ok {
				done <- nil
				return
			}
			s.l.Debug("event recieved", slog.Any("event", event))

			response := &espb.UpcomingEventResponse{
				Event: &espb.EventInfo{
					Id:          event.Id,
					EkpId:       event.EkpId,
					Name:        event.Name,
					Description: event.Description,
					SportSubtype: &espb.SportSubtype{
						Id:   event.SportSubtype.Id,
						Name: event.SportSubtype.Name,
						Parent: &espb.SportSubtype{
							Id:   event.SportSubtype.Parent.Id,
							Name: event.SportSubtype.Parent.Name,
						},
					},
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
				done <- err
			}
		}
	}()

	if err := <-done; err != nil {
		s.l.Error("error", sl.Err(err))
		return err
	}

	return nil
}
