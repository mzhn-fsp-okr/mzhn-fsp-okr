package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/event-service/internal/config"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/internal/services/eventservice"
	"mzhn/event-service/internal/services/sportservice"
	"mzhn/event-service/pb/espb"
	"mzhn/event-service/pkg/sl"
	"net"

	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var _ espb.EventServiceServer = (*Server)(nil)

type Server struct {
	cfg *config.Config
	l   *slog.Logger
	es  *eventservice.Service
	ss  *sportservice.Service
	*espb.UnimplementedEventServiceServer
}

// Event implements espb.EventServiceServer.
func (s *Server) Event(ctx context.Context, in *espb.EventRequest) (*espb.EventResponse, error) {

	event, err := s.es.Find(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if event == nil {
		return nil, status.Errorf(codes.NotFound, "not found")
	}

	response := &espb.EventResponse{
		Info: &espb.EventInfo{
			Id:    event.Id,
			EkpId: event.EkpId,
			SportSubtype: &espb.SportSubtype{
				Id:   event.SportSubtype.Id,
				Name: event.SportSubtype.Name,
				Parent: &espb.SportType{
					Id:   event.SportSubtype.Parent.Id,
					Name: event.SportSubtype.Parent.Name,
				},
			},
			Name:        event.Name,
			Description: event.Description,
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

	return response, nil
}

func New(cfg *config.Config, es *eventservice.Service, ss *sportservice.Service) *Server {
	return &Server{
		cfg: cfg,
		l:   slog.With(sl.Module("grpc")),
		es:  es,
		ss:  ss,
	}
}

func (s *Server) Run(ctx context.Context) error {

	log := slog.With(sl.Module("grpc"))

	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	host := s.cfg.Grpc.Host
	port := s.cfg.Grpc.Port
	addr := fmt.Sprintf("%s:%d", host, port)

	log.Info("starting grpc server", slog.String("addr", addr))

	if s.cfg.Grpc.UseReflection {
		log.Info("enabling reflection")
		reflection.Register(server)
	}

	espb.RegisterEventServiceServer(server, s)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		slog.Error("failed to bind port", slog.String("addr", addr), sl.Err(err))
		return err
	}

	go func() {
		if err := server.Serve(listener); err != nil {
			slog.Error("failed to serve", sl.Err(err))
			return
		}
	}()

	<-ctx.Done()
	log.Info("shutting down grpc server")
	server.GracefulStop()
	return nil
}
