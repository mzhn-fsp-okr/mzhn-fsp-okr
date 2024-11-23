package integrationstorage

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/internal/storage/pg"
	"mzhn/notification-service/pkg/sl"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (s *Storage) find(ctx context.Context, userId string) (*domain.Integrations, error) {
	fn := "integrationstorage.find"
	log := s.l.With(sl.Module(fn))

	c, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer c.Release()

	qb := sq.
		Select("user_id", "telegram_username", "wanna_mail").
		From(pg.INTEGRATIONS).
		Where(sq.Eq{"user_id": userId}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("executing query", slog.String("sql", sql), slog.Any("args", args))

	i := &domain.Integrations{}
	if err := c.
		QueryRow(ctx, sql, args...).
		Scan(
			&i.UserId,
			&i.TelegramUsername,
			&i.WannaMail,
		); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		log.Error("failed to execute query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return i, nil
}
