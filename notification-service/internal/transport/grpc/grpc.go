package grpc

import (
	"context"
	"errors"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/internal/services/integrationservice"
	"mzhn/notification-service/pb/nspb"
	"mzhn/notification-service/pkg/sl"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var _ nspb.NotificationServiceServer = (*Server)(nil)

type Server struct {
	cfg *config.Config
	l   *slog.Logger
	is  *integrationservice.Service
	*nspb.UnimplementedNotificationServiceServer
}

// LinkTelegram implements nspb.NotificationServiceServer.
func (s *Server) LinkTelegram(ctx context.Context, req *nspb.LinkTelegramRequest) (*nspb.LinkTelegramResponse, error) {

	if err := s.is.LinkTelegram(ctx, &domain.LinkTelegramRequest{
		UserId:           req.UserId,
		Token:            req.Token,
		TelegramUsername: req.TelegramUsername,
	}); err != nil {
		if errors.Is(err, domain.ErrVerificationExpired) {
			return nil, status.Errorf(codes.ResourceExhausted, "verification expired")
		}
		if errors.Is(err, domain.ErrVerificationInvalidToken) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid token")
		}
		if errors.Is(err, domain.ErrVerificationNotFound) {
			return nil, status.Errorf(codes.NotFound, "verification not found")
		}
		return nil, err
	}

	return &nspb.LinkTelegramResponse{}, nil
}

func New(cfg *config.Config, is *integrationservice.Service) *Server {
	return &Server{
		cfg: cfg,
		l:   slog.With(sl.Module("grpc")),
		is:  is,
	}
}

func (s *Server) Run(ctx context.Context) error {

	log := slog.With(sl.Module("grpc"))

	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	addr := s.cfg.Grpc.Addr()

	log.Info("starting grpc server", slog.String("addr", addr))

	if s.cfg.Grpc.UseReflection {
		log.Info("enabling reflection")
		reflection.Register(server)
	}

	nspb.RegisterNotificationServiceServer(server, s)

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
