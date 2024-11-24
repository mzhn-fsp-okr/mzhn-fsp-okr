package sportstorage

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

func (s *Storage) Find(ctx context.Context, id string) (*model.SportTypeWithSubtypes, error) {
	fn := "sport-storage.Find"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer conn.Release()

	ctx = context.WithValue(ctx, "tx", conn)

	qb := sq.
		Select("st.id, st.sport_type").
		From(fmt.Sprintf("%s st", pg.SPORT_TYPES)).
		Where(sq.Eq{"st.id": id}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("executing query", slog.String("sql", sql), slog.Any("args", args))

	var e model.SportTypeWithSubtypes
	if err := conn.QueryRow(ctx, sql, args...).
		Scan(
			&e.SportType.Id,
			&e.SportType.Name,
		); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		log.Error("failed to execute query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("listing requirements for event")
	rr, err := s.listSubtypes(ctx, e.Id)
	if err != nil {
		log.Error("failed to list requirements for event", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	e.Subtypes = rr

	return &e, nil
}
