package verificationstorage

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/notification-service/internal/storage/model"
	"mzhn/notification-service/internal/storage/pg"
	"mzhn/notification-service/pkg/sl"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
)

func (s *Storage) Find(ctx context.Context, userId string) (*model.Verification, error) {
	return s.find(ctx, userId)
}

func (s *Storage) find(ctx context.Context, userId string) (*model.Verification, error) {
	fn := "verificationstorage.find"
	log := s.l.With(sl.Module(fn))

	c, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer c.Release()

	qb := sq.
		Select("user_id", "token", "created_at").
		From(pg.VERIFICATIONS).
		Where(sq.Eq{"user_id": userId}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("executing query", slog.String("sql", sql), slog.Any("args", args))

	i := new(model.Verification)
	if err := c.
		QueryRow(ctx, sql, args...).
		Scan(
			&i.UserId,
			&i.Token,
			&i.CreatedAt,
		); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		log.Error("failed to execute query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return i, nil
}
