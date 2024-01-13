package watcher

import (
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
)

func (s *Service) embedError(
	err error,
) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "Error",
		Type:  discordgo.EmbedTypeRich,
		Color: int(ColorError),
	}
}

func (s *Service) embedClan(
	clan *wg.ClanListEntry,
) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
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
		},
	}
}
