package middleware

import (
	"errors"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/internal/services/authservice"
	"mzhn/notification-service/pkg/responses"
	"mzhn/notification-service/pkg/sl"

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
					return responses.Unauthorized(c)
				}

				slog.Debug("user authenticated", slog.Any("user", user))
				c.Set(USER, user)

				return next(c)
			}
		}
	}
}
