package http

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/services/authservice"
	"mzhn/notification-service/internal/services/integrationservice"
	"mzhn/notification-service/internal/transport/http/handlers"
	"mzhn/notification-service/internal/transport/http/middleware"
	"mzhn/notification-service/pkg/sl"
	"strings"

	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo

	cfg    *config.Config
	logger *slog.Logger

	as *authservice.Service
	is *integrationservice.Service
}

func New(cfg *config.Config, as *authservice.Service, is *integrationservice.Service) *Server {
	return &Server{
		Echo:   echo.New(),
		logger: slog.Default().With(sl.Module("http")),
		cfg:    cfg,
		as:     as,
		is:     is,
	}
}

func (h *Server) setup() {

	h.Use(emw.Logger())
	h.Use(emw.CORSWithConfig(emw.CORSConfig{
		AllowOrigins:     strings.Split(h.cfg.Http.Cors.AllowedOrigins, ","),
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowCredentials: true,
	}))

	tokguard := middleware.Token()
	authguard := middleware.RequireAuth(h.as, h.cfg)

	h.GET("/", handlers.ProfileIntegrations(h.is), tokguard(), authguard())
	h.PUT("/", handlers.UpdateIntegrations(h.is), tokguard(), authguard())
}

func (h *Server) Run(ctx context.Context) error {
	h.setup()

	host := h.cfg.Http.Host
	port := h.cfg.Http.Port
	addr := fmt.Sprintf("%s:%d", host, port)
	slog.Info("running http server", slog.String("addr", addr))

	go func() {
		if err := h.Start(addr); err != nil {
			return
		}
	}()

	<-ctx.Done()
	if err := h.Shutdown(ctx); err != nil {
		return err
	}

	slog.Info("shutting down http server\n")
	return nil
}
