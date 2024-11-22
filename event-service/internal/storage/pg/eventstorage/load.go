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
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type saveEvent struct {
	ekpId        string
	name         string
	description  string
	location     string
	subtypeId    string
	participants int
}

type sportType struct {
	id   string
	name string
}

type saveSportSubtype struct {
	name        string
	sportTypeId string
}

type sportSubtype struct {
	id          string
	name        string
	sportTypeId string
}

func (s *Storage) Load(ctx context.Context, in *domain.EventLoadInfo) (string, error) {

	fn := "EventStorage.Load"
	log := s.l.With(sl.Method(fn))

	var stId, sstId string
	st, err := s.findSportType(ctx, in.SportType)
	if err != nil {
		log.Error("failed find sport type", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	if st == nil {
		log.Debug("sport type not found", slog.String("slug", in.SportType))
		id, err := s.saveSportType(ctx, in.SportType)
		if err != nil {
			log.Error("failed save sport type", sl.Err(err))
			return "", fmt.Errorf("%s: %w", fn, err)
		}
		stId = id
	} else {
		stId = st.id
	}

	sst, err := s.findSportSubtype(ctx, in.SportSubtype)
	if err != nil {
		log.Error("failed find sport subtype", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	if sst == nil {
		log.Debug("sport subtype not found", slog.String("slug", in.SportSubtype))
		id, err := s.saveSportSubtype(ctx, saveSportSubtype{name: in.SportSubtype, sportTypeId: stId})
		if err != nil {
			log.Error("failed save sport subtype", sl.Err(err))
			return "", fmt.Errorf("%s: %w", fn, err)
		}
		sstId = id
	}

	eId, err := s.save(ctx, &saveEvent{
		ekpId:        in.EkpId,
		name:         in.Name,
		description:  in.Description,
		location:     in.Location,
		subtypeId:    sstId,
		participants: in.Participants,
	})
	if err != nil {
		log.Error("failed save event", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	if err := s.saveEventDate(ctx, eId, in.Dates); err != nil {
		log.Error("failed save event dates", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	return eId, nil
}

func (s *Storage) findSportType(ctx context.Context, slug string) (*sportType, error) {
	fn := "EventStorage.findSportType"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed acquire connection", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	qb := sq.
		Select("*").
		From(fmt.Sprintf("%s st", pg.SPORT_TYPES)).
		Where(sq.Or{sq.Eq{"st.id": slug}, sq.Eq{"st.sport_type": slug}}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed build sql", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("exectuing sql", slog.String("sql", sql), slog.Any("args", args))

	st := new(sportType)
	if err := conn.QueryRow(ctx, sql, args...).Scan(
		&st.id,
		&st.name,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Debug("sport type not found", slog.String("slug", slug))
			return nil, nil
		}
		log.Error("failed query row", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return st, nil
}

func (s *Storage) findSportSubtype(ctx context.Context, slug string) (*sportSubtype, error) {
	fn := "EventStorage.findSportSubtype"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed acquire connection", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	qb := sq.
		Select("*").
		From(fmt.Sprintf("%s sst", pg.SPORT_SUBTYPES)).
		Where(sq.Or{sq.Eq{"sst.id": slug}, sq.Eq{"sst.sport_subtype": slug}}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed build sql", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("exectuing sql", slog.String("sql", sql), slog.Any("args", args))

	sst := new(sportSubtype)
	if err := conn.QueryRow(ctx, sql, args...).Scan(
		&sst.id,
		&sst.name,
		&sst.sportTypeId,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Debug("sport type not found", slog.String("slug", slug))
			return nil, nil
		}
		log.Error("failed query row", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return sst, nil
}

func (s *Storage) saveSportType(ctx context.Context, name string) (string, error) {
	fn := "EventStorage.saveSportType"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed acquire connection", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	qb := sq.
		Insert(pg.SPORT_TYPES).
		Columns("sport_type").
		Values(name).
		Suffix(fmt.Sprintf("RETURNING id")).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed build sql", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("exectuing sql", slog.String("sql", sql), slog.Any("args", args))

	var id string
	if err := conn.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		log.Error("failed query row", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	return id, nil
}

func (s *Storage) saveSportSubtype(ctx context.Context, in saveSportSubtype) (string, error) {
	fn := "EventStorage.saveSportSubtype"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed acquire connection", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	qb := sq.
		Insert(pg.SPORT_SUBTYPES).
		Columns("sport_subtype", "sport_type_id").
		Values(in.name, in.sportTypeId).
		Suffix(fmt.Sprintf("RETURNING id")).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed build sql", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("exectuing sql", slog.String("sql", sql), slog.Any("args", args))

	var id string
	if err := conn.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		log.Error("failed query row", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	return id, nil
}

func (s *Storage) save(ctx context.Context, in *saveEvent) (string, error) {

	fn := "EventStorage.findSportSubtype"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed acquire connection", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	qb := sq.
		Insert(pg.EVENTS).
		Columns("ekp_id", "sport_subtype_id", "name", "description", "location", "participants").
		Values(in.ekpId, in.subtypeId, in.name, in.description, in.location, in.participants).
		Suffix(fmt.Sprintf("RETURNING id")).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed build sql", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("exectuing sql", slog.String("sql", sql), slog.Any("args", args))

	var id string
	if err := conn.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		log.Error("failed query row", sl.Err(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	return id, nil
}

func (s *Storage) saveEventDate(ctx context.Context, eventId string, dates domain.DateRange) error {
	fn := "EventStorage.saveEventDate"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed acquire connection", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	qb := sq.
		Insert(pg.EVENT_DATES).
		Columns("event_id", "date_from", "date_to").
		Values(eventId, dates.From, dates.To).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed build sql", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("exectuing sql", slog.String("sql", sql), slog.Any("args", args))

	if _, err := conn.Exec(ctx, sql, args...); err != nil {
		var pge pgconn.PgError
		if errors.As(err, &pge) {
			log.Error("failed exec query", sl.PgError(pge))
		} else {
			log.Error("failed exec query", sl.Err(err))
		}
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
