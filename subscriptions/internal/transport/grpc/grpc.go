package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/subscriptions-service/internal/config"
	"mzhn/subscriptions-service/internal/domain"
	"mzhn/subscriptions-service/pb/sspb"
	"mzhn/subscriptions-service/pkg/sl"
	"net"

	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ sspb.SubscriptionServiceServer = (*Server)(nil)

type Server struct {
	cfg *config.Config
	l   *slog.Logger
	ss  domain.SubscriptionsService
	sspb.UnimplementedSubscriptionServiceServer
}

func New(cfg *config.Config, ss domain.SubscriptionsService) *Server {
	return &Server{
		cfg: cfg,
		l:   slog.With(sl.Module("grpc")),
		ss:  ss,
	}
}

func (s *Server) Run(ctx context.Context) error {
	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	host := s.cfg.Grpc.Host
	port := s.cfg.Grpc.Port
	addr := fmt.Sprintf("%s:%d", host, port)

	s.l.Info("starting grpc server", slog.String("addr", addr))

	if s.cfg.Grpc.UseReflection {
		log.Info("enabling reflection")
		reflection.Register(server)
	}

	sspb.RegisterSubscriptionServiceServer(server, s)

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

// GetUsersSubscribedToEvent implements sspb.SubscriptionServiceServer.
func (s *Server) GetUsersSubscribedToEvent(req *sspb.SubscriptionRequest, stream sspb.SubscriptionService_GetUsersSubscribedToEventServer) error {
	userIds, err := s.ss.GetUsersSubscribedToEvent(req.Id)
	if err != nil {
		return err
	}

	for _, userId := range userIds {
		if err := stream.Send(&sspb.SubscriptionResponse{UserId: userId}); err != nil {
			return err
		}
	}

	return nil
}

// GetUsersSubscribedToSport implements sspb.SubscriptionServiceServer.
func (s *Server) GetUsersSubscribedToSport(req *sspb.SubscriptionRequest, stream sspb.SubscriptionService_GetUsersSubscribedToSportServer) error {
	userIds, err := s.ss.GetUsersSubscribedToSport(req.Id)
	if err != nil {
		return err
	}

	for _, userId := range userIds {
		if err := stream.Send(&sspb.SubscriptionResponse{UserId: userId}); err != nil {
			return err
		}
	}

	return nil
}

// GetUsersFromEventByDaysLeft implements sspb.SubscriptionServiceServer.
func (s *Server) GetUsersFromEventByDaysLeft(req *sspb.UsersEventByDaysRequest, stream grpc.ServerStreamingServer[sspb.UsersEventByDaysResponse]) error {
	userIds, err := s.ss.GetUsersFromEventByDaysLeft(req.EventId, req.DaysLeft)
	if err != nil {
		return err
	}

	for _, userId := range userIds {
		if err := stream.Send(&sspb.UsersEventByDaysResponse{UserId: userId}); err != nil {
			return err
		}
	}

	return nil

}

// NotifyUser implements sspb.SubscriptionServiceServer.
func (s *Server) NotifyUser(ctx context.Context, req *sspb.NotifyUserRequest) (*emptypb.Empty, error) {
	if err := s.ss.NotifyUser(req.UserId, req.DaysLeft); err != nil {
		return nil, err
	}

	return nil, nil
}
