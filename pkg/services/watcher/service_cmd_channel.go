package watcher

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
)

func (s *Service) cmdChannelBind(
	event *discordgo.InteractionCreate,
	data *discord.CommandData,
) (*discord.Response, error) {
	instance := models.BotInstance{
		ServerID:  event.GuildID,
		ChannelID: data.String("channel"),
		Type:      models.SubscriptionType(data.String("type")),
	}

	err := s.domain.CreateUpdateInstance(context.TODO(), &instance)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return &discord.Response{
		Content: "Channel bound",
	}, nil
}
