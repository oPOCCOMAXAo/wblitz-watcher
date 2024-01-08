package watcher

import (
	"context"
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
)

func (s *Service) cmdClanAdd(
	event *discordgo.InteractionCreate,
	data *discord.CommandData,
) (*discord.Response, error) {
	req := wg.ClansListRequest{
		Search: data.String("clan"),
		Region: models.Region(data.String("server")),
	}

	clan, err := s.wg.FindClanByTag(context.Background(), req)

	if clan == nil {
		res := &discord.Response{
			Content: "Clan not found",
		}

		if err != nil {
			res.Embeds = append(res.Embeds, s.embedError(err))
		}

		//nolint:wrapcheck
		return res, err
	}

	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return &discord.Response{
		Content: "Debug",
		Embeds: []*discordgo.MessageEmbed{
			{
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "Tag",
						Value: clan.Tag,
					},
					{
						Name:  "Name",
						Value: clan.Name,
					},
					{
						Name:  "Members",
						Value: strconv.Itoa(clan.MembersCount),
					},
					{
						Name:  "DS Server",
						Value: event.GuildID,
					},
				},
			},
		},
	}, nil
}
