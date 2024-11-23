package verificationstorage

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/notification-service/internal/storage/model"
	"mzhn/notification-service/internal/storage/pg"
	"mzhn/notification-service/pkg/sl"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) Save(ctx context.Context, in *model.NewVerification) (*model.Verification, error) {
	fn := "integrationstorage.Save"
	log := s.l.With(sl.Method(fn))

	i, err := s.find(ctx, in.UserId)
	if err != nil {
		log.Error("failed to find user", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	if i != nil {
		v, err := s.update(ctx, in)
		if err != nil {
			log.Error("failed to update user", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", fn, err)

		}
		return v, nil
	}

	v, err := s.save(ctx, in)
	if err != nil {
		log.Error("failed to save user", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return v, nil
}

func (s *Storage) save(ctx context.Context, in *model.NewVerification) (*model.Verification, error) {
	fn := "verificationservice.save"
	log := s.l.With(sl.Module(fn))

	c, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)

	}
	defer c.Release()

	qb := sq.
		Insert(pg.VERIFICATIONS).
		Columns("user_id", "token").
		Values(in.UserId, in.Token).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("executing query", slog.String("sql", sql), slog.Any("args", args))

	v := new(model.Verification)
	if err := c.QueryRow(ctx, sql, args...).Scan(&v.UserId, &v.Token, &v.CreatedAt); err != nil {
		log.Error("failed to execute query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return v, nil
}

func (s *Storage) update(ctx context.Context, in *model.NewVerification) (*model.Verification, error) {
	fn := "verificationstorage.update"
	log := s.l.With(sl.Module(fn))

	c, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer c.Release()

	qb := sq.
		Update(pg.INTEGRATIONS).
		Where(sq.Eq{"user_id": in.UserId}).
		Set("token", in.Token).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("executing query", slog.String("sql", sql), slog.Any("args", args))

	v := new(model.Verification)
	if err := c.QueryRow(ctx, sql, args...).Scan(&v.UserId, &v.Token, &v.CreatedAt); err != nil {
		log.Error("failed to execute query", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return v, nil
}
