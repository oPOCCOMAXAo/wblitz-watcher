package watcher

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/samber/lo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (s *Service) embedError(
	_ error,
) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: MessageError,
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

func (s *Service) embedClanEvent(
	event *models.EventClan,
	clan *wg.ClanInfo,
	user *wg.AccountInfo,
) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Type:      discordgo.EmbedTypeRich,
		Timestamp: time.Unix(event.Time, 0).Format(time.RFC3339),
		Fields:    []*discordgo.MessageEmbedField{},
	}

	switch event.Type {
	case models.ETCEnter:
		embed.Color = int(ColorEnter)
		embed.Description = MessageEnter
	case models.ETCLeave:
		embed.Color = int(ColorLeave)
		embed.Description = MessageLeave
	}

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "Clan",
		Value:  clan.StatName(),
		Inline: false,
	})

	embed.Author = &discordgo.MessageEmbedAuthor{
		Name: user.AuthorName(),
	}

	embed.Fields = append(embed.Fields,
		&discordgo.MessageEmbedField{
			Name:   MessageWins,
			Value:  user.StatWins(),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   MessageDamage,
			Value:  user.StatDamage(),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   MessageBattles,
			Value:  user.StatBattles(),
			Inline: true,
		},
	)

	return embed
}
