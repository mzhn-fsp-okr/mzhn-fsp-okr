package middleware

import (
	"errors"
	"log/slog"
	"mzhn/subscriptions-service/internal/config"
	"mzhn/subscriptions-service/internal/domain"
	"mzhn/subscriptions-service/internal/services/authservice"
	"mzhn/subscriptions-service/pkg/responses"
	"mzhn/subscriptions-service/pkg/sl"

	"github.com/labstack/echo/v4"
)

type RoleFunc func(roles ...domain.Role) echo.MiddlewareFunc

func RequireAuth(as *authservice.Service, cfg *config.Config) RoleFunc {
	return func(roles ...domain.Role) echo.MiddlewareFunc {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				slog.Debug("require auth check")

				token, ok := c.Get(TOKEN).(string)
				if !ok {
					slog.Error("token not found")
					return responses.BadRequest(c, errors.New("token not found"))

				}

				ctx := c.Request().Context()

				user, err := as.Profile(ctx, token)
				if err != nil {
					slog.Error("failed to authenticate token", sl.Err(err))

					if errors.Is(err, domain.ErrUnathorized) {
						return responses.Unauthorized(c)
					}

					return responses.Internal(c, err)
				}

				slog.Debug("user authenticated", slog.Any("user", user))
				c.Set(USER, user)

				return next(c)
			}
		}
	}
}
