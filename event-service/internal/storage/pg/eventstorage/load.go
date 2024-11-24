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

func (s *Storage) Load(ctx context.Context, in *domain.EventLoadInfo) (*domain.EventInfo, bool, error) {

	fn := "EventStorage.Load"
	log := s.l.With(sl.Method(fn))

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed acquire connection", sl.Err(err))
		return nil, false, fmt.Errorf("%s: %w", fn, err)
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Error("failed begin transaction", sl.Err(err))
		return nil, false, fmt.Errorf("%s: %w", fn, err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}

		err = tx.Commit(ctx)
		if err != nil {
			log.Error("failed commit transaction", sl.Err(err))
		}
	}()

	ctx = context.WithValue(ctx, "tx", tx)

	oldEvent, err := s.Find(ctx, in.EkpId)
	if err != nil {
		log.Error("failed find event", sl.Err(err))
		return nil, false, fmt.Errorf("%s: %w", fn, err)
	}

	if oldEvent != nil {
		if err := s.unstale(ctx, oldEvent.Id); err != nil {
			log.Error("failed unstale event", sl.Err(err))
			return nil, false, fmt.Errorf("%s: %w", fn, err)
		}

		event, err := s.Find(ctx, oldEvent.Id)
		if err != nil {
			log.Error("failed find event", sl.Err(err))
			return nil, false, fmt.Errorf("%s: %w", fn, err)
		}

		return event, true, nil
	}

	var stId, sstId string
	st, err := s.findSportType(ctx, in.SportType)
	if err != nil {
		log.Error("failed find sport type", sl.Err(err))
		return nil, false, fmt.Errorf("%s: %w", fn, err)
	}

	if st == nil {
		log.Debug("saving sport type", slog.String("slug", in.SportType))

		stid, err := s.saveSportType(ctx, in.SportType)
		if err != nil {
			log.Error("failed save sport type", sl.Err(err))
			return nil, false, fmt.Errorf("%s: %w", fn, err)
		}
		stId = stid

		sstid, err := s.saveSportSubtype(ctx, saveSportSubtype{name: in.SportSubtype, sportTypeId: stId})
		if err != nil {
			log.Error("failed save sport subtype", sl.Err(err))
			return nil, false, fmt.Errorf("%s: %w", fn, err)
		}

		sstId = sstid
	} else {
		stId = st.id

		sst, err := s.findSportSubtypeByName(ctx, stId, in.SportSubtype)
		if err != nil {
			log.Error("failed find sport subtype", sl.Err(err))
			return nil, false, fmt.Errorf("%s: %w", fn, err)
		}

		if sst == nil {
			log.Debug("sport subtype not found", slog.String("slug", in.SportSubtype))

			sstid, err := s.saveSportSubtype(ctx, saveSportSubtype{name: in.SportSubtype, sportTypeId: stId})
			if err != nil {
				log.Error("failed save sport subtype", sl.Err(err))
				return nil, false, fmt.Errorf("%s: %w", fn, err)
			}
			sstId = sstid
		} else {
			sstId = sst.id
		}
	}

	eid, err := s.save(ctx, &saveEvent{
		ekpId:        in.EkpId,
		name:         in.Name,
		description:  in.Description,
		location:     in.Location,
		subtypeId:    sstId,
		participants: in.Participants,
	})
	if err != nil {
		log.Error("failed save event", sl.Err(err))
		return nil, false, fmt.Errorf("%s: %w", fn, err)
	}

	if err := s.saveEventDate(ctx, eid, in.Dates); err != nil {
		log.Error("failed save event dates", sl.Err(err))
		return nil, false, fmt.Errorf("%s: %w", fn, err)
	}

	if err := s.saveParticipantsRequirements(ctx, eid, in.ParticipantRequirements); err != nil {
		log.Error("failed save participants requirements", sl.Err(err))
		return nil, false, fmt.Errorf("%s: %w", fn, err)
	}

	if err := s.saveEventStaleInfo(ctx, eid, false); err != nil {
		log.Error("failed save event stale info", sl.Err(err))
		return nil, false, fmt.Errorf("%s: %w", fn, err)
	}

	event, err := s.Find(ctx, eid)
	if err != nil {
		log.Error("failed find event", sl.Err(err))
		return nil, false, fmt.Errorf("%s: %w", fn, err)
	}

	return event, false, nil
}

