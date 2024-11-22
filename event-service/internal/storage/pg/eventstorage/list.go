package eventstorage

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/internal/storage/pg"
	"mzhn/event-service/pkg/sl"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) List(ctx context.Context, chEvents chan<- domain.EventInfo) error {

	fn := "EventStorage.List"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}
	defer conn.Release()

	qb := sq.
		Select("e.id, e.ekp_id, st.sport_type, sst.sport_subtype, e.name, e.description, e.location, e.participants, ed.date_from, ed.date_to").
		From(fmt.Sprintf("%s e", pg.EVENTS)).
		InnerJoin(fmt.Sprintf("%s sst on e.sport_subtype_id = sst.id", pg.SPORT_SUBTYPES)).
		InnerJoin(fmt.Sprintf("%s st on sst.sport_type_id = st.id", pg.SPORT_TYPES)).
		InnerJoin(fmt.Sprintf("%s ed on ed.event_id = e.id", pg.EVENT_DATES)).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build query", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("executing query", slog.String("sql", sql), slog.Any("args", args))

	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		log.Error("failed to execute query", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	for rows.Next() {
		var e domain.EventInfo
		if err := rows.Scan(
			&e.Id,
			&e.EkpId,
			&e.SportType,
			&e.SportSubtype,
			&e.Name,
			&e.Description,
			&e.Location,
			&e.Participants,
			&e.Dates.From,
			&e.Dates.To,
		); err != nil {
			log.Error("failed to scan row", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}

		chEvents <- e
	}

	return nil
}
