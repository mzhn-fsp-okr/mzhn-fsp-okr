package http

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/event-service/internal/config"
	"mzhn/event-service/internal/services/authservice"
	"mzhn/event-service/internal/services/eventservice"
	"mzhn/event-service/internal/transport/http/handlers"
	"mzhn/event-service/pkg/sl"
	"strings"

	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo

	cfg    *config.Config
	logger *slog.Logger

	as *authservice.Service
	es *eventservice.Service
}

func New(cfg *config.Config, as *authservice.Service, es *eventservice.Service) *Server {
	return &Server{
		Echo:   echo.New(),
		logger: slog.Default().With(sl.Module("http")),
		cfg:    cfg,
		as:     as,
		es:     es,
	}
}

func (h *Server) setup() {

	h.Use(emw.Logger())
	h.Use(emw.CORSWithConfig(emw.CORSConfig{
		AllowOrigins:     strings.Split(h.cfg.Http.Cors.AllowedOrigins, ","),
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowCredentials: true,
	}))

	// tokguard := middleware.Token()
	// authguard := middleware.RequireAuth(h.as, h.cfg)

	h.GET("/", handlers.Events(h.es) /*, tokguard(), authguard()*/)
	h.GET("/:id", handlers.Event(h.es) /*, tokguard(), authguard()*/)
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
