package cron

import (
	"context"
	"log/slog"
	"mzhn/notification-cron/internal/config"
	"mzhn/notification-cron/internal/domain"
	"mzhn/notification-cron/pkg/sl"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type CronHandler struct {
	l   *slog.Logger
	cfg *config.Config
	cs  domain.CronService

	tasks []gocron.Task
}

func New(cfg *config.Config, cs domain.CronService) *CronHandler {
	return &CronHandler{
		l:     slog.With(sl.Module("cron handler")),
		cfg:   cfg,
		cs:    cs,
		tasks: make([]gocron.Task, 0),
	}
}

func (h *CronHandler) Run(ctx context.Context) error {
	h.setup(ctx)
	s, err := gocron.NewScheduler()
	if err != nil {
		h.l.Error("cannot start cron scheduler")
		return err
	}

	for _, task := range h.tasks {
		s.NewJob(gocron.DurationJob(time.Duration(h.cfg.Cron.Minutes)*time.Minute), task)
	}

	s.Start()

	<-ctx.Done()
	h.l.Info("shutting down cron")
	s.Shutdown()
	return nil
}

func (h *CronHandler) setup(ctx context.Context) {
	h.tasks = append(h.tasks, gocron.NewTask(h.cs.NotifyUsers, ctx, 30))
	h.tasks = append(h.tasks, gocron.NewTask(h.cs.NotifyUsers, ctx, 7))
	h.tasks = append(h.tasks, gocron.NewTask(h.cs.NotifyUsers, ctx, 3))

	for _, daysLeft := range []int{30, 7, 3} {
		go h.cs.NotifyUsers(ctx, uint32(daysLeft))
	}
}
