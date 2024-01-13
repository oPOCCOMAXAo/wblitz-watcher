package domain

import (
	"context"
	"errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (s *Service) EnsureSubscriptionClan(
	ctx context.Context,
	value *models.SubscriptionClan,
) error {
	record, err := s.repo.GetSubscriptionClan(ctx, value)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		//nolint:wrapcheck
		return err
	}

	if record == nil {
		err = s.repo.CreateSubscriptionClan(ctx, value)
		if err != nil {
			//nolint:wrapcheck
			return err
		}

		return nil
	}

	value.ID = record.ID

	return nil
}

type ClanAddRequest struct {
	ServerID string
	ClanTag  string
	Region   models.Region
}

func (s *Service) ClanAdd(
	ctx context.Context,
	request *ClanAddRequest,
) (*wg.ClanListEntry, error) {
	instance := &models.BotInstance{
		ServerID:  request.ServerID,
		ChannelID: "",
		Type:      models.STClan,
	}

	err := s.EnsureInstance(ctx, instance)
	if err != nil {
		return nil, err
	}

	clan, err := s.wg.FindClanByTag(ctx, wg.ClansListRequest{
		Search: request.ClanTag,
		Region: request.Region,
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	subscription := &models.SubscriptionClan{
		InstanceID: instance.ID,
		ClanID:     clan.ClanID,
		Region:     request.Region,
	}

	return clan, s.EnsureSubscriptionClan(ctx, subscription)
}

func (s *Service) ClanRemove(
	ctx context.Context,
	request *ClanAddRequest,
) (*wg.ClanListEntry, error) {
	instance := &models.BotInstance{
		ServerID:  request.ServerID,
		ChannelID: "",
		Type:      models.STClan,
	}

	err := s.EnsureInstance(ctx, instance)
	if err != nil {
		return nil, err
	}

	clan, err := s.wg.FindClanByTag(ctx, wg.ClansListRequest{
		Search: request.ClanTag,
		Region: request.Region,
	})
	if err != nil {
		//nolint:wrapcheck
		return clan, err
	}

	subscription := &models.SubscriptionClan{
		InstanceID: instance.ID,
		ClanID:     clan.ClanID,
		Region:     request.Region,
	}

	return clan, s.repo.DeleteSubscriptionClan(ctx, subscription)
}
