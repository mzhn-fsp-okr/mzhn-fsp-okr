package middleware

import (
	"errors"
	"log/slog"
	"mzhn/event-service/internal/config"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/internal/services/authservice"
	"mzhn/event-service/pkg/responses"
	"mzhn/event-service/pkg/sl"

	"github.com/labstack/echo/v4"
)

type RoleFunc func(roles ...domain.Role) echo.MiddlewareFunc

func RequireAuth(as *authservice.Service, cfg *config.Config) RoleFunc {
	return func(roles ...domain.Role) echo.MiddlewareFunc {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				slog.Debug("require auth check")

				token := c.Get(TOKEN)
				if token == nil {
					slog.Error("token not found")
					return responses.BadRequest(c, errors.New("token not found"))
				}

				ctx := c.Request().Context()

				user, err := as.Authenticate(ctx, token.(string), roles...)
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
