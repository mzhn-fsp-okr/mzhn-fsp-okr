package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/event-service/internal/config"
	"mzhn/event-service/internal/services/eventservice"
	"mzhn/event-service/pb/espb"
	"mzhn/event-service/pkg/sl"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var _ espb.EventServiceServer = (*Server)(nil)

type Server struct {
	*espb.UnimplementedEventServiceServer
	cfg *config.Config
	l   *slog.Logger
	es  *eventservice.Service
}

func New(cfg *config.Config, es *eventservice.Service) *Server {
	return &Server{
		cfg: cfg,
		l:   slog.With(sl.Module("grpc")),
		es:  es,
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
