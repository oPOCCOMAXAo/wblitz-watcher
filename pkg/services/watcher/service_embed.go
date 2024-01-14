package watcher

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/samber/lo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
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

func (s *Service) embedAccountInfo(
	account *wg.AccountInfo,
) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Wins",
				Value:  account.StatWins(),
				Inline: true,
			},
			{
				Name:   "Damage",
				Value:  account.StatDamage(),
				Inline: true,
			},
			{
				Name:   "Battles",
				Value:  account.StatBattles(),
				Inline: true,
			},
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name: account.AuthorName(),
		},
	}
}

func (s *Service) embedClan(
	clan *models.WGClan,
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
				Name:  "Region",
				Value: clan.Region.Pretty(),
			},
		},
	}
}

func (s *Service) embedClanList(
	clans []*models.WGClan,
) []*discordgo.MessageEmbed {
	const maxEmbedFields = 25

	res := make([]*discordgo.MessageEmbed, 0, len(clans)/maxEmbedFields+1)

	for _, group := range lo.Chunk(clans, maxEmbedFields) {
		embed := &discordgo.MessageEmbed{
			Fields: make([]*discordgo.MessageEmbedField, len(group)),
		}
		res = append(res, embed)

		for i, clan := range group {
			embed.Fields[i] = &discordgo.MessageEmbedField{
				Name:  fmt.Sprintf("[%s] (%s)", clan.Tag, clan.Region.Pretty()),
				Value: clan.Name,
			}
		}
	}

	return res
}