func (s *Storage) findSportType(ctx context.Context, slug string) (*sportType, error) {
	fn := "EventStorage.findSportType"
	log := s.l.With(sl.Method(fn))

	conn := ctx.Value("tx").(pgx.Tx)

	qb := sq.
		Select("*").
		From(fmt.Sprintf("%s st", pg.SPORT_TYPES)).
		PlaceholderFormat(sq.Dollar)

	if _, err := uuid.Parse(slug); err != nil {
		qb = qb.Where(sq.Eq{"st.sport_type": slug})
	} else {
		qb = qb.Where(sq.Eq{"st.id": slug})
	}

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

func (s *Storage) findSportSubtypeById(ctx context.Context, id string) (*sportSubtype, error) {
	fn := "EventStorage.findSportSubtypeById"
	log := s.l.With(sl.Method(fn))

	conn := ctx.Value("tx").(pgx.Tx)

	qb := sq.
		Select("*").
		From(fmt.Sprintf("%s sst", pg.SPORT_SUBTYPES)).
		Where(sq.Eq{"sst.id": id}).
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
			log.Debug("sport type not found", slog.String("id", id))
			return nil, nil
		}
		log.Error("failed query row", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return sst, nil
}

func (s *Storage) findSportSubtypeByName(ctx context.Context, sportTypeId, name string) (*sportSubtype, error) {
	fn := "EventStorage.findSportSubtype"
	log := s.l.With(sl.Method(fn))

	conn := ctx.Value("tx").(pgx.Tx)

	qb := sq.
		Select("*").
		From(fmt.Sprintf("%s sst", pg.SPORT_SUBTYPES)).
		Where(sq.And{sq.Eq{"sst.sport_subtype": name}, sq.Eq{"sst.sport_type_id": sportTypeId}}).
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
			log.Debug("sport type not found", slog.String("name", name), slog.String("sportTypeId", sportTypeId))
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

	conn := ctx.Value("tx").(pgx.Tx)

	qb := sq.
		Insert(pg.SPORT_TYPES).
		Columns("sport_type").
		Values(name).
		Suffix("RETURNING id").
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

	conn := ctx.Value("tx").(pgx.Tx)

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

	conn := ctx.Value("tx").(pgx.Tx)

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

	conn := ctx.Value("tx").(pgx.Tx)

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
		var pge *pgconn.PgError
		if errors.As(err, &pge) {
			log.Error("failed exec query", sl.PgError(pge))
		} else {
			log.Error("failed exec query", sl.Err(err))
		}
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) saveParticipantsRequirements(ctx context.Context, eventId string, requirements []domain.ParticipantRequirements) error {
	fn := "EventStorage.saveEventDate"
	log := s.l.With(sl.Method(fn))

	conn := ctx.Value("tx").(pgx.Tx)

	qb := sq.
		Insert(pg.EVENT_PARTICIPANTS_REQUIREMENTS).
		Columns("event_id", "gender", "min_age", "max_age").
		PlaceholderFormat(sq.Dollar)

	for _, r := range requirements {
		if r.MaxAge != nil && r.MinAge != nil {
			qb = qb.Values(eventId, r.Gender, *r.MinAge, *r.MaxAge)
		} else if r.MaxAge != nil {
			qb = qb.Values(eventId, r.Gender, nil, *r.MaxAge)
		} else if r.MinAge != nil {
			qb = qb.Values(eventId, r.Gender, *r.MinAge, nil)
		} else {
			qb = qb.Values(eventId, r.Gender, nil, nil)
		}
	}

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed build sql", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("exectuing sql", slog.String("sql", sql), slog.Any("args", args))

	if _, err := conn.Exec(ctx, sql, args...); err != nil {
		var pge *pgconn.PgError
		if errors.As(err, &pge) {
			log.Error("failed exec query", sl.PgError(pge))
		} else {
			log.Error("failed exec query", sl.Err(err))
		}
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) saveEventStaleInfo(ctx context.Context, eventId string, stale bool) error {
	fn := "EventStorage.saveEventStaleInfo"
	log := s.l.With(sl.Method(fn))

	conn := ctx.Value("tx").(pgx.Tx)

	qb := sq.
		Insert(pg.EVENT_STALE).
		Columns("event_id", "is_stale").
		Values(eventId, stale).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed build sql", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("exectuing sql", slog.String("sql", sql), slog.Any("args", args))

	if _, err := conn.Exec(ctx, sql, args...); err != nil {
		var pge *pgconn.PgError
		if errors.As(err, &pge) {
			log.Error("failed exec query", sl.PgError(pge))
		} else {
			log.Error("failed exec query", sl.Err(err))
		}
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
