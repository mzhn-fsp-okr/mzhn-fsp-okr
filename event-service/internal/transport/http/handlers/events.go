package handlers

import (
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/internal/services/eventservice"
	"mzhn/event-service/pkg/sl"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type EventsResponse struct {
	Events []domain.EventInfo `json:"events"`
}

func Events(es *eventservice.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		ctx := c.Request().Context()

		res := EventsResponse{}

		chEvents := make(chan domain.EventInfo, 10)
		done := make(chan error, 1)

		go func() {
			done <- es.List(ctx, chEvents)
		}()

		go func() {
			for event := range chEvents {
				res.Events = append(res.Events, event)
			}
		}()

		if err := <-done; err != nil {
			log.Error("failed to list events", sl.Err(err))
			return err
		}

		return c.JSON(200, res)
	}
}
