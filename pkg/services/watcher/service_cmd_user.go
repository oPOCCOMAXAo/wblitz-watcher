package watcher

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/jsonutils"
)

func (s *Service) cmdUserStats(
	event *discordgo.InteractionCreate,
) (*discordgo.InteractionResponse, error) {
	var req wg.AccountListRequest

	for _, option := range event.ApplicationCommandData().Options {
		switch option.Name {
		case "username":
			req.Search = option.StringValue()
		case "server":
			req.Region = wg.RegionFromName(option.StringValue())
		}
	}

	user, err := s.getUserStatsByNick(req)

	if user == nil {
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "User not found",
			},
		}, nil
	}

	if err != nil {
		return nil, err
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: event.ApplicationCommandData().Name,
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
		},
	}, nil
}

func (s *Service) getUserStatsByNick(
	req wg.AccountListRequest,
) (*wg.AccountInfo, error) {
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

	for _, ale := range list {
		if ale.Nickname == req.Search {
			id = ale.AccountID

			break
		}
	}

	if id == 0 {
		return nil, nil
	}

	info, err := s.wg.AccountInfo(context.Background(), wg.AccountInfoRequest{
		Region: req.Region,
		IDs:    []int64{id},
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return info[jsonutils.MaybeInt(id)], nil
}
