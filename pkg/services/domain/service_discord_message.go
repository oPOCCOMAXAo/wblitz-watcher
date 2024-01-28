package domain

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (s *Service) CreateDiscordMessagesForEventClan(
	ctx context.Context,
) error {
	err := s.repo.CreateDiscordMessagesFromEventClan(ctx)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	err = s.repo.UpdateEventClanProcessed(ctx)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}

type DiscordMessageData struct {
	Message *models.DiscordMessage
	User    *wg.AccountInfo
	Clan    *wg.ClanInfo
}

func (s *Service) GetFirstUnsentDiscordMessage(
	ctx context.Context,
) (*DiscordMessageData, error) {
	var (
		res DiscordMessageData
		err error
	)

	res.Message, err = s.repo.GetFirstUnsentDiscordMessage(ctx)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	res.User, err = s.wg.GetAccountByID(ctx, models.WGAccountID{
		ID:     res.Message.AccountID,
		Region: res.Message.Region,
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	res.Clan, err = s.wg.GetClanByID(ctx, models.WGClanID{
		ID:     res.Message.ClanID,
		Region: res.Message.Region,
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return &res, nil
}

func (s *Service) UpdateDiscordMessageProcessed(
	ctx context.Context,
	id int64,
) error {
	err := s.repo.UpdateDiscordMessagesProcessed(ctx, []int64{id})
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
