package app

import (
	"context"
	"log"
	"time"
	"wblitz-watcher/app/diff"
	"wblitz-watcher/sender"
	"wblitz-watcher/sender/discord"
	"wblitz-watcher/wg/types"
)

func (app *App) NotifyClanDiff(
	ctx context.Context,
	clanInfo *types.ClanInfo,
	clanDiff *diff.Total,
) error {
	log.Printf("%s\n", clanDiff.Pretty())

	return nil
}

func (app *App) buildClanEvent(
	clan *types.ClanInfo,
	player *types.Player,
	eventType diff.Type,
) *discord.Message {
	embed := discord.Embed{
		Type:      discord.EmbedTypeRich,
		Timestamp: time.Now(),
	}

	switch eventType {
	case DiffEnter:
		embed.Color = discord.ColorEnter
		embed.Description = StringEnter
	case DiffLeave:
		embed.Color = discord.ColorLeave
		embed.Description = StringLeave
	default:
		return nil
	}

	return &discord.Message{
		Embeds: []discord.Embed{embed},
	}
}

func (app *App) SendDiscordMessage(ctx context.Context, message *discord.Message) error {
	return app.sender.Request(ctx, &sender.Request{
		URL:  app.config.DiscordURL,
		Body: message,
	})
}
