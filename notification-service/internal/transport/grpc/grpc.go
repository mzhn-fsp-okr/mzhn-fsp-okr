package grpc

import (
	"context"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/internal/services/integrationservice"
	"mzhn/notification-service/pb/nspb"
	"mzhn/notification-service/pkg/sl"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
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
	if err := s.is.Save(ctx, &domain.SetIntegrations{
		UserId:           req.UserId,
		TelegramUsername: &req.TelegramUsername,
	}); err != nil {
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
