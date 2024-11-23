package handlers

import (
	"log/slog"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/internal/services/eventservice"
	"mzhn/event-service/internal/storage/model"
	"mzhn/event-service/pkg/sl"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type EventsRequest struct {
	Page      int    `query:"page"`
	PageSize  int    `query:"page_size"`
	StartDate string `query:"start_date"`
	EndDate   string `query:"end_date"`
}

type EventsResponse struct {
	Events []domain.EventInfo `json:"events"`
}

func Events(es *eventservice.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := EventsRequest{}
		ctx := c.Request().Context()

		binder := &echo.DefaultBinder{}
		if err := binder.BindQueryParams(c, &req); err != nil {
			log.Error("failed to bind events request", sl.Err(err))
			return echo.NewHTTPError(500, err.Error())
		}

		slog.Info("list events", slog.Any("filters", req))

		filters := model.EventsFilters{}

		if req.Page != 0 && req.PageSize != 0 {
			limit := uint64(req.PageSize)
			filters.Limit = &limit

			offset := (uint64(req.Page) - 1) * limit
			filters.Offset = &offset
		}

		if req.StartDate != "" {
			startDate, err := time.Parse("02.01.2006", req.StartDate)
			if err != nil {
				log.Error("failed to parse start date", sl.Err(err))
				return echo.NewHTTPError(400, "bad start date")
			}
			filters.StartDate = &startDate
		}

		if req.EndDate != "" {
			endDate, err := time.Parse("02.01.2006", req.EndDate)
			if err != nil {
				log.Error("failed to parse end date", sl.Err(err))
				return echo.NewHTTPError(400, "bad end date")
			}
			filters.EndDate = &endDate
		}

		res := EventsResponse{}

		chEvents := make(chan domain.EventInfo, 10)
		done := make(chan error, 1)

		go func() {
			done <- es.List(ctx, chEvents, filters)
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
