package eventstorage

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"mzhn/event-service/internal/storage/model"
	"mzhn/event-service/internal/storage/pg"
	"mzhn/event-service/pkg/sl"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (s *Storage) Count(ctx context.Context, filters model.EventsFilters) (int64, error) {

	fn := "EventStorage.Count"
	log := s.l.With(sl.Method(fn), slog.Any("filters", filters))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", fn, err)
	}
	defer conn.Release()

	qb := sq.
		Select("count(*)").
		From(fmt.Sprintf("%s e", pg.EVENTS)).
		InnerJoin(fmt.Sprintf("%s sst on e.sport_subtype_id = sst.id", pg.SPORT_SUBTYPES)).
		InnerJoin(fmt.Sprintf("%s st on sst.sport_type_id = st.id", pg.SPORT_TYPES)).
		InnerJoin(fmt.Sprintf("%s ed on ed.event_id = e.id", pg.EVENT_DATES)).
		PlaceholderFormat(sq.Dollar)

	qb = s.applyListFilters(qb, filters)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build query", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("query", slog.String("sql", sql), slog.Any("args", args))

	var count int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		log.Error("failed to execute query", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", fn, err)
	}

	return count, nil
}
