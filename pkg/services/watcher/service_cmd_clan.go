package watcher

import (
	"context"
	"fmt"

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
		ServerID:  event.GuildID,
		ChannelID: event.ChannelID,
		ClanTag:   data.String("clan"),
		Region:    models.Region(data.String("server")),
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

		return &res, nil
	}

	used := int64(len(clans.ClansEnabled))

	res.Content = fmt.Sprintf("Clan list. Current limit: `%d/%d`",
		used,
		clans.Limit,
	)

	if used >= clans.Limit {
		res.Content += "\n\nLimit reached. Extra clans will be disabled."
	}

	if len(clans.ClansEnabled) > 0 {
		res.Embeds = append(res.Embeds, s.embedClanList(clans.ClansEnabled, false)...)
	}

	if len(clans.ClansDisabled) > 0 {
		res.Embeds = append(res.Embeds, s.embedClanList(clans.ClansDisabled, true)...)
	}

	return &res, nil
}
