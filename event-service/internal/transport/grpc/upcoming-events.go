package grpc

import (
	"log/slog"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/internal/storage/model"
	"mzhn/event-service/pb/espb"
	"mzhn/event-service/pkg/sl"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetUpcomingEvents implements espb.EventServiceServer.
func (s *Server) GetUpcomingEvents(req *emptypb.Empty, stream espb.EventService_GetUpcomingEventsServer) error {

	chEvents := make(chan domain.EventInfo, 2)
	done := make(chan error, 1)

	startDate := time.Now()
	endDate := startDate.AddDate(0, 1, 0)
	log.Info("get upcoming events", slog.Any("endDate", endDate))

	go func() {
		err := s.es.List(stream.Context(), chEvents, model.EventsFilters{
			StartDate: &startDate,
			EndDate:   &endDate,
		})

		log.Debug("list ended execution")

		if err != nil {
			log.Error("failed to list events", slog.Any("err", err))
		}

		done <- err
		close(chEvents)
	}()

	go func() {
		for {
			select {
			case <-stream.Context().Done():
				log.Debug("context done")
				done <- nil
				return

			case event, ok := <-chEvents:
				if !ok {
					log.Info("channel closed")
					done <- nil
					return
				}
				log.Debug("event recieved", slog.Any("event", event))

				daysLeft := time.Until(event.Dates.From).Hours() / 24

				response := &espb.UpcomingEventResponse{
					DaysLeft: uint32(daysLeft),
					Event: &espb.EventInfo{
						Id:          event.Id,
						EkpId:       event.EkpId,
						Name:        event.Name,
						Description: event.Description,
						SportSubtype: &espb.SportSubtype{
							Id:   event.SportSubtype.Id,
							Name: event.SportSubtype.Name,
							Parent: &espb.SportType{
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

				log.Info("streaming event", slog.String("eventId", event.Id))

				log.Debug("sending response", slog.Any("response", response))
				if err := stream.Send(response); err != nil {
					log.Error("failed to send response", slog.Any("err", err), slog.Any("response", response))
					done <- err
				}
			}
		}
	}()

	if err := <-done; err != nil {
		log.Error("error", sl.Err(err))
		return err
	}

	return nil
}
