package watcher

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/domain"
)

func (s *Service) cmdChannelBind(
	ctx context.Context,
	event *discordgo.InteractionCreate,
	data *discord.CommandData,
) (*discord.Response, error) {
	request := domain.ChannelBindRequest{
		ServerID:  event.GuildID,
		ChannelID: data.String("channel"),
		Type:      models.SubscriptionType(data.String("type")),
	}

	err := s.domain.ChannelBind(ctx, &request)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return &discord.Response{
		Content: "Channel bound",
	}, nil
}
