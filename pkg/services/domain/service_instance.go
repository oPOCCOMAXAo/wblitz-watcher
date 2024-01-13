package domain

import (
	"context"
	"errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (s *Service) EnsureInstance(
	ctx context.Context,
	instance *models.BotInstance,
) error {
	record, err := s.repo.GetInstance(ctx, instance)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		//nolint:wrapcheck
		return err
	}

	if record == nil {
		err = s.repo.CreateInstance(ctx, instance)
		if err != nil {
			//nolint:wrapcheck
			return err
		}

		return nil
	}

	instance.ID = record.ID

	return nil
}

type ChannelBindRequest struct {
	ServerID  string
	ChannelID string
	Type      models.SubscriptionType
}

func (s *Service) ChannelBind(
	ctx context.Context,
	request *ChannelBindRequest,
) error {
	instance := &models.BotInstance{
		ServerID:  request.ServerID,
		ChannelID: request.ChannelID,
		Type:      request.Type,
	}

	err := s.EnsureInstance(ctx, instance)
	if err != nil {
		return err
	}

	err = s.repo.UpdateInstance(ctx, instance)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
