package userservice

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"mzhn/auth/internal/domain"
	"mzhn/auth/internal/domain/entity"
	"mzhn/auth/internal/storage"
	"mzhn/auth/pkg/sl"
)

func (a *Service) Find(ctx context.Context, slug string) (*entity.User, error) {

	fn := "userService.Find"
	log := a.logger.With(sl.Method(fn))

	u, err := a.userProvider.Find(ctx, slug)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Debug("user not found", slog.String("id", slug))
			return nil, fmt.Errorf("%s: %w", fn, domain.ErrUserNotFound)
		}

		log.Debug("cannot provide user", sl.Err(err))
		return nil, err
	}
	log.Debug("found user", slog.Any("user", u))

	return nil, fmt.Errorf("%s: %w", fn, domain.ErrInsufficientPermission)
}
