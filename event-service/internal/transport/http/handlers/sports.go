package handlers

import (
	"mzhn/event-service/internal/services/sportservice"
	"mzhn/event-service/internal/storage/model"
	"mzhn/event-service/pkg/sl"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type SportsListResponse struct {
	Sports []model.SportTypeWithSubtypes `json:"sportTypes"`
	Total  int64                         `json:"total"`
}

func Sports(ss *sportservice.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		ctx := c.Request().Context()

		// req := EventsRequest{}
		// binder := &echo.DefaultBinder{}
		// if err := binder.BindQueryParams(c, &req); err != nil {
		// 	log.Error("failed to bind events request", sl.Err(err))
		// 	return echo.NewHTTPError(500, err.Error())
		// }

		// slog.Info("list events", slog.Any("filters", req))

		res := SportsListResponse{}

		ch := make(chan model.SportTypeWithSubtypes, 10)
		done := make(chan error, 1)

		go func() {
			done <- ss.List(ctx, ch)
		}()

		go func() {
			for event := range ch {
				res.Sports = append(res.Sports, event)
			}
		}()

		if err := <-done; err != nil {
			log.Error("failed to list events", sl.Err(err))
			return err
		}

		return c.JSON(200, res)
	}
}
