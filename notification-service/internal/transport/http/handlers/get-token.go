package handlers

import (
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/internal/services/integrationservice"
	"mzhn/notification-service/internal/transport/http/middleware"
	"mzhn/notification-service/pkg/sl"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type GetTokenResponse struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expireAt"`
}

func GetToken(es *integrationservice.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		user, ok := c.Get(middleware.USER).(*domain.User)
		if !ok {
			return echo.NewHTTPError(401, "invalid user")
		}

		ctx := c.Request().Context()

		v, err := es.GetVerificationCode(ctx, &domain.VerificationCodeRequest{
			UserId: user.Id,
		})
		if err != nil {
			log.Error("failed to get token", sl.Err(err))
			return err
		}

		return c.JSON(200, &GetTokenResponse{
			Token:    v.Token,
			ExpireAt: v.ExpireAt,
		})
	}
}
