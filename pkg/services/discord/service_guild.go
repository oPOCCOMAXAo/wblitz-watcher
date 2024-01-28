package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func (s *Service) GetAllGuilds(
	ctx context.Context,
) ([]*discordgo.UserGuild, error) {
	var (
		res    []*discordgo.UserGuild
		lastID string
	)

	for {
		guilds, err := s.session.UserGuilds(100, "", lastID, s.requestOptions(ctx)...)
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}

		if len(guilds) == 0 {
			break
		}

		res = append(res, guilds...)
		lastID = guilds[len(guilds)-1].ID
	}

	return res, nil
}
