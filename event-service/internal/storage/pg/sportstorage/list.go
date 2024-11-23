package sportstorage

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/event-service/internal/storage/model"
	"mzhn/event-service/internal/storage/pg"
	"mzhn/event-service/pkg/sl"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) List(ctx context.Context, chout chan<- model.SportTypeWithSubtypes) error {
	fn := "sport-storage.List"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}
	defer conn.Release()

	ctx = context.WithValue(ctx, "tx", conn)

	qb := sq.
		Select("st.id, st.sport_type").
		From(fmt.Sprintf("%s st", pg.SPORT_TYPES)).
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
		var e model.SportTypeWithSubtypes
		log.Debug("scan row")
		if err := rows.Scan(
			&e.SportType.Id,
			&e.SportType.Name,
		); err != nil {
			log.Error("failed to scan row", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}

		log.Debug("listing requirements for event")
		rr, err := s.listSubtypes(ctx, e.Id)
		if err != nil {
			log.Error("failed to list requirements for event", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}

		e.Subtypes = rr

		log.Debug("sending event to channel")
		chout <- e
	}

	return nil
}

func (s *Storage) listSubtypes(ctx context.Context, sportId string) ([]model.Subtype, error) {

	fn := "sport-storage.listSubtypes"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer conn.Release()

	qb := sq.
		Select("sst.id, sst.sport_subtype").
		From(fmt.Sprintf("%s sst", pg.SPORT_SUBTYPES)).
		Where(sq.Eq{"sst.sport_type_id": sportId}).
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

	rr := make([]model.Subtype, 0, 16)
	for rows.Next() {
		var e model.Subtype
		if err := rows.Scan(
			&e.Id,
			&e.Name,
		); err != nil {
			log.Error("failed to scan row", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", fn, err)
		}

		rr = append(rr, e)
	}

	return rr, nil
}
