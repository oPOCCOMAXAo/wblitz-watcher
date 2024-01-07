package watcher

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	du "github.com/opoccomaxao/wblitz-watcher/pkg/utils/discordutils"
)

func (s *Service) cmdChannel(
	event *discordgo.InteractionCreate,
) (*discord.Response, error) {
	data := event.ApplicationCommandData()

	switch data.Options[0].Name {
	case "bind":
		return s.cmdChannelBind(event)
	}

	return nil, models.ErrNotFound
}

func (s *Service) cmdChannelBind(
	event *discordgo.InteractionCreate,
) (*discord.Response, error) {
	err := s.discord.VerifyAccess(event)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	var instance models.BotInstance

	err = du.ParseOptions(event.ApplicationCommandData().Options[0].Options, du.DecodersMap{
		"channel": du.DecoderChannelID(&instance.ChannelID),
		"type":    du.DecoderString(&instance.Type),
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	instance.ServerID = event.GuildID

	err = s.domain.CreateUpdateInstance(context.TODO(), &instance)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return &discord.Response{
		Content: "Channel bound",
	}, nil
}
