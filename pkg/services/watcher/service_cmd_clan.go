package watcher

import (
	"context"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/app"
	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	du "github.com/opoccomaxao/wblitz-watcher/pkg/utils/discordutils"
)

func (s *Service) cmdClan(
	event *discordgo.InteractionCreate,
) (*discordgo.InteractionResponse, error) {
	data := event.ApplicationCommandData()

	log.Printf("%s %s\n", data.Name, data.Options[0].Name)

	switch data.Options[0].Name {
	case "add":
		return s.cmdClanAdd(event)
	}

	return nil, app.ErrNotFound
}

func (s *Service) cmdClanAdd(
	event *discordgo.InteractionCreate,
) (*discordgo.InteractionResponse, error) {
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
		data := &discordgo.InteractionResponseData{
			Title: "Clan not found",
			Flags: discordgo.MessageFlagsEphemeral,
		}

		if err != nil {
			data.Content = err.Error()
		}

		//nolint:wrapcheck
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: data,
		}, err
	}

	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
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
			Flags: discordgo.MessageFlagsEphemeral,
		},
	}, nil
}
