package domain

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (s *Service) UpdateAllSubscriptionClanActivation(
	ctx context.Context,
) error {
	instances, err := s.repo.GetInstancesByType(ctx, models.STClan)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	for _, instance := range instances {
		err = s.ActivateSubscriptionClansForInstance(ctx, instance)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) ActivateSubscriptionClansForInstance(
	ctx context.Context,
	instance *models.BotInstance,
) error {
	limit, err := s.GetSubscriptionClanLimitForServer(ctx, instance.ServerID)
	if err != nil {
		return err
	}

	clans, err := s.repo.GetSubscriptionClanListByInstance(ctx, instance.ID)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	disableID := make([]int64, 0, len(clans))
	enableID := make([]int64, 0, len(clans))

	for i, clan := range clans {
		if int64(i) < limit {
			if clan.IsDisabled {
				enableID = append(enableID, clan.ID)
			}
		} else {
			if !clan.IsDisabled {
				disableID = append(disableID, clan.ID)
			}
		}
	}

	err = s.repo.UpdateIsDisabledSubscriptionClanByID(ctx, true, disableID)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	err = s.repo.UpdateIsDisabledSubscriptionClanByID(ctx, false, enableID)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
