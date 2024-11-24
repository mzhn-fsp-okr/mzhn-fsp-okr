package handlers

import (
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/internal/services/integrationservice"
	"mzhn/notification-service/internal/transport/http/middleware"

	"github.com/labstack/echo/v4"
)

type UpdateIntegrationsRequest struct {
	TelegramUsername *string `json:"telegramUsername"`
	WannaMail        *bool   `json:"wannaMail"`
}

func UpdateIntegrations(es *integrationservice.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		var r UpdateIntegrationsRequest

		if err := c.Bind(&r); err != nil {
			return err
		}

		user, ok := c.Get(middleware.USER).(*domain.User)
		if !ok {
			return echo.NewHTTPError(401, "invalid user")
		}

		ctx := c.Request().Context()

		es.Save(ctx, &domain.SetIntegrations{
			UserId:           user.Id,
			TelegramUsername: r.TelegramUsername,
			WannaMail:        r.WannaMail,
		})

		return c.NoContent(200)
	}
}
