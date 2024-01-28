package watcher

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/domain"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

func (s *Service) cmdClanAdd(
	ctx context.Context,
	event *discordgo.InteractionCreate,
	data *discord.CommandData,
) (*discord.Response, error) {
	request := domain.ClanAddRequest{
		ServerID: event.GuildID,
		ClanTag:  data.String("clan"),
		Region:   models.Region(data.String("server")),
	}

	clan, err := s.domain.ClanAdd(ctx, &request)

	var res discord.Response

	if err != nil {
		res.Content = MessageError

		telemetry.RecordError(ctx, err)
	} else {
		res.Content = "Clan added"
	}

	if clan != nil {
		res.Embeds = append(res.Embeds, s.embedClan(clan))
	}

	return &res, nil
}

func (s *Service) cmdClanRemove(
	ctx context.Context,
	event *discordgo.InteractionCreate,
	data *discord.CommandData,
) (*discord.Response, error) {
	request := domain.ClanAddRequest{
		ServerID: event.GuildID,
		ClanTag:  data.String("clan"),
		Region:   models.Region(data.String("server")),
	}

	clan, err := s.domain.ClanRemove(ctx, &request)

	var res discord.Response

	if err != nil {
		res.Content = MessageError

		telemetry.RecordError(ctx, err)
	} else {
		res.Content = "Clan removed"
	}

	if clan != nil {
		res.Embeds = append(res.Embeds, s.embedClan(clan))
	}

	return &res, nil
}

func (s *Service) cmdClanList(
	ctx context.Context,
	event *discordgo.InteractionCreate,
	_ *discord.CommandData,
) (*discord.Response, error) {
	request := domain.ClanListRequest{
		ServerID: event.GuildID,
	}

	clans, err := s.domain.ClanList(ctx, &request)

	var res discord.Response

	if err != nil {
		res.Content = MessageError

		telemetry.RecordError(ctx, err)
	} else {
		res.Content = "Clans"
	}

	if clans != nil {
		res.Embeds = append(res.Embeds, s.embedClanList(clans)...)
	}

	return &res, nil
}
