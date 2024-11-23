package handlers

import (
	"log/slog"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/internal/services/eventservice"
	"mzhn/event-service/pkg/sl"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type EventRequest struct {
	Id string `param:"id"`
}

type EventResponse struct {
	Event *domain.EventInfo `json:"events"`
}

func Event(es *eventservice.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := EventRequest{}
		ctx := c.Request().Context()

		if err := echo.PathParamsBinder(c).
			String("id", &req.Id).
			BindError(); err != nil {
			log.Error("failed to bind events request", sl.Err(err))
			return echo.NewHTTPError(500, err.Error())
		}

		if _, err := uuid.Parse(req.Id); err != nil {
			log.Error("failed to parse event id", sl.Err(err), slog.String("id", req.Id))
			return echo.NewHTTPError(400, "event id must be uuid")
		}

		event, err := es.Find(ctx, req.Id)
		if err != nil {
			log.Error("failed to find event", sl.Err(err), slog.String("id", req.Id))
			return echo.NewHTTPError(500, err.Error())
		}

		if event == nil {
			return echo.NewHTTPError(404, "event not found")
		}

		return c.JSON(200, &EventResponse{
			Event: event,
		})
	}
}
