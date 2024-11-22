package eventstorage

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/internal/storage/pg"
	"mzhn/event-service/pkg/sl"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Storage) Find(ctx context.Context, id string) (*domain.EventInfo, error) {

	fn := "EventStorage.Find"
	log := r.l.With(sl.Method(fn))

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	qb := sq.
		Select("e.id, e.ekp_id, e.name, e.description, st.sport_type, sst.sport_subtype, ed.date_from, ed.date_to, e.location, e.participants").
		From(fmt.Sprintf("%s e", pg.EVENTS)).
		InnerJoin(fmt.Sprintf("%s ed e.id = ed.id", pg.EVENT_DATES)).
		InnerJoin(fmt.Sprintf("%s sst e.sport_subtype_id = sst.id", pg.SPORT_SUBTYPES)).
		InnerJoin(fmt.Sprintf("%s st sst.sport_type_id = st.id", pg.SPORT_TYPES)).
		PlaceholderFormat(sq.Dollar)

	if _, err := uuid.Parse(id); err != nil {
		qb = qb.Where(sq.Eq{"e.ekp_id": id})
	} else {
		qb = qb.Where(sq.Eq{"e.id": id})
	}

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build sql", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("executing query", slog.String("sql", sql), slog.Any("args", args))

	e := new(domain.EventInfo)
	if err := conn.QueryRow(ctx, sql, args...).Scan(
		&e.Id,
		&e.EkpId,
		&e.Name,
		&e.Description,
		&e.SportType,
		&e.SportSubtype,
		&e.Dates.From,
		&e.Dates.To,
		&e.Location,
		&e.Participants,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		log.Error("failed to execute query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return e, nil
}
