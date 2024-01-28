package discord

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (s *Service) cmdHelp(
	ctx context.Context,
	event *discordgo.InteractionCreate,
	_ *CommandData,
) (*Response, error) {
	cmds, err := s.getServerCommands(ctx, event.GuildID)
	if err != nil {
		return nil, err
	}

	ids := s.parseCommandIDs(cmds)

	embed := &discordgo.MessageEmbed{
		Title: "Available commands",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Public commands",
				Value: fmt.Sprintf(
					`</help:%s> - print this help
</ping:%s> - check server
</user stats:%s> - print user stats
</clan list:%s> - list of clans for notifications`,
					ids[CommandFullName{Name: "help"}],
					ids[CommandFullName{Name: "ping"}],
					ids[CommandFullName{Name: "user", SubName: "stats"}],
					ids[CommandFullName{Name: "clan", SubName: "list"}],
				),
				Inline: false,
			},
			{
				Name: "Admin commands",
				Value: fmt.Sprintf(
					`</clan add:%s> - add clan to watch
</clan remove:%s> - remove clan from watch
</channel bind:%s> - bind channel for notifications`,
					ids[CommandFullName{Name: "clan", SubName: "add"}],
					ids[CommandFullName{Name: "clan", SubName: "remove"}],
					ids[CommandFullName{Name: "channel", SubName: "bind"}],
				),
				Inline: false,
			},
		},
		Footer: s.copyrightFooter(),
	}

	return &Response{
		Embeds: []*discordgo.MessageEmbed{embed},
	}, nil
}
