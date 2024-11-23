package integrationservice

import (
	"context"
	"fmt"
	"log/slog"
	"mzhn/notification-service/internal/config"
	"mzhn/notification-service/internal/domain"
	"mzhn/notification-service/internal/storage/model"
	"mzhn/notification-service/pkg/sl"
	"time"

	"github.com/google/uuid"
)

type IntegrationsSaver interface {
	Create(ctx context.Context, userId string) error
	Save(context.Context, *domain.SetIntegrations) error
}
type IntegrationsProvider interface {
	Find(ctx context.Context, userId string) (*domain.Integrations, error)
}

type VerificationSaver interface {
	Save(context.Context, *model.NewVerification) (*model.Verification, error)
}
type VerificationProvider interface {
	Find(ctx context.Context, userId string) (*model.Verification, error)
}

type Service struct {
	l   *slog.Logger
	cfg *config.Config
	is  IntegrationsSaver
	ip  IntegrationsProvider
	vs  VerificationSaver
	vp  VerificationProvider
}

func New(
	cfg *config.Config,
	is IntegrationsSaver,
	ip IntegrationsProvider,
	vs VerificationSaver,
	vp VerificationProvider,
) *Service {
	return &Service{
		l:   slog.With(sl.Module("integration-service")),
		cfg: cfg,
		is:  is,
		ip:  ip,
		vs:  vs,
		vp:  vp,
	}
}

func (s *Service) Find(ctx context.Context, userId string) (*domain.Integrations, error) {
	fn := "integrationservice.Find"
	log := s.l.With(sl.Module(fn))

	i, err := s.ip.Find(ctx, userId)
	if err != nil {
		log.Error("failed to find integrations", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	if i == nil {
		if err := s.is.Create(ctx, userId); err != nil {
			return nil, err
		}

		newI, err := s.ip.Find(ctx, userId)
		if err != nil {
			return nil, err
		}

		i = newI
	}

	return i, nil
}

func (s *Service) Save(ctx context.Context, req *domain.SetIntegrations) error {
	fn := "integrationservice.Save"
	log := s.l.With(sl.Module(fn))

	if err := s.is.Save(ctx, req); err != nil {
		log.Error("failed to save integrations", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Service) generateVerificationCode() string {
	return uuid.NewString()
}

func (s *Service) GetVerificationCode(ctx context.Context, in *domain.VerificationCodeRequest) (*domain.Verification, error) {
	fn := "integrationservice.GetVerificationCode"
	log := s.l.With(sl.Module(fn))

	token := s.generateVerificationCode()

	v, err := s.vs.Save(ctx, &model.NewVerification{
		UserId: in.UserId,
		Token:  token,
	})
	if err != nil {
		log.Error("failed to save verification", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &domain.Verification{
		UserId:   v.UserId,
		Token:    token,
		ExpireAt: s.calculateExpireAt(v.CreatedAt),
	}, nil
}

func (s *Service) LinkTelegram(ctx context.Context, in *domain.LinkTelegramRequest) error {

	fn := "integrationservice.LinkTelegram"
	log := s.l.With(sl.Method(fn))

	if err := s.verify(ctx, &domain.VerificationRequest{
		UserId: in.UserId,
		Token:  in.Token,
	}); err != nil {
		log.Error("failed to verify token", sl.Err(err), slog.Any("request", in))
		return fmt.Errorf("%s: %w", fn, err)
	}

	if err := s.is.Save(ctx, &domain.SetIntegrations{
		UserId:           in.UserId,
		TelegramUsername: &in.TelegramUsername,
	}); err != nil {
		log.Error("failed to save integrations", sl.Err(err), slog.Any("request", in))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Service) calculateExpireAt(createdAt time.Time) time.Time {
	return createdAt.Add(time.Minute * time.Duration(s.cfg.Verificator.TTL))
}

func (s *Service) verify(ctx context.Context, in *domain.VerificationRequest) error {
	fn := "integrationservice.verify"
	log := s.l.With(sl.Module(fn))

	v, err := s.vp.Find(ctx, in.UserId)
	if err != nil {
		log.Error("failed to find verification", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	if v == nil {
		return fmt.Errorf("%s: %w", fn, domain.ErrVerificationNotFound)
	}

	if v.Token != in.Token {
		return fmt.Errorf("%s: %w", fn, domain.ErrVerificationInvalidToken)
	}

	expireAt := s.calculateExpireAt(v.CreatedAt)

	if time.Now().After(expireAt) {
		return fmt.Errorf("%s: %w", fn, domain.ErrVerificationExpired)
	}

	return nil
}
