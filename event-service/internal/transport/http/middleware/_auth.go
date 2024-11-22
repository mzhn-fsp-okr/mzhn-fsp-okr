package middleware

import (
	"errors"
	"log/slog"
	"mzhn/event-service/internal/config"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/internal/domain/dto"
	"mzhn/event-service/internal/domain/entity"
	"mzhn/event-service/internal/services/authservice"
	"mzhn/event-service/pkg/responses"
	"mzhn/event-service/pkg/sl"

	"github.com/labstack/echo/v4"
)

type RoleFunc func(roles ...entity.Role) echo.MiddlewareFunc

func RequireAuth(as *authservice.AuthService, cfg *config.Config) RoleFunc {
	return func(roles ...entity.Role) echo.MiddlewareFunc {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				slog.Debug("require auth check")

				token := c.Get(TOKEN)
				if token == nil {
					slog.Error("token not found")
					return responses.BadRequest(c, errors.New("token not found"))
				}

				ctx := c.Request().Context()

				user, err := as.Authenticate(ctx, &dto.Authenticate{
					AccessToken: token.(string),
					Roles:       roles,
				})
				if err != nil {
					slog.Error("failed to authenticate token", sl.Err(err))

					if errors.Is(err, domain.ErrTokenInvalid) {
						return responses.Unauthorized(c)
					} else if errors.Is(err, domain.ErrTokenExpired) {
						return responses.Unauthorized(c)
					} else if errors.Is(err, domain.ErrUserNotFound) {
						return responses.Unauthorized(c)
					} else if errors.Is(err, domain.ErrInsufficientPermission) {
						return responses.Forbidden(c)
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
