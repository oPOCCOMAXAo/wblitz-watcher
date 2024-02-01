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

	err = s.EnsureSubscriptionClan(ctx, subscription)
	if err != nil {
		return &res, err
	}

	err = s.ActivateSubscriptionClansForInstance(ctx, instance)
	if err != nil {
		return &res, err
	}

	return &res, nil
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

	err = s.ActivateSubscriptionClansForInstance(ctx, instance)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

type ClanListRequest struct {
	ServerID string
}

type ClanListResponse struct {
	ClansEnabled  []*models.WGClan
	ClansDisabled []*models.WGClan
	Limit         int64
}

func (s *Service) ClanList(
	ctx context.Context,
	request *ClanListRequest,
) (*ClanListResponse, error) {
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

	subs, err := s.repo.GetSubscriptionClanListByInstance(ctx, instance.ID)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	disabledByID := make(map[int64]bool, len(subs))
	for _, sub := range subs {
		disabledByID[sub.ClanID] = sub.IsDisabled
	}

	clans, err := s.repo.GetWGClanListByInstance(ctx, instance.ID)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	var res ClanListResponse

	res.Limit, err = s.GetSubscriptionClanLimitForServer(ctx, instance.ServerID)
	if err != nil {
		return nil, err
	}

	for _, clan := range clans {
		if disabledByID[clan.ID] {
			res.ClansDisabled = append(res.ClansDisabled, clan)
		} else {
			res.ClansEnabled = append(res.ClansEnabled, clan)
		}
	}

	return &res, nil
}
