package handlers

import (
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/internal/services/integrationservice"
	"mzhn/notification-service/internal/transport/http/middleware"
	"mzhn/notification-service/pkg/sl"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ProfileIntegrationsResponse struct {
	Telegram  *string `json:"telegram"`
	WannaMail bool    `json:"wannaMail"`
}

func ProfileIntegrations(es *integrationservice.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		user, ok := c.Get(middleware.USER).(*domain.User)
		if !ok {
			return echo.NewHTTPError(401, "invalid user")
		}

		ctx := c.Request().Context()

		i, err := es.Find(ctx, user.Id)
		if err != nil {
			log.Error("failed to find integration", sl.Err(err))
			return err
		}

		return c.JSON(200, &ProfileIntegrationsResponse{
			Telegram:  i.TelegramUsername,
			WannaMail: i.WannaMail,
		})
	}
}
