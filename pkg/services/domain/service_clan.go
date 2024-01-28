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
	ServerID  string
	ChannelID string
	ClanTag   string
	Region    models.Region
}

func (s *Service) ClanAdd(
	ctx context.Context,
	request *ClanAddRequest,
) (*models.WGClan, error) {
	instance := &models.BotInstance{
		ServerID:  request.ServerID,
		ChannelID: request.ChannelID,
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

	res := models.WGClan{
		ID:     clan.ClanID,
		Region: request.Region,
		Tag:    clan.Tag,
		Name:   clan.Name,
	}

	err = s.EnsureWGClan(ctx, &res)
	if err != nil {
		return &res, err
	}

	subscription := &models.SubscriptionClan{
		InstanceID: instance.ID,
		ClanID:     clan.ClanID,
		Region:     request.Region,
	}

	return &res, s.EnsureSubscriptionClan(ctx, subscription)
}

func (s *Service) ClanRemove(
	ctx context.Context,
	request *ClanAddRequest,
) (*models.WGClan, error) {
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

	res := models.WGClan{
		ID:     clan.ClanID,
		Region: request.Region,
		Tag:    clan.Tag,
		Name:   clan.Name,
	}

	err = s.EnsureWGClan(ctx, &res)
	if err != nil {
		return &res, err
	}

	subscription := &models.SubscriptionClan{
		InstanceID: instance.ID,
		ClanID:     clan.ClanID,
		Region:     request.Region,
	}

	err = s.EnsureSubscriptionClan(ctx, subscription)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

type ClanListRequest struct {
	ServerID string
}

func (s *Service) ClanList(
	ctx context.Context,
	request *ClanListRequest,
) ([]*models.WGClan, error) {
	instance := &models.BotInstance{
		ServerID:  request.ServerID,
		ChannelID: "",
		Type:      models.STClan,
	}

	instance, err := s.repo.GetInstance(ctx, instance)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	clans, err := s.repo.GetWGClanListByInstance(ctx, instance.ID)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return clans, nil
}
