package watcher

import (
	"context"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	du "github.com/opoccomaxao/wblitz-watcher/pkg/utils/discordutils"
)

func (s *Service) cmdClan(
	event *discordgo.InteractionCreate,
) (*discord.Response, error) {
	data := event.ApplicationCommandData()

	log.Printf("%s %s\n", data.Name, data.Options[0].Name)

	switch data.Options[0].Name {
	case "add":
		return s.cmdClanAdd(event)
	case "remove":
		// TODO
	case "list":
		// TODO
	}

	return nil, models.ErrNotFound
}

func (s *Service) cmdClanAdd(
	event *discordgo.InteractionCreate,
) (*discord.Response, error) {
	err := s.discord.VerifyAccess(event)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	var req wg.ClansListRequest

	err = du.ParseOptions(event.ApplicationCommandData().Options[0].Options, du.DecodersMap{
		"clan":   du.DecoderString(&req.Search),
		"server": du.DecoderString(&req.Region),
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
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
