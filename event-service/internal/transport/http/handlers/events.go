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
	Page            *int     `query:"page"`
	PageSize        *int     `query:"page_size"`
	StartDate       *string  `query:"start_date"`
	EndDate         *string  `query:"end_date"`
	SportTypeId     []string `query:"sport_type_id"`
	SportSubtypeId  []string `query:"sport_subtype_id"`
	MinAge          *int     `query:"min_age"`
	MaxAge          *int     `query:"max_age"`
	Sex             *bool    `query:"sex"`
	MinParticipants *int     `query:"min_participants"`
	MaxParticipants *int     `query:"max_participants"`
	Location        *string  `query:"location"`
	Name            *string  `query:"name"`
}

type EventsResponse struct {
	Events []domain.EventInfo `json:"events"`
	Total  int64              `json:"total"`
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

		filters := model.EventsFilters{
			Pagination: new(model.Pagination),
		}

		if req.PageSize != nil {
			page := uint64(*req.PageSize)
			filters.Limit = page
		} else {
			filters.Limit = 100
		}

		if req.Page != nil {
			offset := (uint64(*req.Page) - 1) * filters.Limit
			filters.Offset = offset
		}

		if req.StartDate != nil {
			startDate, err := time.Parse("02.01.2006", *req.StartDate)
			if err != nil {
				log.Error("failed to parse start date", sl.Err(err))
				return echo.NewHTTPError(400, "bad start date")
			}
			filters.StartDate = &startDate
		}

		if req.EndDate != nil {
			endDate, err := time.Parse("02.01.2006", *req.EndDate)
			if err != nil {
				log.Error("failed to parse end date", sl.Err(err))
				return echo.NewHTTPError(400, "bad end date")
			}
			filters.EndDate = &endDate
		}

		filters.MinAge = req.MinAge
		filters.MaxAge = req.MaxAge
		filters.Sex = req.Sex
		filters.MinParticipants = req.MinParticipants
		filters.MaxParticipants = req.MaxParticipants
		filters.Location = req.Location
		filters.SportTypesId = req.SportTypeId
		filters.SportSubtypesId = req.SportSubtypeId
		filters.Name = req.Name

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

		count, err := es.Count(ctx, filters)
		if err != nil {
			log.Error("failed to count events", sl.Err(err))
			return err
		}
		res.Total = count

		return c.JSON(200, res)
	}
}
