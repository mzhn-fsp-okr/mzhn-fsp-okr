package eventstorage

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/event-service/internal/domain"
	"mzhn/event-service/internal/storage/model"
	"mzhn/event-service/internal/storage/pg"
	"mzhn/event-service/pkg/sl"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) List(ctx context.Context, chEvents chan<- domain.EventInfo, filters model.EventsFilters) error {

	fn := "EventStorage.List"
	log := s.l.With(sl.Method(fn), slog.Any("filters", filters))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}
	defer conn.Release()

	ctx = context.WithValue(ctx, "tx", conn)

	qb := sq.
		Select("e.id, e.ekp_id, st.sport_type, sst.sport_subtype, e.name, e.description, e.location, e.participants, ed.date_from, ed.date_to").
		From(fmt.Sprintf("%s e", pg.EVENTS)).
		InnerJoin(fmt.Sprintf("%s sst on e.sport_subtype_id = sst.id", pg.SPORT_SUBTYPES)).
		InnerJoin(fmt.Sprintf("%s st on sst.sport_type_id = st.id", pg.SPORT_TYPES)).
		InnerJoin(fmt.Sprintf("%s ed on ed.event_id = e.id", pg.EVENT_DATES)).
		OrderBy("ed.date_from ASC").
		PlaceholderFormat(sq.Dollar)

	if filters.Limit != nil {
		qb = qb.Limit(*filters.Limit)
	}

	if filters.Offset != nil {
		qb = qb.Offset(*filters.Offset)
	}

	if filters.StartDate != nil {
		qb = qb.Where(sq.GtOrEq{"ed.date_from": *filters.StartDate})
	}

	if filters.EndDate != nil {
		qb = qb.Where(sq.LtOrEq{"ed.date_from": *filters.EndDate})
	}

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

		rr, err := s.listRequirementsFor(ctx, e.Id)
		if err != nil {
			log.Error("failed to list requirements for event", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}

		e.ParticipantRequirements = rr

		chEvents <- e
	}

	return nil
}

func (s *Storage) listRequirementsFor(ctx context.Context, eventId string) ([]domain.ParticipantRequirements, error) {

	fn := "EventStorage.listRequirementsFor"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer conn.Release()

	qb := sq.
		Select("epr.gender, epr.min_age, epr.max_age").
		From(fmt.Sprintf("%s epr", pg.EVENT_PARTICIPANTS_REQUIREMENTS)).
		Where(sq.Eq{"epr.event_id": eventId}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("executing query", slog.String("sql", sql), slog.Any("args", args))

	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		log.Error("failed to execute query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	rr := make([]domain.ParticipantRequirements, 0, 16)
	for rows.Next() {
		var e domain.ParticipantRequirements
		if err := rows.Scan(
			&e.Gender,
			&e.MinAge,
			&e.MaxAge,
		); err != nil {
			log.Error("failed to scan row", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", fn, err)
		}

		rr = append(rr, e)
	}

	return rr, nil
}
