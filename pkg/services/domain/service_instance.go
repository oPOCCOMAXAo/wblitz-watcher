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

	if record.ChannelID == "" {
		record.ChannelID = instance.ChannelID

		err = s.repo.UpdateInstance(ctx, record)
		if err != nil {
			//nolint:wrapcheck
			return err
		}
	}

	instance.ID = record.ID

	return nil
}

func (s *Service) EnsureInstancesForAllTypes(
	ctx context.Context,
	params *FastFixParams,
) error {
	for _, typ := range models.SubscriptionTypes {
		err := s.EnsureInstance(ctx, &models.BotInstance{
			ServerID:  params.ServerID,
			ChannelID: params.ChannelID,
			Type:      typ,
		})
		if err != nil {
			return err
		}
	}

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

func (s *Service) CleanupDeletedInstances(
	ctx context.Context,
) error {
	err := s.repo.DeleteSubscriptionClansForDeletedInstances(ctx)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	err = s.repo.DeleteDiscordMessagesForDeletedInstances(ctx)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	err = s.repo.HardDeleteCleanedInstances(ctx)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
