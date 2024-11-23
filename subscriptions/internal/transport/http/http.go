package http

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/subscriptions-service/internal/config"
	"mzhn/subscriptions-service/internal/services/authservice"
	"mzhn/subscriptions-service/internal/services/subscriptionservice"
	"mzhn/subscriptions-service/internal/transport/http/handlers"
	"mzhn/subscriptions-service/internal/transport/http/middleware"
	"mzhn/subscriptions-service/pkg/sl"
	"strings"

	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo

	cfg    *config.Config
	logger *slog.Logger

	as *authservice.Service
	ss *subscriptionservice.Service
}

func New(cfg *config.Config, as *authservice.Service, ss *subscriptionservice.Service) *Server {
	return &Server{
		Echo:   echo.New(),
		logger: slog.Default().With(sl.Module("http")),
		cfg:    cfg,
		as:     as,
		ss:     ss,
	}
}

func (h *Server) setup() {
	h.Use(emw.Logger())
	h.Use(emw.CORSWithConfig(emw.CORSConfig{
		AllowOrigins:     strings.Split(h.cfg.Http.Cors.AllowedOrigins, ","),
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowCredentials: true,
	}))
	h.Validator = &CustomValidator{}

	tokguard := middleware.Token()
	authguard := middleware.RequireAuth(h.as, h.cfg)
	sport := h.Group("/sport", tokguard(), authguard())
	sport.POST("/subscribe", handlers.SubscribeToSport(h.ss))
	sport.POST("/unsubscribe", handlers.UnsubscribeFromSport(h.ss))

	event := h.Group("/event", tokguard(), authguard())
	event.GET("/", handlers.GetUserEvents(h.ss))
	event.POST("/subscribe", handlers.SubscribeToEvent(h.ss))
	event.POST("/unsubscribe", handlers.UnsubscribeFromEvent(h.ss))
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
