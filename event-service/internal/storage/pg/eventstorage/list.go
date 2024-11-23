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

func (s *Storage) applyListPagination(qb sq.SelectBuilder, filters model.Pagination) sq.SelectBuilder {

	qb = qb.Limit(filters.Limit)
	qb = qb.Offset(filters.Offset)

	return qb
}
func (s *Storage) applyListFilters(qb sq.SelectBuilder, filters model.EventsFilters) sq.SelectBuilder {

	if filters.StartDate != nil {
		qb = qb.Where(sq.GtOrEq{"ed.date_from": *filters.StartDate})
	}

	if filters.EndDate != nil {
		qb = qb.Where(sq.LtOrEq{"ed.date_from": *filters.EndDate})
	}

	if len(filters.SportTypesId) != 0 {
		clause := sq.Or{}
		for _, id := range filters.SportTypesId {
			clause = append(clause, sq.Eq{"st.id": id})
		}
		qb = qb.Where(clause)
	}

	if len(filters.SportSubtypesId) != 0 {
		clause := sq.Or{}
		for _, id := range filters.SportSubtypesId {
			clause = append(clause, sq.Eq{"sst.id": id})
		}
		qb = qb.Where(clause)
	}

	if filters.Location != nil {
		qb = qb.Where(sq.ILike{"e.location": fmt.Sprintf("%%%s%%", *filters.Location)})
	}

	if filters.Name != nil {
		qb = qb.Where(sq.ILike{"e.name": fmt.Sprintf("%%%s%%", *filters.Name)})
	}

	if filters.MinParticipants != nil {
		qb = qb.Where(sq.GtOrEq{"e.participants": *filters.MinParticipants})
	}

	if filters.MaxParticipants != nil {
		qb = qb.Where(sq.LtOrEq{"e.participants": *filters.MaxParticipants})
	}

	return qb
}

func (s *Storage) List(ctx context.Context, chEvents chan<- domain.EventInfo, filters ...model.EventsFilters) error {

	fn := "EventStorage.List"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}
	defer conn.Release()

	ctx = context.WithValue(ctx, "tx", conn)

	qb := sq.
		Select("e.id, e.ekp_id, st.id, st.sport_type, sst.id, sst.sport_subtype, e.name, e.description, e.location, e.participants, ed.date_from, ed.date_to").
		From(fmt.Sprintf("%s e", pg.EVENTS)).
		InnerJoin(fmt.Sprintf("%s sst on e.sport_subtype_id = sst.id", pg.SPORT_SUBTYPES)).
		InnerJoin(fmt.Sprintf("%s st on sst.sport_type_id = st.id", pg.SPORT_TYPES)).
		InnerJoin(fmt.Sprintf("%s ed on ed.event_id = e.id", pg.EVENT_DATES)).
		OrderBy("ed.date_from ASC").
		PlaceholderFormat(sq.Dollar)

	if len(filters) != 0 {
		f := filters[0]
		qb = s.applyListFilters(qb, f)
		if f.Pagination != nil {
			qb = s.applyListPagination(qb, *f.Pagination)
		}
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
		log.Debug("scan row")
		if err := rows.Scan(
			&e.Id,
			&e.EkpId,
			&e.SportSubtype.Parent.Id,
			&e.SportSubtype.Parent.Name,
			&e.SportSubtype.Id,
			&e.SportSubtype.Name,
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

		log.Debug("listing requirements for event")
		rr, err := s.listRequirementsFor(ctx, e.Id, filters...)
		if err != nil {
			log.Error("failed to list requirements for event", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}

		e.ParticipantRequirements = rr

		log.Debug("sending event to channel")
		chEvents <- e
	}

	return nil
}

func (s *Storage) listRequirementsFor(ctx context.Context, eventId string, filters ...model.EventsFilters) ([]domain.ParticipantRequirements, error) {

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

	if len(filters) != 0 {
		f := filters[0]
		if f.Sex != nil {
			qb = qb.Where(sq.Eq{"epr.gender": *f.Sex})
		}

		if f.MaxAge != nil {
			qb = qb.Where(sq.LtOrEq{"epr.max_age": *f.MaxAge})
		}

		if f.MinAge != nil {
			qb = qb.Where(sq.GtOrEq{"epr.min_age": *f.MinAge})
		}
	}

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
