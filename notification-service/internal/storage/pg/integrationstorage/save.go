package integrationstorage

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/internal/storage/pg"
	"mzhn/notification-service/pkg/sl"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) Create(ctx context.Context, userId string) error {
	return s.save(ctx, &domain.SetIntegrations{UserId: userId})
}

func (s *Storage) Save(ctx context.Context, in *domain.SetIntegrations) error {
	fn := "integrationstorage.Save"
	log := s.l.With(sl.Method(fn))

	i, err := s.find(ctx, in.UserId)
	if err != nil {
		log.Error("failed to find user", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	if i != nil {
		if err := s.update(ctx, in); err != nil {
			log.Error("failed to update user", sl.Err(err))
			return fmt.Errorf("%s: %w", fn, err)
		}
		return nil
	}

	if err := s.save(ctx, in); err != nil {
		log.Error("failed to save user", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) save(ctx context.Context, in *domain.SetIntegrations) error {
	fn := "integrationstorage.save"
	log := s.l.With(sl.Module(fn))

	c, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}
	defer c.Release()

	values := make([]interface{}, 3)
	values[0] = in.UserId
	values[1] = in.TelegramUsername
	values[2] = true
	if in.WannaMail != nil {
		values[2] = *in.WannaMail
	}

	qb := sq.
		Insert(pg.INTEGRATIONS).
		Columns("user_id", "telegram_username", "wanna_mail").
		Values(values...).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build query", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("executing query", slog.String("sql", sql), slog.Any("args", args))

	if _, err := c.Exec(ctx, sql, args...); err != nil {
		log.Error("failed to execute query", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) update(ctx context.Context, in *domain.SetIntegrations) error {
	fn := "integrationstorage.update"
	log := s.l.With(sl.Module(fn))

	c, err := s.pool.Acquire(ctx)
	if err != nil {
		log.Error("failed to acquire connection", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}
	defer c.Release()

	qb := sq.
		Update(pg.INTEGRATIONS).
		Where(sq.Eq{"user_id": in.UserId}).
		PlaceholderFormat(sq.Dollar)

	if in.TelegramUsername != nil {
		qb = qb.Set("telegram_username", *in.TelegramUsername)
	}

	if in.WannaMail != nil {
		qb = qb.Set("wanna_mail", *in.WannaMail)
	}

	sql, args, err := qb.ToSql()
	if err != nil {
		log.Error("failed to build query", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	log.Debug("executing query", slog.String("sql", sql), slog.Any("args", args))

	if _, err := c.Exec(ctx, sql, args...); err != nil {
		log.Error("failed to execute query", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
