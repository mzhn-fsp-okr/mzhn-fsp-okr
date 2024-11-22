package eventstorage

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/event-service/internal/storage/pg"
	"mzhn/event-service/pkg/sl"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (s *Storage) StaleAll(ctx context.Context) error {
	fn := "eventstorage.StaleAll"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}
	defer conn.Release()

	qb := sq.Update(pg.EVENT_STALE).Set("is_stale", true).PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build query", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("executing query", slog.String("sql", sql), slog.Any("args", args))

	if _, err := conn.Exec(ctx, sql, args...); err != nil {
		log.Error("failed to execute query", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) unstale(ctx context.Context, eventId string) error {
	fn := "eventstorage.unstale"
	log := s.l.With(sl.Method(fn))

	conn := ctx.Value("tx").(pgx.Tx)

	qb := sq.Update(pg.EVENT_STALE).Set("is_stale", false).Where(sq.Eq{"event_id": eventId}).PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build query", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("executing query", slog.String("sql", sql), slog.Any("args", args))

	if _, err := conn.Exec(ctx, sql, args...); err != nil {
		log.Error("failed to execute query", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
