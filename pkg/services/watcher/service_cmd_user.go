package watcher

import (
	"context"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/jsonutils"
)

func (s *Service) cmdUserStats(
	_ *discordgo.InteractionCreate,
	data *discord.CommandData,
) (*discord.Response, error) {
	req := wg.AccountListRequest{
		Search: data.String("username"),
		Region: models.Region(data.String("server")),
	}

	user, err := s.getUserStatsByNick(req)

	if user == nil {
		return &discord.Response{
			Content: "User not found",
		}, nil
	}

	if err != nil {
		return nil, err
	}

	return &discord.Response{
		Content: "User stats",
		Embeds: []*discordgo.MessageEmbed{
			{
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Wins",
						Value:  user.StatWins(),
						Inline: true,
					},
					{
						Name:   "Damage",
						Value:  user.StatDamage(),
						Inline: true,
					},
					{
						Name:   "Battles",
						Value:  user.StatBattles(),
						Inline: true,
					},
				},
				Author: &discordgo.MessageEmbedAuthor{
					Name: user.AuthorName(),
				},
			},
		},
	}, nil
}

func (s *Service) getUserStatsByNick(
	req wg.AccountListRequest,
) (*wg.AccountInfo, error) {
	log.Printf("userstats %s [%s]\n", req.Search, req.Region.Pretty())

	ctx := context.Background()

	list, err := s.wg.AccountList(ctx, req)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	var id int64

	req.Search = strings.ToLower(req.Search)

	for _, ale := range list {
		if strings.ToLower(ale.Nickname) == req.Search {
			id = ale.AccountID

			break
		}
	}

	if id == 0 {
		return nil, nil
	}

	info, err := s.wg.AccountInfo(ctx, wg.AccountInfoRequest{
		Region: req.Region,
		IDs:    []int64{id},
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return info[jsonutils.MaybeInt(id)], nil
}
