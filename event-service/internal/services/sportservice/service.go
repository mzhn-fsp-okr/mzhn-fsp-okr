package sportservice

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/event-service/internal/config"
	"mzhn/event-service/internal/storage/model"
	"mzhn/event-service/pkg/sl"
)

type SportProvider interface {
	Find(ctx context.Context, id string) (*model.SportTypeWithSubtypes, error)
	List(ctx context.Context, chout chan<- model.SportTypeWithSubtypes) error
}

type Service struct {
	l   *slog.Logger
	cfg *config.Config
	sp  SportProvider
}

func New(cfg *config.Config, sp SportProvider) *Service {
	return &Service{
		l:   slog.With(sl.Module("sport-service")),
		cfg: cfg,
		sp:  sp,
	}
}

func (s *Service) List(ctx context.Context, chout chan<- model.SportTypeWithSubtypes) error {

	fn := "SportService.List"
	log := s.l.With(sl.Method(fn))

	err := s.sp.List(ctx, chout)
	if err != nil {
		log.Error("failed to list sports", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Service) Find(ctx context.Context, id string) (*model.SportTypeWithSubtypes, error) {

	fn := "SportService.List"
	log := s.l.With(sl.Method(fn))

	sp, err := s.sp.Find(ctx, id)
	if err != nil {
		log.Error("failed to list sports", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return sp, nil
}
