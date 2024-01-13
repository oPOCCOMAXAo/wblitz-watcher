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
	event *discordgo.InteractionCreate,
	data *discord.CommandData,
) (*discord.Response, error) {
	request := domain.ClanAddRequest{
		ServerID: event.GuildID,
		ClanTag:  data.String("clan"),
		Region:   models.Region(data.String("server")),
	}

	clan, err := s.domain.ClanAdd(context.TODO(), &request)

	var res discord.Response

	if err != nil {
		res.Content = "Error"

		telemetry.RecordErrorBackground(err)
	} else {
		res.Content = "Clan added"
	}

	if clan != nil {
		res.Embeds = append(res.Embeds, s.embedClan(clan))
	}

	return &res, nil
}

func (s *Service) cmdClanRemove(
	event *discordgo.InteractionCreate,
	data *discord.CommandData,
) (*discord.Response, error) {
	request := domain.ClanAddRequest{
		ServerID: event.GuildID,
		ClanTag:  data.String("clan"),
		Region:   models.Region(data.String("server")),
	}

	clan, err := s.domain.ClanRemove(context.TODO(), &request)

	var res discord.Response

	if err != nil {
		res.Content = "Error"

		telemetry.RecordErrorBackground(err)
	} else {
		res.Content = "Clan removed"
	}

	if clan != nil {
		res.Embeds = append(res.Embeds, s.embedClan(clan))
	}

	return &res, nil
}
