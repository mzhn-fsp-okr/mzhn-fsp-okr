package handlers

import (
	"mzhn/notification-service/internal/services/integrationservice"

	"github.com/labstack/echo/v4"
)

type UpdateIntegrationsRequest struct {
	TelegramUsername *string `json:"telegramUsername"`
	WannaMail        *bool   `json:"wannaMail"`
}

func UpdateIntegrations(es *integrationservice.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(200)
	}
}
